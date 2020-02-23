package bytez

func Index(p []byte, toSearch byte) int {
	for i, bt := range p {
		if bt == toSearch {
			return i
		}
	}
	return -1
}
