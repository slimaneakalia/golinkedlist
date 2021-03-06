package golinkedlist

import "fmt"

type LinkedList struct{
	Value interface{}
	Next *LinkedList
}

type EqualityComparatorFunc func(interface{}, interface{}) bool

func AddToEndOrCreateList(targetList *LinkedList, value interface{}) *LinkedList{
	newNode := &LinkedList{
		Value: value,
		Next: nil,
	}

	if targetList == nil {
		return newNode
	}
	
	tmpPointer := targetList
	for tmpPointer.Next != nil {
		tmpPointer = tmpPointer.Next
	}

	tmpPointer.Next = newNode
	return targetList
}

func (head *LinkedList) AddValue(v interface{}) *LinkedList{
	return &LinkedList{
		Value: v,
		Next: head,
	}
}

func DefaultEqualityComparator(a interface{}, b interface{}) bool {
	return a == b
}


func (head *LinkedList) FindValue(value interface{},
	equalityComparator EqualityComparatorFunc) *LinkedList{
	tmp := head
	eqCmp := equalityComparator
	if equalityComparator == nil {
		eqCmp = DefaultEqualityComparator
	}

	for tmp != nil {
		if eqCmp(tmp.Value, value) {
			return tmp
		}

		tmp = tmp.Next
	}

	return nil
}

func (head *LinkedList) Print() string{
	str := ""
	tmp := head
	for tmp != nil {
		str += fmt.Sprintf("%v - ", tmp.Value)
		tmp = tmp.Next
	}

	return str
}