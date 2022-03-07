// find windows logical drives (drive letters), for both avaliable and occupied ones
package main

import (
	"fmt"
	"sort"
	. "strconv"
	"syscall"
)

func GetLogicalDrives() (occup []string, avail []string) {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	GetLogicalDrives := kernel32.MustFindProc("GetLogicalDrives")
	n, _, _ := GetLogicalDrives.Call()
	s := FormatInt(int64(n), 2)
	var drives_all = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	avaliable_map := make(map[string]int)
	for i, v := range drives_all {
		avaliable_map[v] = i
	}
	temp := drives_all[0:len(s)]
	var occupied []string
	for i, v := range s {
		l := len(s) - i - 1
		if v == 49 { // 49 '1', 48 '0'
			occupied = append(occupied, temp[l])
		}
	}
	for _, v := range occupied {
		delete(avaliable_map, v)
	}
	var avaliable []string
	for v := range avaliable_map {
		avaliable = append(avaliable, v)
	}
	sort.Slice(occupied, func(i, j int) bool {
		return occupied[i] < occupied[j]
	})
	sort.Slice(avaliable, func(i, j int) bool {
		return avaliable[i] < avaliable[j]
	})

	return occupied, avaliable
}

func main() {
	occupied, avaliable := GetLogicalDrives()
	fmt.Printf("occupied  drives: %s\n", occupied)
	fmt.Printf("avaliable drives: %s\n", avaliable)
}
