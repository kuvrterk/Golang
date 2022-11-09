package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//Global variables for random gen
const MAX, MIN int = 10000, 10

func GetData(arr []int, count int) {

	for i := 0; i < count; i++ {
		var randNum int = rand.Intn(1000)
		arr[i] = randNum
	}

}

func DisplayNums(arr []int, count int) {

	var countLine int = 0
	for i := 0; i < count; i++ {
		fmt.Print(arr[i])
		fmt.Printf("  ")

		if countLine == 10 {
			fmt.Print("\n")
			countLine = 0
		}

		countLine++
	}
	fmt.Println("")
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])

	return arr
}

func stats(arr []int, count int) {

	var average, variance, standard_dev, i float64
	var floatCount = float64(count)

	for i = 0; i < floatCount; i++ {
		average += float64(arr[int32(i)])
	}

	average = average / floatCount

	for i = 0; i < floatCount; i++ {

		variance = math.Pow((float64(arr[int32(i)]) - average), 2)
	}

	variance = variance / (floatCount - 1)
	standard_dev = math.Sqrt(variance)

	fmt.Println("The average is: ", average)
	fmt.Println("The Variance is: ", variance)
	fmt.Println("The Standard Deviation is ", standard_dev)
}

func main() {
	//Get cmd line argumrnts then check between min and max

	no, err := strconv.Atoi(os.Args[1])
	var count int = 0

	//Error checking
	if err != nil {
		fmt.Println(err)
		fmt.Println("Couldn't convert: " + os.Args[1])
	}

	//If the count
	if no > MIN && no < MAX {
		count = no
	} else {
		fmt.Println("The inputted number is more or less then the MAX and Min")
		os.Exit(1)
	}

	var randList [MAX]int

	var arrayList = randList[0:count]
	GetData(arrayList, count)
	fmt.Println("Here is the array before the sort\n")
	DisplayNums(arrayList, count)

	quickSort(arrayList)
	fmt.Println("Here is the array After the sort\n")
	DisplayNums(arrayList, count)
	stats(arrayList, count)

}
