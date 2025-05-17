package evm

import (
	"testing"
	"root/internal/evm"
)

func TestVote(t *testing.T) {
	// Test the Vote function
	evm.Vote("nowka")
	evm.Vote("langol")
	evm.Vote("dhan")
	evm.Vote("nowka")

	// Check the votes
	votes := evm.GetVotes()

	if votes["nowka"] != 2 {
		t.Errorf("Expected 2 votes for nowka, got %d", votes["nowka"])
	}
	if votes["langol"] != 1 {
		t.Errorf("Expected 1 vote for langol, got %d", votes["langol"])
	}
	if votes["dhan"] != 1 {
		t.Errorf("Expected 1 vote for dhan, got %d", votes["dhan"])
	}
}