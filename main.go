package main

import (
	"fmt"
	"math"
	"sync"
)

// Current playing with 8 cores
type Core struct {
	Value int
	Counter int
	Trigger int
}


var (
	cores = []Core{
		{Value: 8, Counter: 0, Trigger: -1},
		{Value: 19, Counter: 0, Trigger: 0},
		{Value: 7, Counter: 0, Trigger: 1},
		{Value: 15, Counter: 0, Trigger: 0},
		{Value: 7, Counter: 0, Trigger: 2},
		{Value: 13, Counter: 0, Trigger: 0},
		{Value: 12, Counter: 0, Trigger: 1},
		{Value: 14, Counter: 0, Trigger: 0},
	}
)

func main() {
	var wg sync.WaitGroup

	wg.Add(len(cores))

	for i := range cores {
		go (&cores[i]).start(i, &wg)
	}

	wg.Wait()

	debugPrint(0)
	debugPrint(1)
	debugPrint(2)
	debugPrint(3)
	debugPrint(4)
	debugPrint(5)
	debugPrint(6)
	debugPrint(7)

	fmt.Println("Global Sum", cores[0].Value)
}

func (core *Core) start(index int, wg *sync.WaitGroup) {
	iteration := 0
	if index % 2 == 1 {
		prevIndex := index - int(math.Pow(2, float64(iteration)))
		cores[prevIndex].receive(prevIndex, iteration, core.Value)
	}

	wg.Done()
}

func (core *Core) receive(index, iteration, value int) {
	iteration++
	core.Counter++
	core.Value += value
	
	fmt.Println("core", index, "receiving", value, "trigger", core.Trigger)

	if core.Trigger == core.Counter {
		prevIndex := index - int(math.Pow(2, float64(iteration)))
		if prevIndex > -1 {
			cores[prevIndex].receive(prevIndex, iteration, core.Value)
		}
	}
}

func debugPrint(index int)  {
	fmt.Println("index", index, "Value", cores[index].Value, "Counter", cores[index].Counter, "Trigger", cores[index].Trigger)
}