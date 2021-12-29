package main

import (
	"fmt"
	"github.com/jaisonjose89m/go/learngo/collections"
	"github.com/jaisonjose89m/go/learngo/flow"
	"github.com/jaisonjose89m/go/learngo/functions"
	"github.com/jaisonjose89m/go/learngo/primitives"
)

func main() {
	fmt.Println("Hello World !!!")
	workingWithPrimitives()
	collection()
	functions.InvokeFunc()
	loops()
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
