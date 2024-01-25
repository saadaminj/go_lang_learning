package main

import (
	"fmt"
	"time"

	tinytime "github.com/wagslane/go-tinytime"
)

func main() {
	currentTime := time.Now()
	tt := tinytime.New(uint32(currentTime.Unix()))
	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
