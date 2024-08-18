package algorithm

type simpleAlgoScale struct {
	FIVE_KIND  int
	FOUR_KIND  int
	FULL_HOUSE int
	THREE_KIND int
	TWO_PAIR   int
	ONE_PAIR   int
	HIGH_CARD  int
}

var SIMPLE_ALGORITHM_MAP = simpleAlgoScale{
	FIVE_KIND:  7,
	FOUR_KIND:  6,
	FULL_HOUSE: 5,
	THREE_KIND: 4,
	TWO_PAIR:   3,
	ONE_PAIR:   2,
	HIGH_CARD:  1,
}

type SimpleAlgorithm struct {
}

func NewSimpleAlgorithm() *SimpleAlgorithm {
	return &SimpleAlgorithm{}
}

func (sa *SimpleAlgorithm) Calculate(kinds []string, handIndex map[string]int) int {
	fiveKind := false
	fourKind := false
	threeKind := false
	pair := 0
	for _, kind := range kinds {
		switch handIndex[kind] {
		case 5:
			fiveKind = true
		case 4:
			fourKind = true
		case 3:
			threeKind = true
		case 2:
			pair += 1
		}
	}

	if fiveKind {
		return SIMPLE_ALGORITHM_MAP.FIVE_KIND
	}

	if fourKind {
		return SIMPLE_ALGORITHM_MAP.FOUR_KIND
	}

	if threeKind && pair == 1 {
		return SIMPLE_ALGORITHM_MAP.FULL_HOUSE
	}

	if threeKind {
		return SIMPLE_ALGORITHM_MAP.THREE_KIND
	}

	if pair == 2 {
		return SIMPLE_ALGORITHM_MAP.TWO_PAIR
	}

	if pair == 1 {
		return SIMPLE_ALGORITHM_MAP.ONE_PAIR

	}

	return 1
}

func (sa *SimpleAlgorithm) GetValue(letter string) int {
	switch letter {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	}

	return 0
}
