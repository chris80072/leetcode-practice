func smallerNumbersThanCurrent(nums []int) []int {
    var r []int
    m := make(map[int]int)
    
    tmpNums := make([]int, len(nums))
    copy(tmpNums, nums)
    sort.Ints(tmpNums)
    
    for i, n := range tmpNums {
        if _, ok := m[n]; !ok {
            m[n] = i
        }
    }
    
    for _, n := range nums {
        r = append(r, m[n])
    }
    
    return r
}