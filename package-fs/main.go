package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"syscall"
)

func main() {
	long := flag.Bool("l", false, "use a long listing format")

	flag.Parse()

	dir, err := os.Open("/etc")
	if err != nil {
		fmt.Printf("Error opening current directory: %v\n", err)
		return
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	// calculate total size
	var totalSize int64
	for _, entry := range entries {
		totalSize += entry.Size()
	}
	fmt.Printf("total %d\n", totalSize/512)

	if *long {
		for _, entry := range entries {
			stat := entry.Sys().(*syscall.Stat_t)
			fmt.Sprintln(entry.Sys().(*syscall.Stat_t))
			usr, err := user.LookupId(fmt.Sprint(stat.Uid))
			if err != nil {
				fmt.Printf("Error looking up user: %v\n", err)
				continue
			}
			grp, err := user.LookupGroupId(fmt.Sprint(stat.Gid))
			if err != nil {
				fmt.Printf("Error looking up group: %v\n", err)
				continue
			}
			fmt.Printf(
				"%s %d %s %s %d %s %s\n",
				entry.Mode().String(),
				stat.Nlink,
				usr.Username,
				grp.Name,
				entry.Size(),
				entry.ModTime().Format("Jan 02 15:04"),
				entry.Name(),
			)
		}
	} else {
		for _, entry := range entries {
			fmt.Println(entry.Name())
		}
	}
}
