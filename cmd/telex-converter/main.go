package main

import (
	"fmt"

	"github.com/haunt98/ioe-go"
	"github.com/haunt98/telex-go"
)

func main() {
	fmt.Println("Convert to Telex")
	text := ioe.ReadInput()
	result := telex.ConvertText(text)
	fmt.Println(result)
}
