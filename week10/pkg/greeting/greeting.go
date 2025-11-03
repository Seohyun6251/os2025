package greeting

import "fmt"

//함수이름 소문자로 시작하면 외부에 노출 X
func Hello() {
	fmt.Printf("hello\n")
}
func Hi() {
	fmt.Printf("hi\n")
}
