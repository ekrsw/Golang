package alib

func Average(s []int) int {
	total := 0
	for _, n := range s {
		total += n
	}
	return int(total / len(s))
}