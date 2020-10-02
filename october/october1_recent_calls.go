package october

type RecentCounter struct {
	pingTimerArr []int
}

func Constructor1() RecentCounter {
	pingTimerArray := []int{}
	return RecentCounter{pingTimerArray}
}

func (this *RecentCounter) Ping(t int) int {
	result := 0
	this.pingTimerArr = append(this.pingTimerArr, t)
	lowLimit := t - 3000
	for i := 0; i < len(this.pingTimerArr); i++ {
		if this.pingTimerArr[i] >= lowLimit {
			result++
		}
	}
	return result
}
