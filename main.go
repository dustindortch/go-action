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

func SetOutput(k string, v string) error {
	f, err := os.OpenFile(os.Getenv("GITHUB_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(fmt.Sprintf("%s=%s\n", k, v))); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func Printenv(e string) {
	log.Println(os.Getenv(e))
}

func main() {
	Printenv("GITHUB_OUTPUT")
	greeting := fmt.Sprintf("Hello, %v!", Name)
	log.Println(greeting)
	err := SetOutput("greeting", greeting)
	if err != nil {
		log.Fatalf("Error setting output: %v", err)
	}
}
