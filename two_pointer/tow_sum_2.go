package main


func twoSum(numbers []int, target int) []int {
	var left int = 0
	var right int = len(numbers) - 1
	for left < right {
		if numbers[left] + numbers[right] == target {
			return []int{left, right}
		}
		if numbers[left] + numbers[right] > target {
			right -= 1
		} else {
			left += 1
		}
	}
	return []int{}
}

//func main() {
//	numbers := twoSum([]int{2, 7, 11, 15}, 9)
//	fmt.Println(numbers) // Выводит [0 1], так как 2 + 7 = 9
//}
//https://habr.com/ru/articles/774796/