package main

import "math"

func main() {
	const filename = "consts.go"
	const a, b = 3, 4
	var c = int(math.Sqrt(a*a + b*b))
	println(filename, c)

	enums()
	enumsIota()
}

func enums() {
	const (
		cpp    = 0
		java   = 1
		python = 2
		golang = 3
	)
	println(cpp, java, python, golang)
}

func enumsIota() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	println(cpp, python, golang, javascript)
	println(b, kb, mb, gb, tb, pb)
}
