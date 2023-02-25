package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	pl("What is your name ?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err == nil {
		pl("Hello ", name)
	} else {
		log.Fatal(err)
	}

	var vName string = "Salim"
	var v1, v2 = 1.2, 3.4
	var v3 = "Hello"
	v4 := 2.4

	cV1 := 1.5
	cV2 := int(cV1)
}
