package main

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool) //Создаем самодельный Set
	for _, item := range arr {   //Проходимся по массиву
		if _, ok := set[item]; !ok { //Проверяем наличие значения в Set
			set[item] = true //Если значения нет, добавляем
		}
	}
}
