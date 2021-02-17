func diStringMatch(S string) []int {
    var r []int
    b, e := 0, len(S)
    
    for _, v := range S {
        if v == 'I' {
            r = append(r, b)
            b++
        } else {
            r = append(r, e)
            e--
        }
    }
    
    r = append(r, b)
    
    return r
}