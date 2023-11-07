package main

func main() {
	arr := []int{88, 222, 12, 5, 9, 6, 73}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, left int, right int) {
	if left < right { //Проверяем не состит ли отрезок из одного элемента
		anchor := right                            //Опорным элементом выбираем последний
		bAnchor := left                            //Приравниваем больший якорь к первому элементу
		mAnchor := left                            //Приравниваем меньший якорь к первому элементу
		for mAnchor < anchor && bAnchor < anchor { //Проверяем не дошел ли один из якорей до опорного элемента
			for i := bAnchor; i < anchor; i++ { //Ищем тот, что больше последнего
				if arr[bAnchor] > arr[anchor] {
					break
				}
				bAnchor++
			}
			for i := mAnchor; i < anchor; i++ { //Ищем тот, что меньше последнего
				if arr[mAnchor] < arr[anchor] {
					break
				}
				mAnchor++
			}
			if mAnchor < anchor && bAnchor < anchor { //Проверяем не дошел ли один из якорей до опорного элемента
				arr[bAnchor], arr[mAnchor] = arr[mAnchor], arr[bAnchor] //Переставляем опорный элемент с самым большим,
				// там слева от большего у нас останутся только меньшие опрного элемента
				bAnchor++
				mAnchor++
			}
		}
		arr[bAnchor], arr[right] = arr[right], arr[bAnchor] //В последний раз переставляем опорный элемент с самым большим
		quickSort(arr, left, bAnchor-1)                     //Сортируем массив слева от опорного элемента
		quickSort(arr, bAnchor+1, right)                    //Сортируем массив справа от опорного элемента
	}
}
