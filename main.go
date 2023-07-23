package main

import (
	"fmt"
)

type Node struct {
}

type Queue struct {
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

/**
 *
 * func to run a cache
 **/
func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

/**
 *
 * func to create a queue
 **/
func NewQueue() Queue {

}

func main() {
	fmt.Println("Start cache")
	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "tree", "text", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
