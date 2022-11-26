package star

// 冒泡排序
func BubbleSort(nums []int) {
	n := len(nums)
	for i := n - 1; i >= 1; i-- {
		for j := 0; j <= i-1; j++ {
			if nums[j] > nums[j+1] {
				// 相当于swap函数
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}
