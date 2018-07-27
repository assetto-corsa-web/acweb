package util

// Interface to treat sql.Row and sql.Rows as the same, when parsing results.
type RowScanner interface {
	Scan(...interface{}) error
}
