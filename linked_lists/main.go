package main

func main() {
	linkedList := NewLinkedList()
	linkedList.Append(10)
	linkedList.Append(11)
	linkedList.Append(12)
	linkedList.Append(13)
	linkedList.Print()
	linkedList.Prepend(5)
	linkedList.Print()
	linkedList.Insert(2, 9)
	linkedList.Insert(0, 8)
	linkedList.Insert(20, 88)
	linkedList.Print()
	linkedList.Remove(0)
	linkedList.Print()
}
