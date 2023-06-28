package validators

import "regexp"

func IsAlpha(str string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z]+$`)
	return re.MatchString(str)
}
