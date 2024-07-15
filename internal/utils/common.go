package utils

import "time"

// ParseAsTime parse string as time.DateTime
func ParseAsTime(s string) (time.Time, error) {
	return time.Parse(time.DateTime, s)
}

// GetStrValOrDef get s value if exist else def value else empty string
func GetStrValOrDef(s *string, def *string) string {
	if s != nil {
		return *s
	}
	if def != nil {
		return *def
	}
	return ""
}
