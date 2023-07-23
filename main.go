package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Val   string
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
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
	head := &Node{}
	tail := &Node{}

	//create an empty queue where the head right points to the tail and tail left points to head
	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail}
}

/**
 *
 * check if the value already exits in the queue
 * if exist remove and add at the beginning
 * else add at the beginning
 **/
func (c *Cache) Check(str string) {

	//create empty node
	node := &Node{}

	//check if value exists
	if val, ok := c.Hash[str]; ok {
		//remove val
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}

	//add node
	c.Add(node)

	c.Hash[str] = node
}

func main() {
	fmt.Println("Start cache")
	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "tree", "text", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
