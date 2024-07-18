package helloworld

const (
	enPrefix = "Hello, "
	esPrefix = "Hola, "
	frPrefix = "Bonjour, "
	uzPrefix = "As Salomu Alaykum, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = esPrefix
	case "French":
		prefix = frPrefix
	case "Uzbek":
		prefix = uzPrefix
	default:
		prefix = enPrefix
	}

	return
}
