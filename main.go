package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func forEach[T any](list *List[T], f func(index int, item *T)) {
	for i, item := 0, list; item != nil; i++ {
		f(i, &item.val)
		item = item.next
		if item == list {
			break
		}
	}
}

func addAfter[T any](list *List[T], val T) *List[T] {
	item := &List[T]{val: val}
	if list == nil {
		item.next = item
	} else {
		list.next, item.next = item, list.next
	}
	return item
}

func removeAfter[T any](list *List[T]) *List[T] {
	if list == nil {
		return nil
	}
	if list == list.next || list.next == nil {
		list.next = nil
		return nil
	}
	list.next, list.next.next = list.next.next, nil
	return list
}

func printList[T any](list *List[T]) {
	forEach(list, func(index int, item *T) {
		fmt.Printf("item[%d]=%v\n", index, *item)
	})
}

func main() {
	var list *List[int]
	for i, last := 0, list; i < 10; i++ {
		last = addAfter(last, i)
		if list == nil {
			list = last
		}
	}
	fmt.Println("Created list:")
	printList(list)
	forEach(list, func(index int, item *int) {
		*item += index
	})
	fmt.Println("Processed list:")
	printList(list)
	list = removeAfter(list)
	list = removeAfter(list)
	fmt.Println("Reduced list:")
	printList(list)
	for list != nil {
		list = removeAfter(list)
	}
	fmt.Println("Empty list:")
	printList(list)
}
