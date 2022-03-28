package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var mounted []int

func main() {
	ok, info := EnsureAbsDirExists(os.Args[1])
	fmt.Println(ok)
	fmt.Println(info)

	fmt.Printf("empty: %v\n", IsDirEmpty(os.Args[1]))

	for i := 0; i < 10; i++ {
		mounted_append(i + 1000)
	}
	fmt.Println(mounted)

	fmt.Println(mounted_check(1010))

	mounted_remove(1010)
	fmt.Println(mounted)
}

func EnsureAbsDirExists(path string) (ok bool, errinfo string) {
	if !filepath.IsAbs(path) {
		return false, fmt.Sprintf("Absolute path should be used.(%s is not!)\n", path)
	}

	exist, isdir, errinfo := CheckPath(path)
	if exist {
		if isdir {
			return true, ""
		}
		return false, fmt.Sprintf("Path %s exists and is regular file\n", path)
	}

	return MakePath(path)
}

func CheckPath(path string) (exist, isdir bool, errinfo string) {
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, false, err.Error()
		}
		return false, false, "[state Error] " + err.Error()
	}

	return true, fileinfo.IsDir(), ""
}

func MakePath(path string) (ok bool, errinfo string) {
	if err := os.MkdirAll(path, os.ModeDir); err != nil {
		return false, err.Error()
	}

	return true, ""
}

func IsDirEmpty(path string) bool {
	if dir, err := ioutil.ReadDir(path); err == nil {
		fmt.Print(dir)
		return len(dir) == 0
	}
	return false
}

func mounted_append(m int) {
	mounted = append(mounted, m)
}

func mounted_remove(to_remove int) {
	for i, v := range mounted {
		if v == to_remove {
			mounted = append(mounted[:i], mounted[i+1:]...)
		}
	}
}

func mounted_check(iii int) bool {
	for _, v := range mounted {
		if iii == v {
			return true
		}
	}

	return false
}
