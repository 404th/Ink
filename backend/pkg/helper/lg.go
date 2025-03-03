package helper

import (
	"strings"

	"github.com/404th/Ink/config"
)

func SplitErrorMessage(message string) string {
	msgs := strings.Split(message, config.ErrorSplitter)

	if len(msgs) < 1 {
		return "Noma'lum xatolik"
	}

	return msgs[len(msgs)-1]
}

func SplitInfoMessage(message string) string {
	msgs := strings.Split(message, config.InfoSplitter)

	if len(msgs) < 1 {
		return "Noma'lum ma'lumot"
	}

	return msgs[len(msgs)-1]
}
