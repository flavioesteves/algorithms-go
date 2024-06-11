package sort

func BubbleSort(arr []int) {
	arrLen := len(arr)

	for i := 0; i < arrLen-1; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if arr[j] > arr[j+1] {

				// Single statement swap (recommended)
				arr[j], arr[j+1] = arr[j+1], arr[j]

				// Two step approach
				// tmp := arr[j]
				// arr[j] = arr[j+1]
				// arr[j+1] = tmp
			}
		}
	}
}
