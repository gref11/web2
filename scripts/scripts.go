package scripts

import (
	"net/http"
	"strings"
)

func In(haystack []string, needle string) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}
	return false
}

func EnableCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func ConvertToMnemonic(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp")
	s = strings.ReplaceAll(s, "<", "&lt")
	s = strings.ReplaceAll(s, ">", "&gt")
	s = strings.ReplaceAll(s, "'", "&#39")
	s = strings.ReplaceAll(s, `"`, "&quot")
	return s
}
