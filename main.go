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
	return fmt.Sprintf(`echo "%s=%s" >> $GITHUB_OUTPUT`, k, v)
}

func main() {
	greeting := fmt.Sprintf("Hello, %v!", Name)
	log.Println(greeting)
	fmt.Println(output("greeting", greeting))
}
