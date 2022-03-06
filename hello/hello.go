package main

import (
	"fmt"

	greetings "example/hello/greeting"
)

func main() {
	message := greetings.Hello("thuan")
	fmt.Println(message)
}
