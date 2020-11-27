package main

import "fmt"

const spanish = "Spanish"
const prefix = "Hello,"
const sprefix = "Hola,"
const french = "French"
const fprefix = "French"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	//p := prefix
	//switch language {
	//case sprefix:
	//	p = sprefix
	//case fprefix:
	//	p = fprefix
	//}

	return getPrefix(language) + name
}

func getPrefix(language string) (p string) {
	switch language {
	case spanish:
		p = sprefix
	case french:
		p = fprefix
	default:
		p = prefix
	}
	return
}

func main() {
	fmt.Println(Hello("long", ""))
}
