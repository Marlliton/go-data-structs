package main

import (
	"fmt"

	hashtable "github.com/Marlliton/temp/pkg/hash-table"
)

// NOTE: arquivo usado para testar coisas esporádicas
func main() {

	// var hash int32
	// for _, char := range key {
	// 	hash += char
	// }

	ht := hashtable.New()
	ht.Put("marlliton", "asdf")

	v, b := ht.Get("marlliton")
	fmt.Println("\n\n hash final", v, b)
}
