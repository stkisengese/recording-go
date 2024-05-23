package piscine

func IsPrime(nb int) bool {
	for i := 3; i*i <= nb; i += 2 {
		switch {
		case nb < 2:
			return false
		case nb == 2, nb == 3:
			return true
		case nb%2 == 0:
			return false
		case nb%i == 0:
			return false
		}
	}
	return true
}
