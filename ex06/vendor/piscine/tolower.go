package piscine

func ToUpper(s string) string {
	var result string
	for _, c := range s {
		if (c >= 'A' && c <= 'Z') {
			result += string(c+'a'-'A')
		} else {
			result += string(c)
		}
	}
	return string(result)
}
