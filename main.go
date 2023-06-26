package main

import (
	"Dirscan/cmd"
	"github.com/gookit/color"
	"time"
)



func main() {
	start := time.Now()

	cmd.Run()

	end := time.Since(start) 

	color.HiGreen.Printf("\n运行时间为: %v\n", end)
}
