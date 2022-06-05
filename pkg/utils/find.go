package utils

import "strings"

func TextContainOne(text string, contain ...string) bool {
	for _, v := range contain {
		if strings.Contains(text, v) {
			return true
		}
	}
	return false
}

func TextContainAll(text string, contain ...string) bool {
	for _, v := range contain {
		if !strings.Contains(text, v) {
			return false
		}
	}
	return true
}
