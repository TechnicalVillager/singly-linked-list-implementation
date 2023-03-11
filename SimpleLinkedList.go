package main

import "fmt"

// Node struct for Linked List
type node struct {
	data int
	link *node
}

// Linked List struct with Length, Head, Tail
type linkedList struct {
	length int
	head   *node
	tail   *node
}

// The append function adds a node after the last node
// and increases the List length
func (List *linkedList) append(Node *node) {
	if List.length == 0 {
		List.head = Node
		List.tail = List.head
	} else {
		List.tail.link = Node
		List.tail = List.tail.link
	}
	List.length++
}

// The prepend function adds a node before the first node
// and increases the List length
func (List *linkedList) prepend(Node *node) {
	if List.length == 0 {
		List.head = Node
		List.tail = List.head
	} else {
		current := List.head
		List.head = Node
		List.head.link = current
	}
	List.length++
}

// The insert function adds a node based on the key
// position and increases the List length
func (List *linkedList) insert(Node *node, key int) {

	if (key <= List.length) && (key > 0) {

		if key != 1 {
			previous := List.head
			count := 1
			for (previous != nil) && (count < (key - 1)) {
				previous = previous.link
				count++
			}
			next := previous.link
			Node.link = next
			previous.link = Node
			List.length++

		} else {
			List.prepend(Node)
		}

	} else {
		fmt.Printf("%d = Invalid Key!\n", key)
	}
}

// The removeLast function removes the last node
// and decreases the List length
func (List *linkedList) removeLast() {

	var previous *node
	current := List.head

	for (current != nil) && (current.link != nil) {
		previous = current
		current = current.link
	}

	if previous != nil {
		previous.link = nil
		List.tail = previous
	} else {
		List.head = previous
		List.tail = List.head
	}

	if List.length > 0 {
		List.length--
	} else {
		fmt.Println("\n!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Can't Delete: List is Empty")
		fmt.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n")
	}
}

// The removeFirst function removes the first node
// and decreases the List length
func (List *linkedList) removeFirst() {
	current := List.head

	if (current != nil) && (current.link != nil) {
		current.data = current.link.data
		current.link = current.link.link
	} else {
		List.head = nil
		List.tail = List.head
	}

	if List.length > 0 {
		List.length--
	} else {
		fmt.Println("\n!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Can't Delete: List is Empty")
		fmt.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!\n\n")
	}
}

// The removeWithKey function removes a node based
// on the key position and decreases the List length
func (List *linkedList) removeWithKey(key int) {

	if (key <= List.length) && (key > 0) {
		if (key == 1) || (key == List.length) {
			if key == List.length {
				List.removeLast()
			} else {
				List.removeFirst()
			}
		} else {
			previous := List.head
			count := 1
			for (previous != nil) && (count < (key - 1)) {
				previous = previous.link
				count++
			}
			next := previous.link
			previous.link = next.link
			List.length--
		}
	} else {
		fmt.Printf("%d = Invalid Key!\n", key)
	}
}

// General function to print the Linked List
func (List linkedList) printList() {

	fmt.Println("--------------------------------")
	fmt.Println("No. of Element in the List:", List.length)
	fmt.Println("--------------------------------")

	if List.length > 0 {
		fmt.Println("Head:", *List.head)
		fmt.Println("Tail:", *List.tail)
		fmt.Println("------------------")
		count := 1
		for List.head != nil {
			fmt.Printf("Node %d: {Data = %d, Link (Next) = %p}\n", count, List.head.data, List.head.link)
			List.head = List.head.link
			count++
		}
	} else {
		fmt.Println("\n!!!!!!!!!!!!!")
		fmt.Println("List is Empty")
		fmt.Print("!!!!!!!!!!!!!\n\n")
	}
}

func main() {
	myList := linkedList{}
	nodes := [5]int{10, 20, 30, 40, 50}
	myList.removeLast()
	for _, value := range nodes {
		myList.append(&node{data: value})
	}
	myList.printList()

	for range nodes {
		myList.removeLast()
	}
	myList.removeLast()
	myList.printList()

	nodes1 := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, value := range nodes1 {
		myList.append(&node{data: value})
	}
	myList.printList()
	for range nodes {
		myList.removeFirst()
	}
	myList.printList()
	myList.insert(&node{data: 5}, 6)
	myList.insert(&node{data: 15}, 1)
	myList.printList()

	for _, value := range nodes {
		myList.prepend(&node{data: value})
	}
	myList.printList()
	myList.insert(&node{data: 12}, 9)
	myList.printList()
	myList.insert(&node{data: 13}, 6)
	myList.printList()
	myList.removeWithKey(6)
	myList.printList()
	myList.removeWithKey(10)
	myList.printList()
	myList.removeWithKey(1)
	myList.printList()
	myList.removeWithKey(0)
	myList.printList()
}
