/*
	This file is part of BlockyUtil project. It is licensed under The MIT License.

	The MIT License (MIT)
	Copyright (c) 2024 BlockyDeer <blockydeer@outlook.com>
	Permission is hereby granted, free of charge, to any person obtaining a copy of this
	software and associated documentation files (the “Software”), to deal in the Software
	without restriction, including without limitation the rights to use, copy, modify,
	merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
	permit persons to whom the Software is furnished to do so, subject to the following
	conditions:
	The above copyright notice and this permission notice shall be included in all copies
	or substantial portions of the Software.
	The Software is provided “as is”, without warranty of any kind, express or implied,
	including but not limited to the warranties of merchantability, fitness for a
	particular purpose and noninfringement. In no event shall the authors or copyright
	holders be liable for any claim, damages or other liability, whether in an action of
	contract, tort or otherwise, arising from, out of or in connection with the software
	or the use or other dealings in the Software.
*/
package blockyutil

type List[DataType any] struct {
	head *listNode[DataType]
	tail *listNode[DataType]
	len  uint
}

func CreateLinkedList[T any]() List[T] {
	return List[T]{
		head: nil,
		tail: nil,
		len:  0,
	}
}

func (list *List[T]) PushBack(data T) {
	node := newListNode(data)
	if list.IsEmpty() {
		list.head = node
		list.tail = node
		list.len = 1
	} else {
		list.tail.next = node
		node.pre = list.tail
		list.len++
	}
}

func (list *List[T]) PushFront(data T) {
	node := newListNode(data)
	if list.IsEmpty() {
		list.head = node
		list.tail = node
		list.len = 1
	} else {
		node.next = list.head
		list.head.pre = node
		list.len++
	}
}

func (list *List[DataType]) IsEmpty() bool {
	return list.head == nil
}

func (list *List[DataType]) Front() ListIterator[DataType] {
	return createListIterator[DataType](list, list.head)
}

// return the next of the last node of the linked list
func (list *List[DataType]) End() ListIterator[DataType] {
	return createListIterator[DataType](list, list.tail.next)
}

func (list *List[DataType]) Tail() ListIterator[DataType] {
	iter := list.End()
	return iter.Next()
}

func (list *List[T]) Len() uint {
	return list.len
}

type listNode[T any] struct {
	pre  *listNode[T]
	next *listNode[T]
	data T
}

func newListNode[T any](data T) *listNode[T] {
	return &listNode[T]{
		data: data,
	}
}

type ListIterator[DataType any] struct {
	list *List[DataType]
	nodePointer *listNode[DataType]
}

func createListIterator[DataType any](list *List[DataType], node *listNode[DataType]) ListIterator[DataType] {
	return ListIterator[DataType] {
		list: list,
		nodePointer: node,
	}
}

func (iter *ListIterator[DataType]) GetList() List[DataType] {
	return *(iter.list)
}

func (iter *ListIterator[DataType]) Next() ListIterator[DataType] {
	return createListIterator[DataType](iter.list, iter.nodePointer.next)
}

func (iter *ListIterator[DataType]) Prev() ListIterator[DataType] {
	return createListIterator[DataType](iter.list, iter.nodePointer.pre)
}
