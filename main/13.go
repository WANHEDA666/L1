package main

func main() {
	reverse(1, 2)
}

func reverse(a int, b int) (int, int) {
	a, b = b, a //Чтобы не создавать новую переменную, Golang может просто свайпнуть ссылки на данные в памяти
	return a, b
}
