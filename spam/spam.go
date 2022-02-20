package main

import (
	"fmt"
	"os"
)

func replaceAtIndex(in string, r rune, i int) string {

	out := []rune(in)
	out[i] = r
	return string(out)

}

func main() {
	if len(os.Args) == 0 {
		fmt.Println("Please enter your information correctly")
	}
	args := os.Args[1:]
	buffer := make([]string, len(args))

	const mark string = "http://"
	for _, k := range args {
		if len(k) < len(mark) {
			buffer = append(buffer, k)
		} else {
			if k[0:7] == mark {

				for i := 7; i < len(k); i++ {
					k = replaceAtIndex(k, '*', i)

				}
				buffer = append(buffer, k)
			}
		}
	}
	fmt.Println(mark)
	fmt.Println(buffer)
}
