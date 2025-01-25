# Golang Reverse String Function

Function for reverse a string in golang

* Function Script
```golang
func ReturnReverseString(str string) string {
	reverse := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverse += string(str[i])
	}
	return reverse
}
```

* Usage
```golang
ResultReverse := ReturnReverseString("Test Reverse String")
println(ResultReverse)
```

* Result
```
gnirtS esreveR tseT
```
