package tool

import (
	"fmt"
	"math/rand"
	"time"
)

func buildSlice(size int) []int {
	var (
		result []int
	)
	index := 0
	for index < size {
		result = append(result, index)
		index += 1
	}
	return result
}
func swap(slice []int, i, j int) {
	tmp := slice[i]
	slice[i] = slice[j]
	slice[j] = tmp
}
func shuffle(slice []int) {
	length := len(slice) - 1
	for length > 0 {
		rand.Seed(time.Now().Unix())
		swap(slice, length, rand.Intn(length))
		length -= 1
	}
}

func quickSort(array []int, low, high int) {
	if high-low <=0{
		return
	}
	i:=low
	j:=high
	mid := array[i]
	for i < j {
		if array[j] >= mid && i < j {
			j--
		}
		array[i] = array[j]
		if array[i] <= mid && i < j {
			i++
		}
		array[j] = array[i]
	}
	array[i] = mid
	quickSort(array, low, i-1)
	quickSort(array, i+1, high)
}

func sort(array []int) {
	i := 0
	j := len(array) - 1
	quickSort(array, i, j)
}

func sortTest() {
	array := buildSlice(10)
	shuffle(array)
	fmt.Printf("%v\n", array)
	sort(array)
	fmt.Printf("%v\n", array)
}
