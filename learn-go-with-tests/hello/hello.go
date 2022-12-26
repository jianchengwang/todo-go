package hello

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(name string, language string) (prefix string) {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
