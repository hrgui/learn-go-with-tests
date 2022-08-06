package main

func Repeat(s string, max int) string {
	repeated := ""

	for i := 0; i < max; i++ {
		repeated += s
	}

	return repeated
}
