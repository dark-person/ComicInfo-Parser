// Package that contains definitions of different types, that is shared between packages.
package definitions

// Category enum-like constants. Consistent with database definitions.
type CategoryType int

const (
	CategoryGenre      CategoryType = iota + 1 // Database value for Genre.
	CategoryPublisher                          // Database value for Publisher.
	CategoryTranslator                         // Database value for Translator.
)
