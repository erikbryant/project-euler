package main

import (
	"fmt"
	"strconv"
)

func mul(s string, n int) string {
	val, _ := strconv.Atoi(s)
	val *= n

	// We only care about the last ten digits
	str := strconv.Itoa(val)
	if len(str) > 10 {
		str = str[len(str)-10:]
	}
	return str
}

func main() {
	product := "1"
	for i := 1; i <= 7830457; i++ {
		product = mul(product, 2)
	}
	p, _ := strconv.Atoi(product)
	p *= 28433
	p++
	fmt.Println(p)
}
