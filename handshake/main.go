package handshake

func GetParticipants(h int) int {

	if h == 0 {
		return 0
	}
	if h == 1 {
		return 2
	}

	for i := 2; ; i++ {
		p := (i * (i - 1)) / 2
		if p >= h {
			return i
		}
	}
}
