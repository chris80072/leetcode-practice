func balancedStringSplit(s string) int {
    r, t := 0, 0
    
    for _, v := range s {
        if v == 'R' {
            t++
        } else {
            t --
        }
        
        if(t == 0) {
            r++
        }
    }
    return r
}