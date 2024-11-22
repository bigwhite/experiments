package foo

import (
	"bisect-demo/bisect"
	"flag"
)

var (
	bisectFlag = flag.String("bisect", "", "bisect pattern")
	matcher    *bisect.Matcher
)

// Features represents different features that might cause issues
const (
	FeatureRangeIteration  = "range-iteration"  // Using range vs classic for loop
	FeatureConcurrentLogic = "concurrent-logic" // Adding concurrent modifications
)

func Init() {
	flag.Parse()
	if *bisectFlag != "" {
		matcher, _ = bisect.New(*bisectFlag)
	}
}

func ProcessItems(items []int) []int {
	result := make([]int, 0, len(items))

	// First potential problematic change: different iteration approach
	id1 := bisect.Hash(FeatureRangeIteration)
	if matcher == nil || matcher.ShouldEnable(id1) {
		if matcher != nil && matcher.ShouldReport(id1) {
			println(bisect.Marker(id1), "enabled feature:", FeatureRangeIteration)
		}
		// Potentially problematic implementation using range
		for i := range items {
			result = append(result, items[i]*2)
		}
	} else {
		// Correct implementation using value iteration
		for _, v := range items {
			result = append(result, v*2)
		}
	}

	// Second potential problematic change: concurrent modifications
	id2 := bisect.Hash(FeatureConcurrentLogic)
	if matcher == nil || matcher.ShouldEnable(id2) {
		if matcher != nil && matcher.ShouldReport(id2) {
			println(bisect.Marker(id2), "enabled feature:", FeatureConcurrentLogic)
		}
		// Potentially problematic implementation with concurrency
		for i := 0; i < len(result); i++ {
			go func(idx int) {
				result[idx] += 1 // Race condition
			}(i)
		}
	}

	return result
}
