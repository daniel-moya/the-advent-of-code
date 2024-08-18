package ranking

import (
	"github.com/daniel-moya/aoc/challenges/day7/hand"
)

type Ranking struct {
	Index []hand.Hand
	Size  int
}

func NewSortedRanking(hands []hand.Hand) *Ranking {
	sorted := mergeSort(hands)

	return &Ranking{Index: sorted, Size: len(sorted)}

}

func (r *Ranking) GetBidTotal() int {
	total := 0
	for index, hand := range r.Index {
		total += hand.Bid * (index + 1)
	}
	return total
}

func mergeSort(items []hand.Hand) []hand.Hand {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []hand.Hand, b []hand.Hand) []hand.Hand {
	final := []hand.Hand{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		result := hand.CompareHands(a[i], b[j])
		if result < 0 {
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
