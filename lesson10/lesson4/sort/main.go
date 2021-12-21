package sort

func Sort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		number := numbers[i]
		j := i
		for ; j >= 1 && numbers[j-1] > number; j-- {
			numbers[j] = numbers[j-1]
		}
		numbers[j] = number
	}
}
