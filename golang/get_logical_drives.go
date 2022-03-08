// func GetLogicalDrives() (occup []string, avail []string)
//     :find windows logical drives (drive letters), for both avaliable and occupied ones
//
// func ShuffleSlice(slice interface{})
//     :shuffle ANY slice (works for whatever the type of slice element is)

package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

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

func ShuffleSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		panic("Must be slice")
	}

	len := rv.Len()
	if len < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := len - 1; i >= 0; i-- {
		j := rand.Intn(len)
		swap(i, j)
	}
}

func main() {
	occupied, avaliable := GetLogicalDrives()
	fmt.Printf("occupied  drives: %s\n", occupied)
	fmt.Printf("avaliable drives: %s\n", avaliable)

	/* 随机乱序 */
	// ***** this works *****
	// rand.Seed(time.Now().UnixNano())
	// rand.Shuffle(len(avaliable), func(i, j int) {
	// 	avaliable[i], avaliable[j] = avaliable[j], avaliable[i]
	// })

	// ***** this better *****
	ShuffleSlice(avaliable)

	fmt.Printf("Shuffle avaliable: %s\n", avaliable)
}
