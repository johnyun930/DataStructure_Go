package main

import (
	"dataStructure"
)

func main() {
	linkedList := dataStructure.MakeLinkedList()

	for i := 0; i < 10; i++ {
		linkedList.Add(dataStructure.Data{Val: i})
	}
	linkedList.PrintData()

}
