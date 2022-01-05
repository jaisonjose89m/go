package main

import (
	"fmt"
	"github.com/jaisonjose89m/go/learngo/collections"
	"github.com/jaisonjose89m/go/learngo/flow"
	"github.com/jaisonjose89m/go/learngo/primitives"
)

type Animal struct {
	Name string
	Kind string
}

type Cat struct {
	Animal
	Domesticated bool
}

func main() {
	fmt.Println("Hello World !!!")
	cat := new(Cat)
	cat.Name = "Tom"
	cat.Kind = "Mammal"
	cat.Domesticated = true
	//fmt.Println(cat)

	list := []string{"abc", "cde"}
	formattedString := fmt.Sprintf("Sample list is %v", list)
	fmt.Println(formattedString)
	//workingWithPrimitives()
	//collection()
	//functions.InvokeFunc()
	//loops()
	//channels.UnbufferedChannelDemo()
	//channels.BufferedChannelDemo()
}

func loops() {
	flow.SimpleFor()
	flow.ClassicFor()
	flow.InfiniteLoop()
	flow.LoopOverSlice()
	flow.LoopOverMap()
}

func collection() {
	collections.Array()
	collections.Slice()
	collections.Map()
	collections.Struct()
}

func workingWithPrimitives() {
	primitives.Primitive()
	primitives.Pointer()
	primitives.Constant()
}
