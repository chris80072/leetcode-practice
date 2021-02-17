func minCostToMoveChips(position []int) int {
    e, o := 0, 0
    
    for _, v := range position {
        if v % 2 == 0 {
            e++
        } else {
            o++
        }
    }
    
    if e < o {
        return e
    }
    
    return o
}