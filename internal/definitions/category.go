package definitions

// Category enum-like constants. Consistent with database definitions.
type CategoryType int

const (
	CategoryGenre      CategoryType = iota + 1 // Database value for Genre.
	CategoryPublisher                          // Database value for Publisher.
	CategoryTranslator                         // Database value for Translator.
	CategoryTag                                // Database value for Tag.
	CategoryWriter                             // Database value for Writer
	CategoryCharacters                         // Database value for characters.
	CategoryTeams                              // Database value for teams.
	CategoryLocations                          // Database value for locations.
)

// Return a slices of category type.
func Categories() []CategoryType {
	return []CategoryType{
		CategoryGenre,
		CategoryPublisher,
		CategoryTranslator,
		CategoryTag,
		CategoryWriter,
		CategoryCharacters,
		CategoryTeams,
		CategoryLocations,
	}
}
