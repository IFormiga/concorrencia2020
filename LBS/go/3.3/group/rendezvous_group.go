package main

/*
3.3 Rendezvous

Puzzle: Generalize the signal pattern so that it works both ways.
Thread A has to wait for Thread B and vice versa.

In other words, given this code...

  Thread A
 ┌────────────────
1│ statement a1
2│ statement a2

  Thread B
 ┌────────────────
1│ statement b1
2│ statement b2

...we want to guarantee that `a1` happens before `b2` and `b1` happens before `a2`.
In writing your solution, be sure to specify the names and initial values of your semaphores
(little hint there).

Your solution should not enforce too many constraints.
For example, we don’t care about the order of a1 and b1.
In your solution, either order should be possible.

This synchronization problem has a name; it’s a rendezvous.
The idea is that two threads rendezvous at a point of execution,
and neither is allowed to proceed until both have arrived.
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func show(name, activity string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("%s:\t%s\n", name, activity)
}

func aTask(randezvous *sync.WaitGroup) {
	defer wg.Done()
	show("A", "doing a1")
	a1Done = true
	randezvous.Done()
	randezvous.Wait()
	if !(a1Done || b1Done) {
		panic("a1Done or b1Done is false")
	}
	show("A", "doing a2")
}

func bTask(randezvous *sync.WaitGroup) {
	defer wg.Done()
	show("B", "doing b1")
	b1Done = true
	randezvous.Done()
	randezvous.Wait()
	if !(a1Done || b1Done) {
		panic("b1Done or a1Done is false")
	}
	show("B", "doing b2")
}

var wg sync.WaitGroup
var a1Done, b1Done bool

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var randezvous sync.WaitGroup
	wg.Add(2)
	randezvous.Add(2)
	go aTask(&randezvous)
	go bTask(&randezvous)
	wg.Wait()
}