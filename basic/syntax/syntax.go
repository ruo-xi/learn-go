package main

/**
包内变量
*/
var aa = 3
var bb = "kkk"

var (
	cc = "asd"
	dd = 123
)

//bb := true
func variableZeroValue() {
	var a int
	var s string
	println(a, s)
	println("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	println(a, b, c, s)
}
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	println(a, b, c, s)
}

func main() {
	println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
}
