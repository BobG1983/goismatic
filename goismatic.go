package goismatic

// Lang represents the languages supported by this library
type Lang int

// List of supported languages
const (
	English Lang = iota
	Russian
)

// Quote returns a random quote in either English or Russian from forismatic.com
func Quote(lang Lang) string {
	return "Party on wayne"
}
