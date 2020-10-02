package october

type RecentCounter struct {
	times []int
}

func Constructor1() RecentCounter {
	return RecentCounter{
		times: make([]int, 0, 10000),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.times = append(rc.times, t)

	for len(rc.times) > 0 && rc.times[0]+3000 < t {
		rc.times = rc.times[1:]
	}

	return len(rc.times)
}
