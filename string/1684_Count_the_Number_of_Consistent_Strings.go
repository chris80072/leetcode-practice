func countConsistentStrings(allowed string, words []string) int {
    r := len(words)
    s := make([]bool, 26)
    
    for _, v := range allowed {
        s[v - 'a'] = true
    }
    
    for _, w := range words {
        for _, v := range w {
            if !s[v - 'a'] {
                r--
                break
            }
        }
    }
    
    return r
}	