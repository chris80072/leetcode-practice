func decompressRLElist(nums []int) []int {
    var r []int
    
    for i := 0; i < len(nums) - 1; i += 2 {
        for j := 0; j < nums[i]; j++ {
            r = append(r, nums[i+1])
        }
    }
    
    return r
}