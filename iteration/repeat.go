package iteration

func Repeat(char string, nums int) string {
	var repeated string
	for i := 0; i < nums; i++ {
		repeated += char
	}
	return repeated
}