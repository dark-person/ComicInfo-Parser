package definitions

// Method for export comicinfo cbz.
type ExportMethod int

const (
	ExportMethodOnlyCBZ        ExportMethod = iota + 1 // Only export single cbz file.
	ExportMethodDefaultWrapCBZ                         // Export cbz file with a default folder wrapped, same with cbz filename.
	ExportMethodCustomWrapCBZ                          // Export cbz file with a custom folder wrapped.
)
