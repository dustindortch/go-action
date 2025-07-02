package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
)

var Name string

func init() {
	Name = cmp.Or(
		os.Getenv("INPUT_NAME"),
		"World",
	)
}

func output(k string, v string) string {
	return fmt.Sprintf(`::set-output name=%s::%s`, k, v)
}

func Printenv(e string) {
	log.Println(os.Getenv(e))
}

func main() {
	Printenv("GITHUB_OUTPUT")
	greeting := fmt.Sprintf("Hello, %v!", Name)
	log.Println(greeting)
	fmt.Println(output("greeting", greeting))
}
