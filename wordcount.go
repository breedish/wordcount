package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(0)
		return
	}
	fmt.Println(countWords(args[1]))
}

func countWords(str string) int {
	return len(strings.Split(str, " "))
}

func readString() (string, error) {
	r := bufio.NewReader(os.Stdin)
	str, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", err
	}

	return strings.TrimSpace(str), nil
}
