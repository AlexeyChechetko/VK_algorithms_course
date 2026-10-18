package main

import "fmt"

// Node — узел односвязного списка.
type Node struct {
	Value int
	Next  *Node
}

// LinkedList хранит указатели на первый и последний узлы.
// Пустой список готов к использованию без дополнительной инициализации.
type LinkedList struct {
	head *Node
	tail *Node
}

// InsertFirst вставляет значение в начало за O(1).
func (list *LinkedList) InsertFirst(value int) {
	node := &Node{Value: value, Next: list.head}
	list.head = node
	if list.tail == nil {
		list.tail = node
	}
}

// InsertAfter вставляет value после первого узла со значением target.
// Возвращает false, если target не найден. Поиск занимает O(n).
func (list *LinkedList) InsertAfter(target, value int) bool {
	for current := list.head; current != nil; current = current.Next {
		if current.Value == target {
			node := &Node{Value: value, Next: current.Next}
			current.Next = node
			if list.tail == current {
				list.tail = node
			}
			return true
		}
	}
	return false
}

// InsertLast вставляет значение в конец за O(1) благодаря указателю tail.
func (list *LinkedList) InsertLast(value int) {
	node := &Node{Value: value}
	if list.tail == nil {
		list.head = node
	} else {
		list.tail.Next = node
	}
	list.tail = node
}

// Print выводит значения в порядке обхода списка.
func (list *LinkedList) Print() {
	for current := list.head; current != nil; current = current.Next {
		fmt.Printf("%d -> ", current.Value)
	}
	fmt.Println("nil")
}

func main() {
	var list LinkedList
	fmt.Print("Пустой список: ")
	list.Print()

	list.InsertFirst(20)
	list.InsertFirst(10)
	fmt.Print("После вставок в начало: ")
	list.Print() // 10 -> 20 -> nil

	list.InsertLast(30)
	fmt.Print("После вставки в конец: ")
	list.Print() // 10 -> 20 -> 30 -> nil

	if list.InsertAfter(20, 25) {
		fmt.Print("После вставки 25 после 20: ")
		list.Print() // 10 -> 20 -> 25 -> 30 -> nil
	}

	// Вставка после последнего узла также обновляет tail.
	list.InsertAfter(30, 40)
	list.InsertLast(50)
	fmt.Print("После вставок в хвост: ")
	list.Print() // 10 -> 20 -> 25 -> 30 -> 40 -> 50 -> nil

	if !list.InsertAfter(100, 200) {
		fmt.Println("Число 100 не найдено, список не изменён.")
	}
}
