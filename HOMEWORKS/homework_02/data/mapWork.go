package data

import (
	"sort"
)

// SortByValue : Method to sort a map by value
func SortByValue(wordFrequencies map[string]int) (PairList, float64) {
	var total float64
	pl := make(PairList, len(wordFrequencies))
	i := 0
	// get total count of languages
	for _, v := range wordFrequencies {
		total += float64(v)
	}

	// next loop is for creating a PairList
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pl))

	return pl, total

}

// GetTopFiveByPercent : from a Sorted PairList and total sum of all values inside
// return new PairList with 6 elements and value is converted to percentage
func GetTopFiveByPercent(sorted PairList, totalValue float64) PairList {
	count := 0
	percent := 0

	final := make(PairList, 6)
	for _, v := range sorted {
		if count == 5 {
			final[count].Key = "Other"
			final[count].Value = 100 - percent
			break
		}
		final[count] = Pair{v.Key, int((float64(v.Value) / totalValue) * 100)}
		percent += int((float64(v.Value) / totalValue) * 100)
		count++

	}
	return final
}


// GetYearsByPercent : list in ordered line last five years by distribution
func GetYearsByPercent(sorted PairList, totalValue float64) PairList {
	count := 0
	percent := 0

	final := make(PairList, 5)
	for _, v := range sorted {
		final[count] = Pair{v.Key, int((float64(v.Value) / totalValue) * 100)}
		percent += int((float64(v.Value) / totalValue) * 100)
		count++

	}
	return final
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }





