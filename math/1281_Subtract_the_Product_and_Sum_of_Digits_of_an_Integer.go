func subtractProductAndSum(n int) int {
    p, s := 1, 0
    
    for n > 0 {
        t := n % 10
        n = n / 10
        p *= t
        s += t
    }
    
    return p - s
}