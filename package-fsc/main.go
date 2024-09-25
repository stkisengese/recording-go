package main

import (
	"flag"
	"fmt"
	"io/fs"
	"math"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

type Options struct {
	LongFormat    bool
	Recursive     bool
	ShowHidden    bool
	Reverse       bool
	SortByTime    bool
	HumanReadable bool
	SortBySize    bool
	Classify      bool
	OnePerLine    bool
}

type FileInfo struct {
	Name    string
	Size    int64
	Mode    os.FileMode
	ModTime time.Time
	IsDir   bool
	Nlink   uint64
	Uid     uint32
	Gid     uint32
}

func main() {
	options, dirs := parseFlags()

	if len(dirs) == 0 {
		dirs = []string{"."}
	}

	for i, dir := range dirs {
		if len(dirs) > 1 {
			if i > 0 {
				fmt.Println()
			}
			fmt.Printf("%s:\n", dir)
		}
		if options.Recursive {
			listRecursive(dir, options)
		} else {
			listDir(dir, options)
		}
	}
}

func parseFlags() (Options, []string) {
	var options Options
	flag.BoolVar(&options.LongFormat, "l", false, "use a long listing format")
	flag.BoolVar(&options.Recursive, "R", false, "list subdirectories recursively")
	flag.BoolVar(&options.ShowHidden, "a", false, "include hidden files")
	flag.BoolVar(&options.Reverse, "r", false, "reverse order while sorting")
	flag.BoolVar(&options.SortByTime, "t", false, "sort by modification time")
	flag.BoolVar(&options.HumanReadable, "h", false, "print sizes in human readable format")
	flag.BoolVar(&options.SortBySize, "S", false, "sort by size")
	flag.BoolVar(&options.Classify, "F", false, "append indicator (one of */=>@|) to entries")
	flag.BoolVar(&options.OnePerLine, "1", false, "list one file per line")
	flag.Parse()
	return options, flag.Args()
}

func listRecursive(path string, options Options) {
	filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("ls: cannot access '%s': %v\n", p, err)
			return nil
		}
		if d.IsDir() {
			if p != path {
				fmt.Printf("\n%s:\n", p)
			}
			files, err := readDir(p, options)
			if err != nil {
				fmt.Printf("ls: cannot access '%s': %v\n", p, err)
				return nil
			}
			printFiles(files, options)
		}
		return nil
	})
}

func listDir(path string, options Options) {
	files, err := readDir(path, options)
	if err != nil {
		fmt.Printf("ls: cannot access '%s': %v\n", path, err)
		return
	}
	printFiles(files, options)
}

func readDir(path string, options Options) ([]FileInfo, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1)
	if err != nil {
		return nil, err
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		if !options.ShowHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		info, err := entry.Info()
		if err != nil {
			continue
		}
		stat := info.Sys().(*syscall.Stat_t)
		files = append(files, FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			Mode:    info.Mode(),
			ModTime: info.ModTime(),
			IsDir:   info.IsDir(),
			Nlink:   stat.Nlink,
			Uid:     stat.Uid,
			Gid:     stat.Gid,
		})
	}

	sortFiles(files, options)
	return files, nil
}

func sortFiles(files []FileInfo, options Options) {
	if options.SortByTime {
		sort.Slice(files, func(i, j int) bool {
			if options.Reverse {
				return files[i].ModTime.Before(files[j].ModTime)
			}
			return files[i].ModTime.After(files[j].ModTime)
		})
	} else if options.SortBySize {
		sort.Slice(files, func(i, j int) bool {
			if options.Reverse {
				return files[i].Size < files[j].Size
			}
			return files[i].Size > files[j].Size
		})
	} else if options.Reverse {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name > files[j].Name
		})
	} else {
		sort.Slice(files, func(i, j int) bool {
			return files[i].Name < files[j].Name
		})
	}
}

func printFiles(files []FileInfo, options Options) {
	if options.LongFormat {
		printLongFormat(files, options)
	} else if options.OnePerLine {
		for _, file := range files {
			fmt.Println(formatFileName(file, options))
		}
	} else {
		printColumnar(files, options)
	}
}

func printLongFormat(files []FileInfo, options Options) {
	var totalBlocks int64
	for _, file := range files {
		totalBlocks += file.Size / 512
		if file.Size%512 > 0 {
			totalBlocks++
		}
	}
	fmt.Printf("total %d\n", totalBlocks)

	maxNlinkWidth := 0
	maxUserWidth := 0
	maxGroupWidth := 0
	maxSizeWidth := 0

	for _, file := range files {
		nlinkWidth := len(fmt.Sprintf("%d", file.Nlink))
		if nlinkWidth > maxNlinkWidth {
			maxNlinkWidth = nlinkWidth
		}

		usr, _ := user.LookupId(fmt.Sprint(file.Uid))
		userWidth := len(usr.Username)
		if userWidth > maxUserWidth {
			maxUserWidth = userWidth
		}

		grp, _ := user.LookupGroupId(fmt.Sprint(file.Gid))
		groupWidth := len(grp.Name)
		if groupWidth > maxGroupWidth {
			maxGroupWidth = groupWidth
		}

		sizeWidth := len(fmt.Sprintf("%d", file.Size))
		if sizeWidth > maxSizeWidth {
			maxSizeWidth = sizeWidth
		}
	}

	for _, file := range files {
		usr, _ := user.LookupId(fmt.Sprint(file.Uid))
		grp, _ := user.LookupGroupId(fmt.Sprint(file.Gid))

		size := fmt.Sprintf("%*d", maxSizeWidth, file.Size)
		if options.HumanReadable {
			size = humanReadableSize(file.Size)
		}

		fmt.Printf("%s %*d %-*s %-*s %*s %s %s\n",
			file.Mode.String(),
			maxNlinkWidth, file.Nlink,
			maxUserWidth, usr.Username,
			maxGroupWidth, grp.Name,
			maxSizeWidth, size,
			file.ModTime.Format("Jan _2 15:04"),
			formatFileName(file, options),
		)
	}
}

func printColumnar(files []FileInfo, options Options) {
	termWidth := getTerminalWidth()

	maxWidth := 0
	for _, file := range files {
		width := len(formatFileName(file, options))
		if width > maxWidth {
			maxWidth = width
		}
	}

	colWidth := maxWidth + 2
	numCols := termWidth / colWidth
	if numCols == 0 {
		numCols = 1
	}

	numRows := int(math.Ceil(float64(len(files)) / float64(numCols)))

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			idx := j*numRows + i
			if idx < len(files) {
				fmt.Printf("%-*s", colWidth, formatFileName(files[idx], options))
			}
		}
		fmt.Println()
	}
}

func formatFileName(file FileInfo, options Options) string {
	name := file.Name
	if options.Classify {
		if file.IsDir {
			name += "/"
		} else if file.Mode&0o111 != 0 {
			name += "*"
		}
	}
	return name
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}

func getTerminalWidth() int {
	defaultWidth := 80

	// Try to get the terminal size using TIOCGWINSZ ioctl
	var size [4]uint16
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&size))); err == 0 {
		return int(size[1])
	}

	// If ioctl fails, try to get the COLUMNS environment variable
	if cols := os.Getenv("COLUMNS"); cols != "" {
		if width, err := strconv.Atoi(cols); err == nil {
			return width
		}
	}

	// If all else fails, return the default width
	return defaultWidth
}
