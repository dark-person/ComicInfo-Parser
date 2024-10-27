package history

// Container for history module,
// which wraps all necessary value will be used to insert value.
//
// This type is designed for insert value with different category at a time.
type HistoryVal struct {
	Category categoryType // category to be inserted
	Value    string       // value to be inserted
}
