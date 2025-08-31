package main

import (
	"fmt"
	"sync"
)

/*
Memory Access Synchronization :
Let’s say we have a data race: two concurrent processes are attempting to access the
same area of memory, and the way they are accessing the memory is not atomic. Our
previous example of a simple data race will do nicely with a few modifications.

We’ve added an else clause here so that regardless of the value of data we’ll always
get some output. Remember that as it is written, there is a data race and the output of
the program will be completely nondeterministic.
In fact, there’s a name for a section of your program that needs exclusive access to a
shared resource. This is called a critical section. In this example, we have three critical
sections:

• Our goroutine, which is incrementing the data variables.
• Our if statement, which checks whether the value of data is 0.
• Our fmt.Printf statement, which retrieves the value of data for output.

*/

//func main() {
//	var data int
//	go func() { data++ }()
//	if data == 0 {
//		fmt.Println("the value is 0")
//	} else {
//		fmt.Printf("the value is %v.\n", data)
//	 }
// }

/*
There are various ways to guard your program’s critical sections, and Go has some
better ideas on how to deal with this, but one way to solve this problem is to syn‐
chronize access to the memory between your critical sections.
*/

var memoryAccess sync.Mutex
var value int

func main() {
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}

/*
In this example we’ve created a convention for developers to follow. Anytime devel‐
opers want to access the data variable’s memory, they must first call Lock, and when
they’re finished they must call Unlock. Code between those two statements can then
assume it has exclusive access to data; we have successfully synchronized access to the
memory.
*/