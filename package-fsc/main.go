package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"syscall"
	"time"
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
	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("ls: cannot access '%s': %v\n", p, err)
			return nil
		}
		if info.IsDir() {
			if p != path {
				fmt.Printf("\n%s:\n", p)
			}
			listDir(p, options)
		}
		return nil
	})
}

func listDir(path string, options Options) {
	dir, err := os.Open(path)
	if err != nil {
		fmt.Printf("ls: cannot access '%s': %v\n", path, err)
		return
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("ls: cannot read directory '%s': %v\n", path, err)
		return
	}

	files := make([]FileInfo, 0, len(entries))
	for _, entry := range entries {
		if !options.ShowHidden && entry.Name()[0] == '.' {
			continue
		}
		stat := entry.Sys().(*syscall.Stat_t)
		files = append(files, FileInfo{
			Name:    entry.Name(),
			Size:    entry.Size(),
			Mode:    entry.Mode(),
			ModTime: entry.ModTime(),
			IsDir:   entry.IsDir(),
			Nlink:   stat.Nlink,
			Uid:     stat.Uid,
			Gid:     stat.Gid,
		})
	}

	sortFiles(files, options)

	for _, file := range files {
		printFileInfo(file, options)
	}

	if !options.LongFormat && !options.OnePerLine {
		fmt.Println()
	}
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

func printFileInfo(file FileInfo, options Options) {
	if options.LongFormat {
		usr, _ := user.LookupId(fmt.Sprint(file.Uid))
		grp, _ := user.LookupGroupId(fmt.Sprint(file.Gid))

		size := fmt.Sprintf("%d", file.Size)
		if options.HumanReadable {
			size = humanReadableSize(file.Size)
		}

		fmt.Printf(
			"%s %d %s %s %s %s %s",
			file.Mode.String(),
			file.Nlink,
			usr.Username,
			grp.Name,
			size,
			file.ModTime.Format("Jan 02 15:04"),
			file.Name,
		)
		if options.Classify {
			if file.IsDir {
				fmt.Print("/")
			} else if file.Mode&0o111 != 0 {
				fmt.Print("*")
			}
		}
		fmt.Println()
	} else {
		fmt.Print(file.Name)
		if options.Classify {
			if file.IsDir {
				fmt.Print("/")
			} else if file.Mode&0o111 != 0 {
				fmt.Print("*")
			}
		}
		if options.OnePerLine {
			fmt.Println()
		} else {
			fmt.Print("  ")
		}
	}
}

func humanReadableSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}