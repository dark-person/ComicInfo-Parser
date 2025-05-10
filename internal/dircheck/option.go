package dircheck

// The type for limit developer to pass in ScanOpt
type optEnum int

const (
	// Unspecific criteria, bypass option matching.
	Unspecific optEnum = iota
	// Allow the item to exist
	Allow
	// Allow the item to exist, and not allow other items
	AllowOnly
	// Not Allow the item to exist
	Reject
)

// The option for dircheck package.
type ScanOpt struct {
	// Option for given path has any sub-folder
	SubFolder optEnum

	// Option for given path has any image
	Image optEnum
}

// Check the ScanOpt is valid for process, prevent contradiction among fields.
func (opt *ScanOpt) Valid() bool {
	if opt.Image == AllowOnly && (opt.SubFolder == Allow || opt.SubFolder == AllowOnly) {
		return false
	}

	if opt.SubFolder == AllowOnly && (opt.Image == Allow || opt.Image == AllowOnly) {
		return false
	}

	return true
}
