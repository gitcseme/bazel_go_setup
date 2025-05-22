package evm

import (
	epb "bazel_go_setup/proto"
	"fmt"
	"os"

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

func StoreVote() {
	var evmVoteList []*epb.EvmVote
	for k, v := range votes {
		evmVoteList = append(evmVoteList, &epb.EvmVote{
			Sign:       k,
			VoteNumber: int32(v),
		})
	}

	file, err := os.Create("votes.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	voteList := &epb.EvmVoteList{
		Year:  2025,
		Votes: evmVoteList,
	}

	binaryData, err := gpb.Marshal(voteList)
	if err != nil {
		fmt.Println("Error marshalling votes:", err)
		return
	}

	_, err = file.Write(binaryData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Marshalled votes:", binaryData)
}

func LoadVote() {
	file, err := os.Open("votes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := os.ReadFile("votes.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	voteList := &epb.EvmVoteList{}
	err = gpb.Unmarshal(data, voteList)
	if err != nil {
		fmt.Println("Error unmarshalling votes:", err)
		return
	}

	for _, v := range voteList.Votes {
		votes[v.Sign] = int(v.VoteNumber)
	}

	fmt.Println("Unmarshalled votes:", votes)
}

func Vote(sign string) {
	votes[sign]++
}

func GetVotes() map[string]int {
	return votes
}
