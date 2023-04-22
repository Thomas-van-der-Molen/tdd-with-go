package main

import "fmt"

const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const frenchPrefix = "Bonjour, "
const dutchPrefix = "Hallo, "
const spanish = "Spanish"
const dutch = "Dutch"
const french = "French"
const punctuation = "!"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingWithPrefix(language) + name + punctuation
}

func main() {
	fmt.Println(hello("world", ""))
}

func greetingWithPrefix(language string) (prefix string){
	switch language {
	case french:
		prefix =  frenchPrefix
	case spanish:
		prefix = spanishPrefix 
	case dutch:
		prefix = dutchPrefix
	default:
		prefix = englishPrefix
	}
	return
}
