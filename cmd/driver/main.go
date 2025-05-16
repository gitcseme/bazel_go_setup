package main

import (
	"fmt"
	"root/internal/evm"
	"root/proto"
)

func main() {
	fmt.Println("Hello, World!")

	evm.Vote("nowka")
	evm.Vote("langol")
	evm.Vote("dhan")
	evm.Vote("nowka")

	oneVote := &proto.EvmVote {
		Sign: "nowka",
		VoteNumber: 1,
	}

	fmt.Println("Votes:", evm.GetVotes(), oneVote)
}
