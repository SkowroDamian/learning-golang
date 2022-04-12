package iteration

func Repeat(character string, number int) string {
	repeated := ""

	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}
