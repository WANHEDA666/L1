package main

import "fmt"

func main() {
	firstSet := createSet([]string{"cat", "cat", "dog", "cat", "tree"})        //Создаем из массива Set
	secondSet := createSet([]string{"dog", "cat", "tree", "orange", "street"}) //Создаем из массива Set
	result := findDuplicate([]map[string]bool{firstSet, secondSet})
	fmt.Println(result)
}

func createSet(arr []string) map[string]bool {
	set := make(map[string]bool) //Создаем самодельный Set
	for _, item := range arr {   //Проходимся по массиву
		if _, ok := set[item]; !ok { //Проверяем наличие значения в Set
			set[item] = true //Если значения нет, добавляем
		}
	}
	return set
}

func findDuplicate(maps []map[string]bool) map[string]bool { //Проверяем наличие каждого элемента в обоих массивах
	result := make(map[string]bool)
	for key, _ := range maps[0] {
		for key1, _ := range maps[1] {
			if key == key1 {
				result[key] = true
			}
		}
	}
	return result
}
