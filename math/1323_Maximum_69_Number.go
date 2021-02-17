func maximum69Number (num int) int {
    s := strconv.Itoa(num)
    p := 0
    
    for i, v := range s {
        if v == '6' {
            p = 3 * int(math.Pow(10, float64(len(s) - i -1 )))
            break
        }
    }
    
    return num + p
}