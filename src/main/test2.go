package main

import (
	"fmt"
	"strconv"
)

func main1() {
	fmt.Println( xtob("0xGG"))
}

func xtob(x string) string {
    base, _ := strconv.ParseInt(x, 16, 10)
    return strconv.FormatInt(base, 2)
}

