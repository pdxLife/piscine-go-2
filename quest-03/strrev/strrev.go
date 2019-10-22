package piscine

func StrRev(s string) string {
	var r string
	b := 0
	for range s {
		b++
	}
	for i := range s {
		r = string(s[i]) + r
	}
	return r
}
