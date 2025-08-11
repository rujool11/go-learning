package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var db = []string{"val1", "val2", "val3", "val4", "val5"}
var wg = sync.WaitGroup{}
var results = []string{}
var m = sync.Mutex{}

func main() {

	// goroutines is a way of launching multiple functions and have them execute concurrently
	// note that concurrency is not the same as parallel execution
	// concurrency => multiple tasks in progress at the same time, can jump back and forth between tasks (like context switching in OS)
	// working on t1 and t2, can move to t2 while t1 is processing, and then move back to t1, etc
	// parallelism is one way to achieve concurrency
	// in this multiple CPUs/cores are working on different tasks
	// with goroutines in practise, you do achieve some level of paralleism assuming you have a multi core CPU:w

	t_init := time.Now()
	for i := 0; i < len(db); i++ {
		// dbCall(i) // waits for each call to finish, and then moves on
		// go dbCall(i) // does not wait for completion of function call, and moves on
		// this will just exit and print nothing as this doesnt wait for other calls

		// to wait until tasks have completed and then continue on, we need wait groups
		// we can import these from the sync package

		// wait groups are basically just counters
		// add 1 to the counter at each call, and in dbCall, use .Done() to decrement counter by 1
		// and before Printf statement, just add .Wait()

		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()
	fmt.Printf("total execution time: %v\n", time.Since(t_init))
	fmt.Printf("results: %v\n", results)

	// when we have multiple concurrent thread manipulating same memory, it can
	// lead to unexpected results and corruption of memory
	// to control this, use a mutex

}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	m.Lock()
	results = append(results, db[i])
	m.Unlock()
	wg.Done()
	// when a goroutine reaches a mutex, it will lock it and other goroutines will have to wait for it
	// leads to better concurrency
	// go also offers RWMutex, which offers RLock, RUnlock, WLock, WUnlock (read/write lock/unlock)
}
