package main

import (
	"fmt"
	"log"
)

func main() {
	cache := newLRUCache(5)
	log.Printf("Cache inicial:\n%+v\n\n", cache)

	cache.put("primeiro", "a informacao")
	cache.put("segundo", "a informacao 2")
	cache.put("terceiro", "a informacao 3")
	cache.put("quarto", "a informacao 4")
	cache.put("quinto", "a informacao 5")

	log.Printf("Cache inicializado:\n%+v\n\n", cache)
	for curr := cache.head; curr != nil; curr = curr.next {
		fmt.Printf("\n%+v", curr)
	}
	fmt.Print("\n\n")

	prim, achou := cache.retrieve("primeiro")
	log.Printf("Valor -> %v, achou -> %v \n\n", prim, achou)
	log.Printf("Cache apos a primeira consulta:\n%+v\n\n", cache)

	for curr := cache.head; curr != nil; curr = curr.next {
		fmt.Printf("\n%+v", curr)
	}
	fmt.Print("\n\n")

	cache.put("sexto", "a informacao 6")
	log.Printf("Cache apos colocar o sexto item:\n%+v\n\n", cache)
	for curr := cache.head; curr != nil; curr = curr.next {
		fmt.Printf("\n%+v", curr)
	}
	fmt.Print("\n\n")

	cache.put("setimo", "a informacao 7")
	log.Printf("Cache apos colocar o setimo item:\n%+v\n\n", cache)
	for curr := cache.head; curr != nil; curr = curr.next {
		fmt.Printf("\n%+v", curr)
	}
	fmt.Print("\n\n")
}
