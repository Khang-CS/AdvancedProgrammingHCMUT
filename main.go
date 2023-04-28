package main

import (
	"fmt"
	f "golang/flyweight"
	"strconv"
)

func main() {
	factory := f.NewFactory()

	for i := 1; i <= 15; i++ {
		starContext := f.NewContext(strconv.Itoa(i), 2)
		soldier := factory.CreateSoldier("Army")
		soldier.Promote(starContext)

	}

	for i := 16; i <= 35; i++ {
		starContext := f.NewContext(strconv.Itoa(i), 2)
		soldier := factory.CreateSoldier("Coast Guard")
		soldier.Promote(starContext)
	}
	fmt.Printf("Number of soldier type in map: %d\n", factory.GetSize())
}
