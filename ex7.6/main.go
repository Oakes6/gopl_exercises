package main

import (
	"flag"
	"fmt"
	"tempconv"
)

var temp = tempconv.CelciusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
