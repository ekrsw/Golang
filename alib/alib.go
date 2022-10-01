package alib

func Average(sl []int) int {
	total := 0
	for _, u := range sl {
		total += u
	}
	return int(total / len(sl))
}
