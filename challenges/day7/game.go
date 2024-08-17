package day7

type Ranking struct {
	Index []Hand
	Size  int
}

func NewSortedRanking(hands []Hand) *Ranking {
	sorted := mergeSort(hands)

	return &Ranking{Index: sorted, Size: len(sorted)}

}

func CompareHands(a, b Hand) int {
	if (a.Value - b.Value) != 0 {
		return a.Value - b.Value
	}

	for index := 0; index < 5; index++ {
		compared := CompareCards(a.Index[index], b.Index[index])
		if compared != 0 {
			return compared
		}
	}

	// If all 5 cards are equal in the same order
	return 0
}

func (r *Ranking) GetBidTotal() int {
	total := 0
	for index, hand := range r.Index {
		total += hand.Bid * (index + 1)
	}
	return total
}

func mergeSort(items []Hand) []Hand {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []Hand, b []Hand) []Hand {
	final := []Hand{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		result := CompareHands(a[i], b[j])
		if result <= 0 {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
