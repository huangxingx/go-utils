package order

// 冒泡排序
// 时间复杂度 logN^2

//Bubble 冒泡
func Bubble(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < len(array)-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
