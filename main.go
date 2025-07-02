package main

import (
	"cmp"
	"fmt"
	"log"
	"os"
)

var Name string

var GH_OUTPUT string

func init() {
	Name = cmp.Or(
		os.Getenv("INPUT_NAME"),
		"World",
	)
}

func output(k string, v string) string {
	return fmt.Sprintf(`%s=%s`, k, v)
}

func main() {
	greeting := fmt.Sprintf("Hello, %v!", Name)
	log.Println(greeting)
	GH_OUTPUT = output("greeting", greeting)
	os.Setenv("GITHUB_OUTPUT", GH_OUTPUT)
}
