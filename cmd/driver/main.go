package main

import (
	"fmt"
	"os"

	"bazel_go_setup/internal/evm"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
	} else {
		fmt.Println("Current working directory:", dir)
	}

	evm.Vote("nowka")
	evm.Vote("langol")
	evm.Vote("dhan")
	evm.Vote("nowka")

	fmt.Println("Votes:", evm.GetVotes())

	evm.StoreVote()
	evm.LoadVote()

}
