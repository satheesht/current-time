package main

import (
	"fmt"
	"time"
)

func main()  {
	currentTime := time.Now()

	fmt.Println("current time: ", currentTime.String())
}