package main

func mergeHelper(arr []int, channel chan int) {
	if len(arr) <= 1 {
		channel <- 0
		return
	}

	m := len(arr) / 2
	var left, right []int

	chann := make(chan int)

	left = append(left, arr[:m]...)
	right = append(right, arr[m:]...)

	go mergeHelper(left, chann)
	go mergeHelper(right, chann)

	<-chann
	<-chann

	// Sorting occur here

	Merge(arr, left, right)
	channel <- 0

}

func MergeSortThreaded(arr []int) {
	channel := make(chan int)
	go mergeHelper(arr, channel)
	<-channel
}

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	// Partition occur

	m := len(arr) / 2
	var left, right []int

	left = append(left, arr[:m]...)
	right = append(right, arr[m:]...)

	MergeSort(left)
	MergeSort(right)

	// Sorting
	Merge(arr, left, right)

}

