package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	val  string
	next *Node
}

func main() {
	node := Node{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("github.com/mkilic91/goLinkedList")
	fmt.Println("--------------------------------")

	for {
		fmt.Print("-> ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if strings.Compare("add", input) == 0 {
			fmt.Print("#add -> ")
			addInput, _ := reader.ReadString('\n')
			addInput = strings.Replace(addInput, "\n", "", -1)

			node.add(addInput)
		} else if strings.Compare("print", input) == 0 {
			fmt.Println("--------------------------------")
			node.print()
			fmt.Println("--------------------------------")
		} else if strings.Compare("append", input) == 0 {
			fmt.Print("#append after -> ")
			appendAfterInput, _ := reader.ReadString('\n')
			appendAfterInput = strings.Replace(appendAfterInput, "\n", "", -1)

			fmt.Print("#append new val -> ")
			appendNewVal, _ := reader.ReadString('\n')
			appendNewVal = strings.Replace(appendNewVal, "\n", "", -1)

			afterNode := node.search(appendAfterInput)

			newNode := &Node{
				val:  appendNewVal,
				next: nil,
			}

			node.appendTo(afterNode, newNode)
		} else if strings.Compare("sadd", input) == 0 {
			fmt.Print("#sadd -> ")
			saddInput, _ := reader.ReadString('\n')
			saddInput = strings.Replace(saddInput, "\n", "", -1)

			node.sortedAdd(saddInput)
		}

	}
}

func (node *Node) add(val string) {
	if node.next == nil {
		node.next = &Node{
			val:  val,
			next: nil,
		}

	} else {
		node.next.add(val)
	}
}

func (node *Node) appendTo(after *Node, newNode *Node) {
	temp := after.next
	newNode.next = temp
	after.next = newNode
}

func (node *Node) search(val string) *Node {
	if node.val == val {

		return node
	} else if node.next != nil {

		return node.next.search(val)
	} else {

		return &Node{}
	}
}

func (node *Node) sortedAdd(val string) {
	if node.next == nil {
		node.next.add(val)

	} else if node.next.val < val {
		node.next.sortedAdd(val)

	} else {
		node.appendTo(node, &Node{
			val:  val,
			next: nil,
		})

	}
}

func (node *Node) print() {
	fmt.Println(node.val)
	if node.next != nil {
		node.next.print()
	}
}
