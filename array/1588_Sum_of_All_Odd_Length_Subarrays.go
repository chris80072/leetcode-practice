func sumOddLengthSubarrays(arr []int) int {
    r := 0
    
    for i := 0; i < len(arr); i++ {
        // 條件為連續的奇數，故除了第一個元素，之後的都會兩個兩個一組被計算
        m := (len(arr) - i - 1) / 2 + 1
        
        for j := i; j < len(arr); j += 2 {
            if j == i {
                r += arr[j] * m
            } else {
                r += (arr[j-1] + arr[j]) * m
            }
            m--
        } 
    }
    
    return r
}

func sumOddLengthSubarrays(arr []int) int {
    r := 0
    
    for b, e := 1, len(arr); b -1 < len(arr); b, e = b + 1, e - 1 {
        r += (b * e + 1) / 2 * arr[b - 1]
    }
    
    return r
}