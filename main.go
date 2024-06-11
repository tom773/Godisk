package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"syscall"

	f "github.com/tom773/godisk/utils"
)

type Dirs struct {
	DirPath string
	Size    float64
}

func main() {
	dirPaths := []string{"/usr", "/var", "/etc", "/home", "/bin", "/lib", "/lib64", "/opt", "/root", "/sbin", "/srv", "/boot", "/media"}
	var size float64
	var dirs []Dirs

	totalSizeOfDrive := getTotal()
	for _, dirPath := range dirPaths {
		size_, err := getDirSize(dirPath)
		if err != nil {
			fmt.Println(err)
			return
		}
		dir := Dirs{DirPath: dirPath, Size: size_}
		if size_ != 0.0 {
			dirs = append(dirs, dir)
		}
		size += size_
	}
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Size > dirs[j].Size
	})
	for _, dir := range dirs {
		fmt.Printf("%sPath: %s\n%s%.2fGB%s ", f.Colors.Cyan, dir.DirPath, f.Colors.Green, dir.Size, f.Colors.Reset)
		fmt.Printf("(%s%.2f%%%s)\n\n", f.Colors.Yellow, (dir.Size/totalSizeOfDrive)*100, f.Colors.Reset)
	}
	fmt.Printf("~ Total Size: %.2fGB of %.2fGB usable (%.2f%%)\n", size, totalSizeOfDrive, size/totalSizeOfDrive*100)
}

func getDirSize(path string) (float64, error) {
	excluded := map[string]bool{
		"/proc":      true,
		"/sys":       true,
		"/dev":       true,
		"/run":       true,
		"/tmp":       true,
		"/var/tmp":   true,
		"/var/run":   true,
		"/var/cache": true,
	}
	var size int64
	err := filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			if os.IsNotExist(err) || os.IsPermission(err) {
				return nil
			}
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		if info.IsDir() && (excluded[p] || isSubPathOfExcluded(p, excluded)) {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	sizegb := float64(size) / (1024 * 1024 * 1024)
	return sizegb, err
}

func getTotal() float64 {
	var stat syscall.Statfs_t
	syscall.Statfs("/", &stat)
	return float64(stat.Bsize) * float64(stat.Blocks) / (1024 * 1024 * 1024)
}

func isSubPathOfExcluded(p string, excluded map[string]bool) bool {
	for k := range excluded {
		if len(p) > len(k) && p[:len(k)] == k {
			return true
		}
	}
	return false
}
