func kidsWithCandies(candies []int, extraCandies int) []bool {
    m := 0
    for _, c := range candies {
        if c > m {
            m = c
        }
    }
    
    var r []bool
    
    for _, c := range candies {
        r = append(r, (c + extraCandies >= m))
    }
    
    return r
}