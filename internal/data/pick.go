package data 

import(
	"os"
	"strings"
	"math/rand"
)

type PassageSet struct {
	Passages []string
	Current int
}

func RandomParagraphPick() PassageSet {
	contents, _ := os.ReadFile("internal/data/english.txt")
	
	passages := strings.Split(string(contents), "===")

	for i := range passages {
		passages[i] = strings.TrimSpace(passages[i])
	}

	index := rand.Intn(len(passages))
	
	passageset := PassageSet{
		Passages: passages,
		Current: index,
	}

	return passageset
}
