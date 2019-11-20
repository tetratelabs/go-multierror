package multierror

import (
	"fmt"
	"strings"
)

// ErrorFormatFunc is a function callback that is called by Error to
// turn the list of errors into a string.
type ErrorFormatFunc func([]error) string

// ListFormatFunc is a basic formatter that outputs the number of errors
// that occurred along with a bullet point list of the errors.
//
// Typically this format is used for CLI exit responses.
func ListFormatFunc(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred:\n\t* %s\n\n", es[0])
	}

	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}

	return fmt.Sprintf(
		"%d errors occurred:\n\t%s\n\n",
		len(es), strings.Join(points, "\n\t"))
}

// LineErrorFormatFunc is a basic formatter that outputs the number of errors
// that occurred along with all errors on a single line.
//
// Typically this format is used for logs.
func LineErrorFormatFunc(es []error) string {
	if len(es) == 1 {
		return es[0].Error()
	}

	var b strings.Builder
	for _, err := range es {
		b.WriteString(err.Error())
		b.WriteString("; ")
	}
	return fmt.Sprintf("%d errors occurred: %s", len(es), b.String()[:b.Len()-2])
}
