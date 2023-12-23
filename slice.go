package main

import "fmt"

func main() {
	names := [...]string{"daud", "hidayat", "ramadhan", "bintang", "rahmatullah"}

	slice1 := names[3:5]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[2:]
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	sliceDays1 := days[5:]
	sliceDays1[0] = "Sabtu Lagi"
	sliceDays1[1] = "Minggu Lagi"

	fmt.Println(days)
	fmt.Println(sliceDays1)

	sliceDays2 := append(sliceDays1, "Libur")
	sliceDays2[0] = "Sabtu Lagi Lagi"
	fmt.Println(sliceDays1)
	fmt.Println(sliceDays2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Daud"
	newSlice[1] = "Hidayat"
	//newSlice[2] = "Ramadhan" error
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Bintang")
	//fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Daud Hidayat"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	isArray := [...]int{1, 2, 3, 4, 5}
	isSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(isArray)
	fmt.Println(isSlice)
}
