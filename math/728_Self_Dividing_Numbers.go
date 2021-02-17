func selfDividingNumbers(left int, right int) []int {
    var r []int
    
    for i := left; i <= right; i++ {
        var t int
        
        for t = i; t > 0; t = t/10 {
            s := t % 10
            if s == 0 || i % s != 0 {
                break
            }
        }
        
        if t == 0 {
            r = append(r, i)
        }
    }
    return r
}