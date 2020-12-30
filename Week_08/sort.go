package main

import "math/rand"

//选择排序
func sort(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[min], arr[i] = arr[i], arr[min]
	}
}

//插入排序
func sort(arr []int) {
	l := len(arr)

	for i := 0; i < l; i++ {
		for j := i; j >= i; j-- {
			if j > 0 && arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

//冒泡排序
func sort(arr []int) {
	l := len(arr)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//希尔排序,本质上是插入排序的优化
//是不稳定排序,最坏情况下时间复杂度还是n^2,通常的时间复杂度是nlog2n
//先按照希尔距离进行分组,进行一次插入排序,然后不断往下插入排序
func shellsort(list []int) {
	for inc := len(list) / 2; inc > 0; inc = (inc + 1) * 5 / 11 {
		for i := inc; i < len(list); i++ {
			j, tmp := i, list[i]
			for ; j >= inc && list[j-inc] > list[j]; j -= inc {
				list[j] = list[j-inc]
			}
			list[j] = tmp
		}
	}
}

//高级排序 nlogn
//快速排序
//把一个序列分成较大和较小的两个子序列,递归的进行排序
//递归有可能导致栈溢出(无限循环)
func quickSort(a []int) []int {
	//special
	if len(a) < 2 {
		return a
	}
	//declared
	right, left := len(a)-1, 0
	pivot := rand.Int() % len(a)

	a[right], a[pivot] = a[pivot], a[right]

	//loop,move less to left
	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	//drill down
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

//归并排序
//合并两个有序数组
//稳定的排序,需要额外的存储
//一分为2,递归进行
func MergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	mid := len(slice) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

//合并两个有序数组
func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}

	return slice
}

//堆排序
//1.create a max heap
//2.swap the first and last element,remove the last element after swap
type Heap struct {
}

func (heap *Heap) HeapSort(array []int) {
	heap.BuildHeap(array)

	for length := len(array); length > 1; length-- {
		heap.RemoveTop(array, length)
	}
}

func (heap *Heap) BuildHeap(array []int) {
	for i := len(array) / 2; i >= 0; i-- {
		heap.Heapify(array, i, len(array))
	}
}

func (heap *Heap) RemoveTop(array []int, length int) {
	var lastIndex = length - 1
	array[0], array[lastIndex] = array[lastIndex], array[0]
	heap.Heapify(array, 0, lastIndex)
}

func (heap *Heap) Heapify(array []int, root, length int) {
	var max = root
	var l, r = heap.Left(array, root), heap.Right(array, root)

	if l < length && array[l] > array[max] {
		max = l
	}

	if r < length && array[r] > array[max] {
		max = r
	}

	if max != root {
		array[root], array[max] = array[max], array[root]
		heap.Heapify(array, max, length)
	}
}

func (*Heap) Left(array []int, root int) int {
	return (root * 2) + 1
}

func (*Heap) Right(array []int, root int) int {
	return (root * 2) + 2
}
