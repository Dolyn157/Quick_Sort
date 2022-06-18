package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Partition(arr *[12]int, left int, right int) int {
	var l int = left
	var r int = right
	//  arr[l] 是每次递归挖坑填数操作的第一个坑，在这个位置上的元素被我们用来当作中轴与数组中其他元素比大小。所有比中轴大的元素都会放在它的右边
	//  所有比它小的元素都会放到它的左边。 随着比大小的进行，左索引 l 的位置将会与右索引 r 重合，届时比大小循环结束， l 和 r 同时指向着中轴元素在数组中的新
	//  位置， 将本回合一开始挖出来的中轴元素放入新位置，本回合结束。将新位置返回给分治函数，对剩下的子序列进行挖数填坑排序操作。
	var pivot int = arr[l]

	for l < r {
		// 从右往左寻找小于 pivot 的数来填充 arr[l]
		for l < r {
			if arr[r] <= pivot {
				arr[l] = arr[r] //将刚才找到的那个比 pivot 小的 arr[r] 填到 arr[l]中， arr[r] 的位置就形成了一个新的坑
				l++
				break
			} else {
				r--
			}
		}
		for l < r {
			// 从左往右寻找大于 pivot 的数来填充 arr[r]
			if arr[l] > pivot {
				arr[r] = arr[l] //将刚才找到的那个比 pivot 大的 arr[l] 填到 arr[r]中， arr[l] 的位置就形成了一个新的坑
				r--
				break
			} else {
				l++
			}
		}
	}
	//退出时，l=r ， 将 pivot 填回 arr[l] 这个上面循环最后挖的坑
	arr[l] = pivot
	// 返回调整后中轴所在的下标位置。
	return l
}

//分治法代码
func _quickSort(arr *[12]int, left int, right int) {
	if left < right {
		pivotIndex := Partition(arr, left, right)
		_quickSort(arr, left, pivotIndex-1)  // 递归地对arr[left..pivotIndex-1] 即上回合排序后基准值左边的数字组成的左子序列重复快速排序，使得左子序列有序
		_quickSort(arr, pivotIndex+1, right) //  递归地对arr[pivotIndex+1..right] 即上回合排序后基准值右边的数字组成的右子序列重复快速排序，使得右子序列有序
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	var arr1 [12]int
	for i, _ := range arr1 {
		arr1[i] = rand.Intn(100)
	}
	fmt.Println("\n Before the sort:", arr1)

	_quickSort(&arr1, 0, len(arr1)-1)

	fmt.Println()
	fmt.Println("\n After the sort:", arr1)

	time.Sleep(time.Second * 10)
}
