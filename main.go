package main

import (
	"fmt"
	"runtime/trace"
)

type The struct {
	Best string
}
type Are struct {
	The The
}
type We struct {
	Are Are
}

// ----------------
type Hello struct {
	World string
}
//-----------------



func main() {
	We := We{
		Are: Are{
			The: The{
				Best: "KODA",
			},
		},
	}
	fmt.Println("1. We.Are.The.Best : ", We.Are.The.Best)

	hello := Hello{
		World: "Hello World",
	}
	fmt.Println("2. hello.world : ", hello.World)
}
