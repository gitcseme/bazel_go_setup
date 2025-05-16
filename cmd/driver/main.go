package main

import (
	"fmt"
	"root/internal/evm"
)

func main() {
	fmt.Println("Hello, World!")

	evm.Vote("nowka")
	evm.Vote("langol")
	evm.Vote("dhan")
	evm.Vote("nowka")

	fmt.Println("Votes:", evm.GetVotes())
}
