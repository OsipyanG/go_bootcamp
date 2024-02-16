package stats

import "math"

func (st Statistics) Mean() float64 {
	if st.numbersLen == 0 {
		return math.NaN()
	}
	sum := 0.0
	for _, num := range st.numbers {
		sum += float64(num)
	}
	return sum / float64(st.numbersLen)
}
func (st Statistics) Mode() int {
	var numberCounter = make(map[int]int, st.numbersLen)
	maxValue := 0
	var maxElem = 0
	for _, number := range st.numbers {
		numberCounter[number]++
	}
	for key, value := range numberCounter {
		if value > maxValue {
			maxValue = value
			maxElem = key
		}
	}
	return maxElem
}
func (st Statistics) Median() float64 {
	if st.numbersLen == 0 {
		return 0
	}
	mid := st.numbersLen / 2
	if st.numbersLen%2 != 0 {
		return float64(st.numbers[mid])
	}
	return float64(st.numbers[mid-1]+st.numbers[mid]) / 2
}
func (st Statistics) StandardDeviation() float64 {
	if st.numbersLen == 0 {
		return math.NaN()
	}
	average := st.Mean()
	sum := 0.0

	for _, num := range st.numbers {
		sum += float64((float64(num) - average) * (float64(num) - average))

	}
	return math.Sqrt(sum / float64(st.numbersLen))
}
