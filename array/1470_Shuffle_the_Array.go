func shuffle(nums []int, n int) []int {
    var r []int
    
    for i := 0; i < n; i++ {
        r = append(r, nums[i], nums[i+n])
    }
    
    return r
}