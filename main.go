package main

func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
func main() {
	//fmt.Println(Hello("world", "English"))
}
