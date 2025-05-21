package evm

import (
	epb "bazel_go_setup/proto"
	"fmt"

	gpb "google.golang.org/protobuf/proto"
)

var (
	nowka  string = "nowka"
	langol string = "langol"
	dhan   string = "dhan"

	votes map[string]int = map[string]int{
		nowka:  0,
		langol: 0,
		dhan:   0,
	}
)

func Vote(sign string) {
	votes[sign]++

	vte := &epb.EvmVote{
		Sign:       sign,
		VoteNumber: 1,
	}

	data, err := gpb.Marshal(vte)
	if err != nil {
		fmt.Println("Error marshalling vote data:", err)
		return
	}

	fmt.Println("Vote data:", data)
}

func GetVotes() map[string]int {
	return votes
}
