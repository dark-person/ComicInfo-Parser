package definitions

// SessionData store data that will be used as session data in frontend.
type SessionData struct {
	// Method to export. This will keep save last option until program close.
	ExportMethod ExportMethod `json:"exportMethod"`

	// Delete input folder after export
	DeleteAfterExport bool `json:"deleteAfterExport"`
}
