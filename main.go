package main

import (
	"fmt"
)

const SIZE = 5

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

/**
 *
 * func to remove node form queue
 **/
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)

	left := n.Left
	right := n.Right

	//connect the nodes on the each side of the node that to be deleted
	left.Right = right
	right.Left = left

	c.Queue.Length -= 1

	delete(c.Hash, n.Val)

	return n

}

/**
 *
 * func to add node to queue
 **/
func (c *Cache) Add(n *Node) {
	fmt.Printf("add : %s\n", n.Val)

	//store the current first node in tmp
	tmp := c.Queue.Head.Right

	// connect the new node to the beginning of the queue
	c.Queue.Head.Right = n
	n.Left = c.Queue.Head

	//connect the previous first node with the new node
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++

	//check if the lenght of the queue doesn't exceed the defined size
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}

}

/**
 *
 * func to display cache
 **/
func (c *Cache) Display() {
	c.Queue.Display()
}

/**
 *
 * display method for queue struct
 **/
func (q *Queue) Display() {
	//get the first node
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)

	//loop arround the queue
	for i := 0; i < q.Length; i++ {

		fmt.Printf("{%s}", node.Val)
		if i < q.Length-1 {
			fmt.Printf("<-->")
		}

		//get the next node
		node = node.Right
	}

	fmt.Printf("]")
}

func main() {
	fmt.Println("Start cache")
	cache := NewCache()

	for _, word := range []string{"parrot", "avocado", "tree", "text", "tomato", "tree"} {
		cache.Check(word)
		cache.Display()
	}
}
