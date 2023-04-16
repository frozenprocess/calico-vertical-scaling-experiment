package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("waiting a lot ...")
	time.Sleep(time.Hour * 10000)
}
