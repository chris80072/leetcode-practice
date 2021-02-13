func maxDepth(s string) int {
    r, t := 0, 0
    
    for _, v := range s {
        if v == '(' {
            t++
        } else if v == ')' {
            t--
        }
        
        if t > r {
            r = t
        }
    }
    
    return r
}