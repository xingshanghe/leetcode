package main

import "fmt"

func totalFruit(tree []int) int {
	length := len(tree)
	halfLine := length / 2
	total := int(0)

	for i := 0; i < length; i++ {
		///
		fmt.Println(tree[i:])
		getTotal := getFruit(tree[i:])
		if total < getTotal {
			total = getTotal
		}
		if total > halfLine {
			break
		}
	}

	return total
}

func getFruit(subTree []int) int {
	total := int(0)
	fruit := make(map[int]int)
	for i := 0; i < len(subTree); i++ {
		// 不存在该水果存储
		if _, ok := fruit[subTree[i]]; !ok {
			fruit[subTree[i]] = 1
		} else {
			fruit[subTree[i]] += 1
		}
		if len(fruit) <= 2 {
			continue
		} else {
			delete(fruit, subTree[i])
			break
		}
	}

	for _, v := range fruit {
		total += v
	}
	return total
}

func main() {
	tree := []int{1, 2, 1}
	fmt.Println(totalFruit(tree))
	tree = []int{0, 1, 2, 2}
	fmt.Println(totalFruit(tree))
	tree = []int{1, 2, 3, 2, 2}
	fmt.Println(totalFruit(tree))
	tree = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	fmt.Println(totalFruit(tree))
}
