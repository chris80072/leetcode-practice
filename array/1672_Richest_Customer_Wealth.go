func maximumWealth(accounts [][]int) int {
    r := 0
    
    for _, a := range accounts {
        s := 0
        
        for _, w := range a {
            s += w
        }
        
        if s > r {
            r = s
        }
    }   
    
    return r
}