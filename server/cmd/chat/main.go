package main

import (
	"fmt"
	"github.com/udborets/chat-app/server/internal"
)

func main() {
	fmt.Println("server started")
	internal.Init()
	internal.StartServer()
}
