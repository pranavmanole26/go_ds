// This is a singly linked list implementation

package main

import "fmt"

type node struct {
	val  int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (ls *linkedList) isListEmpty() bool {
	if ls.head == nil {
		fmt.Println("List is empty")
		return true
	}
	return false
}

func (ls *linkedList) append(newNode *node) {
	if ls.isListEmpty() {
		ls.head = newNode
		ls.length++
		return
	}
	trav := ls.head
	for trav.next != nil {
		trav = trav.next
	}
	trav.next = newNode
	ls.length++
}

func (ls *linkedList) insert(newNode node, after int) {
	if ls.isListEmpty() {
		return
	}
	trav := ls.head
	inserted := false
	for trav != nil {
		if trav.val == after {
			tempNext := trav.next
			trav.next = &newNode
			newNode.next = tempNext
			inserted = true
			ls.length++
			break
		}
		trav = trav.next
	}
	if !inserted {
		fmt.Printf("Can not insert newNode. afterNode = %v not found\n", after)
	}
}

func (ls *linkedList) display() {
	if ls.isListEmpty() {
		return
	}
	trav := ls.head
	for trav != nil {
		fmt.Printf("%d -> ", trav.val)
		trav = trav.next
	}
	fmt.Printf(" nil\n")
}

func (ls *linkedList) search(val int) {
	if ls.isListEmpty() {
		return
	}
	trav := ls.head
	valFound := false
	for trav != nil {
		if trav.val == val {
			fmt.Printf("Value %d found in list.\n", val)
			valFound = true
			break
		}
		trav = trav.next
	}
	if !valFound {
		fmt.Printf("Value %d not found in list.", val)
	}
}

func (ls *linkedList) remove(n int) {
	if ls.isListEmpty() {
		return
	}
	trav := ls.head
	valueRemoved := false
	if ls.head.val == n {
		tempNode := ls.head
		ls.head = ls.head.next
		tempNode.next = nil
		return
	}
	for trav.next != nil {
		if trav.next.val == n {
			if trav.next.next != nil {
				trav.next = trav.next.next
			} else {
				trav.next = nil
			}
			valueRemoved = true
			fmt.Printf("Value %d removed from list.\n", n)
			break
		}
		trav = trav.next
	}
	if !valueRemoved {
		fmt.Printf("%d value not removed from the list", n)
	}
}

func main() {
	ls := linkedList{head: nil, length: 0}
	var choice int
	for {
		var newVal int

		fmt.Println("Choices: ")
		fmt.Println("1. Append")
		fmt.Println("2. Insert")
		fmt.Println("3. Search")
		fmt.Println("4. Remove")
		fmt.Println("5. Display")
		fmt.Println("0. Exit")
		fmt.Println("Enter choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			break
		case 1:
			fmt.Println("Add a new val:")
			fmt.Scanln(&newVal)
			n := node{val: newVal, next: nil}
			ls.append(&n)
		case 2:
			fmt.Println("Enter new value to be Inserted:")
			fmt.Scanln(&newVal)
			n := node{val: newVal, next: nil}
			fmt.Println("Enter a node value after which new node to be inserted.")
			var afterNodeVal int
			fmt.Scanln(&afterNodeVal)
			ls.insert(n, afterNodeVal)
		case 3:
			fmt.Println("Enter value to be searched.")
			var searchVal int
			fmt.Scanf("%d", &searchVal)
			ls.search(searchVal)
		case 4:
			fmt.Println("Enter value to be removed from list")
			var removeVal int
			fmt.Scanf("%d", &removeVal)
			ls.remove(removeVal)
		case 5:
			ls.display()
		default:
			fmt.Println("Please enter valid choice")
		}
	}
}
