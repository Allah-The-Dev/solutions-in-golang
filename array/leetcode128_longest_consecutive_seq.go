func longestConsecutive(nums []int) int {
    // idea is that store elements in a hashset
    // for every no keep checking if one less that that do exist in set
    // if it do not exist then do a loop from there
    // this will make sure that for a single increasing number sequence, only one iteration is done
    res := 0
    
    hashSet := make(map[int]bool)
    
    for _, num := range nums { hashSet[num] = true }
    
    for _, num := range nums {
        if _, ok := hashSet[num-1]; !ok {
            currRes := 1
            num++
            for hashSet[num] {
                currRes++
                num++
            }
            if currRes > res { res = currRes }
        }
    }
    return res
}