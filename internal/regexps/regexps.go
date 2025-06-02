package regexps

import (
	"regexp"
)

var (
	NameRegex  = regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ ]{3,150}$`)
	TelRegex   = regexp.MustCompile(`^\+[0-9]{11}$`)
	EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)
