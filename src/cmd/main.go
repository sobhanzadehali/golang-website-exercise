package main

import (
	"fmt"

	"github.com/sobhanzadehali/golang-website-exercise.git/api"
)

func main() {
	api.InitServer()
	fmt.Println("Server is running on port 5005")
}
