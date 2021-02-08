func createTargetArray(nums []int, index []int) []int {
    var r []int
    
    for i:= 0; i < len(nums); i++ {
        r = append(r[:index[i]], append([]int{nums[i]}, r[index[i]:]...)...)
    }
    
    return r
}