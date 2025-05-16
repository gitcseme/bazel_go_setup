package evm

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
}

func GetVotes() map[string]int {
	return votes
}
