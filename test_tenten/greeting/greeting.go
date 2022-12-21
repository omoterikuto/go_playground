package greeting

import "time"

func Do(t time.Time) string {
	switch h := t.Hour(); {
	case h >= 4 && h <= 9:
		return "おはよう"
	case h >= 10 && h <= 16:
		return "こんにちは"
	default:
		return "こんばんは"
	}
}
