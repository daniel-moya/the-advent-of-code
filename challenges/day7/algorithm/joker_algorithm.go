package algorithm

type jokerScale struct {
	FIVE_KIND  int
	FOUR_KIND  int
	FULL_HOUSE int
	THREE_KIND int
	TWO_PAIR   int
	ONE_PAIR   int
	HIGH_CARD  int
}

var JOKER_MAP = jokerScale{
	FIVE_KIND:  7,
	FOUR_KIND:  6,
	FULL_HOUSE: 5,
	THREE_KIND: 4,
	TWO_PAIR:   3,
	ONE_PAIR:   2,
	HIGH_CARD:  1,
}

type JokerAlgorithm struct {
}

func NewJokerAlgorithm() *JokerAlgorithm {
	return &JokerAlgorithm{}
}

func (ja *JokerAlgorithm) Calculate(kinds []string, handIndex map[string]int) int {
	fiveKind := false
	fourKind := false
	threeKind := false
	pair := 0
	jCount := 0
	for _, kind := range kinds {
		if kind == "J" {
			jCount += handIndex[kind]
			continue
		}
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

	if jCount == 5 || jCount == 4 {
		return JOKER_MAP.FIVE_KIND
	}

	if fiveKind {
		return JOKER_MAP.FIVE_KIND
	}

	if fourKind {
		if jCount > 0 {
			return JOKER_MAP.FIVE_KIND

		}
		return JOKER_MAP.FOUR_KIND
	}

	if threeKind {
		if pair == 1 {
			return JOKER_MAP.FULL_HOUSE
		}
		if jCount == 1 {
			return JOKER_MAP.FOUR_KIND
		}
		if jCount == 2 {
			return JOKER_MAP.FIVE_KIND
		}

		return JOKER_MAP.THREE_KIND

	}

	if pair == 2 {
		if jCount == 1 {
			return JOKER_MAP.FULL_HOUSE
		}
		return JOKER_MAP.TWO_PAIR
	}

	if pair == 1 {
		if jCount == 1 {
			return JOKER_MAP.THREE_KIND
		}
		if jCount == 2 {
			return JOKER_MAP.FOUR_KIND
		}
		if jCount == 3 {
			return JOKER_MAP.FIVE_KIND
		}

		return JOKER_MAP.ONE_PAIR

	}

	if jCount == 1 {
		return JOKER_MAP.ONE_PAIR
	}
	if jCount == 2 {
		return JOKER_MAP.THREE_KIND
	}
	if jCount == 3 {
		return JOKER_MAP.FOUR_KIND
	}

	return 1
}

func (ja *JokerAlgorithm) GetValue(letter string) int {
	switch letter {
	case "A":
		return 15
	case "K":
		return 14
	case "Q":
		return 13
	case "T":
		return 11

	}
	// J will be 0 in joker's algo
	return 0
}
