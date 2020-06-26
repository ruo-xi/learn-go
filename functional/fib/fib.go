package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Fibonacci() intGen {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

type intGen func() int

//func (g intGen) Read(p []byte) (n int, err error) {
//	panic("implement me")
//}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d", next)
	return strings.NewReader(s).Read(p)
}

func (i intGen) ReadByte() (byte, error) {
	panic("implement me")
}

func PrintFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
