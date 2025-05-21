package evm

import (
	"testing"
)

func TestVote(t *testing.T) {
	// Test the Vote function
	Vote("nowka")
	Vote("langol")
	Vote("dhan")
	Vote("nowka")

	// Check the votes
	votes := GetVotes()

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
