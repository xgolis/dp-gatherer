package errhelp

import (
	"errors"

	"github.com/xgolis/dp-gatherer/cmd/gatherer/app"
)

var errorMap = map[error]string{
	app.ErrorVersion: "Unfortunately your version of dp-gatherer is outdated.\n" +
		"Please update your version of dp-gatherer",
}

func GetHelp(err error) string {
	for supportedErr, helpMessage := range errorMap {
		if errors.Is(err, supportedErr) {
			return helpMessage
		}
	}
	return ""
}
