package septemberleetcodechallange

import "strconv"

func getHint(secret string, guess string) string {

	bullCount, cowCount := 0, 0

	bucket := [10]int{}

	for i, ch := range secret {
		guessChAti := rune(guess[i])
		if guessChAti == ch {
			bullCount++
		} else {
			bucket[ch-'0']++
			bucket[guessChAti-'0']--
		}
	}

	countOfSecretNotFoundInGuess := 0
	for _, num := range bucket {
		if num > 0 {
			countOfSecretNotFoundInGuess += num
		}
	}

	cowCount = len(secret) - bullCount - countOfSecretNotFoundInGuess

	return strconv.Itoa(bullCount) + "A" + strconv.Itoa(cowCount) + "B"
}
