package main

func main() {
	i := 7
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = append(arr[:i], arr[i+1:]...) //Собираем новый слайс из слайса до i и после i
}
