package main

import (
	"fmt"
	"stack-and-queue-golang/core"
)

func main() {
	queue := core.NewQueue()
	stack := core.NewStack()

	queue.Add(1)
	resultQueue := queue.IsEmpty()
	fmt.Println(resultQueue)
	queue.Peek()
	queue.Remove()
	fmt.Println(resultQueue)

	stack.Push(1)
	fmt.Println(stack.Pop())
	stack.Peek()
	fmt.Println(stack.IsEmpty())

}
