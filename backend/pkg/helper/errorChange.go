package helper

import (
	"errors"
	"strings"

	"github.com/404th/Ink/config"
)

func ChangeErrorForm(err error) (respErr error) {
	switch {
	case strings.Contains(err.Error(), config.TechMessageUniqueConstraint):
		respErr = errors.New(config.TechMessageUniqueConstraintCompatible)

	case strings.Contains(err.Error(), config.TechMessageNoRows):
		respErr = errors.New(config.TechMessageNoRowsCompatible)

	case strings.Contains(err.Error(), config.TechMessageEOF):
		respErr = errors.New(config.TechMessageEOFCompatible)

	case strings.Contains(err.Error(), config.TechMessageIncorrectPassword):
		respErr = errors.New(config.TechMessageIncorrectPasswordCompatible)

	default:
		respErr = errors.New(config.TechMessageInternalServerCompatible)
	}

	return
}
