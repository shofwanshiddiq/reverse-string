package main

func main() {

	ResultReverse := ReturnReverseString("Test Reverse String")
	println(ResultReverse)
}

func ReturnReverseString(str string) string {
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}
