package pointer

func main() {
	a, b := 3, 4
	a, b = swap(a, b)
	println(a, b)
	a, b = swap(a, b)
	swapByPointer(&a, &b)
	println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
func swapByPointer(a, b *int) {
	*b, *a = *a, *b
}
