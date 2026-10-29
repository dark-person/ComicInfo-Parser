package definitions

import (
	"encoding"
	"strings"
)

// Method for export comicinfo cbz.
type ExportMethod int

var _ encoding.TextUnmarshaler = (*ExportMethod)(nil)

const (
	ExportMethodDefaultWrapCBZ ExportMethod = iota // Export cbz file with a default folder wrapped, same with cbz filename.
	ExportMethodOnlyCBZ                            // Only export single cbz file.
	ExportMethodCustomWrapCBZ                      // Export cbz file with a custom folder wrapped.
)

func (l *ExportMethod) UnmarshalText(text []byte) error {
	str := strings.ToLower(string(text))
	switch str {
	case "file":
		*l = ExportMethodOnlyCBZ
	case "default-wrap":
		*l = ExportMethodDefaultWrapCBZ
	case "custom-wrap":
		*l = ExportMethodCustomWrapCBZ
	default:
		*l = ExportMethodDefaultWrapCBZ
	}
	return nil
}
