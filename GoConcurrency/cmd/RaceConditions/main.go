package main

import (
	"fmt"
	"time"
)

var data int

/*
Race Conditions :
A race condition occurs when two or more operations must execute in the correct
order, but the program has not been written so that this order is guaranteed to be
maintained.
Most of the time, this shows up in what’s called a data race, where one concurrent
operation attempts to read a variable while at some undetermined time another con‐
current operation is attempting to write to the same variable.

*/

func main() {
	go func() {
		data++
	}()
    time.Sleep(1*time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Println(data)
	}
}

/*
Have we solved our data race? No. In fact, it’s still possible for all three outcomes to
arise from this program, just increasingly unlikely. The longer we sleep in between
invoking our goroutine and checking the value of data, the closer our program gets to
achieving correctness—but this probability asymptotically approaches logical correct‐
ness; it will never be logically correct.
In addition to this, we’ve now introduced an inefficiency into our algorithm. We now
have to sleep for one second to make it more likely we won’t see our data race. If we utilized the correct tools, we might not have to wait at all, or the wait could be only a microsecond.
The takeaway here is that you should always target logical correctness. Introducing
sleeps into your code can be a handy way to debug concurrent programs, but they are
not a solution.
*/