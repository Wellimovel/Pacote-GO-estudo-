package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	h := sha1.New()
	h.Write([]byte("Codigo hash"))
	v := h.Sum([]byte{})
	fmt.Println(v)
}
