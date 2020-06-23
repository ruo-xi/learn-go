package loop

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	println(convertToBin(5),
		convertToBin(13))

	printFile("basic/abc.txt")

	//forever()
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(fileame string) {
	file, err := os.Open(fileame)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		println(scanner.Text())
	}
}
func forever() {
	for {
		println("gg")
	}
}
