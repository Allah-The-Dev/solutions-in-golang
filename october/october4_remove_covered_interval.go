package october

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	res, ending := 0, 0
	for _, interval := range intervals {
		if interval[1] > ending {
			res++
			ending = interval[1]
		}
	}

	return res
}
