package algorithm

type Algorithm interface {
	Calculate(kinds []string, handIndex map[string]int) int
	GetValue(char string) int
}
