package sorting

//選擇排序法
var a int
func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] > arr[min] {
				min = j
			}
		}
		if min != i {
			// 代表 i 起始位置不是最小的
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}