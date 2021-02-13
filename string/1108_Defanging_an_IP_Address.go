func defangIPaddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}

func defangIPaddr(address string) string {
	r := ""
	
	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			r += "[.]"
		} else {
			r += string(address[i])
		}
	}
	
	return r
}