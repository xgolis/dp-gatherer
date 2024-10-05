package version

import "fmt"

var (
	Time   string
	Commit string
	Branch string
	Tag    string
)

func Get() string {
	if Tag != "" {
		return fmt.Sprintf("%s (%s)", Tag, Time)
	}
	return fmt.Sprintf("%s-%s (%s)", Branch, Commit, Time)
}
