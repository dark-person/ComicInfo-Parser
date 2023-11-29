package scanner

// The type for limit developer to pass in ScanOpt
type optEnum int

const (
	// Unspecific criteria, bypass option matching.
	Unspecific optEnum = iota
	// Contain item that matches option, can include item that not match option
	Contain
	// Only contain item that matches option
	ContainOnly
	// Exclude all item that matches option
	Exclude
)

// The option for scanner package.
type ScanOpt struct {
	// Option for given path has any sub-folder
	SubFolder optEnum

	// Option for given path has any image
	Image optEnum
}
