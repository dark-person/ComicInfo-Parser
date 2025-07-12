// Package that contains definitions of different types, that is shared between packages.
package definitions

// Category enum-like constants. Consistent with database definitions.
type CategoryType int

const (
	CategoryGenre      CategoryType = iota + 1 // Database value for Genre.
	CategoryPublisher                          // Database value for Publisher.
	CategoryTranslator                         // Database value for Translator.
	CategoryTag                                // Database value for Tag.
)

// Return a slices of category type.
func Categories() []CategoryType {
	return []CategoryType{CategoryGenre, CategoryPublisher, CategoryTranslator, CategoryTag}
}
