package common

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"strings"
)

func IsSupportedImageType(list map[string]bool, data string) bool {
	splitContentType := strings.Split(data, "/")
	if strings.ToUpper(splitContentType[0]) != "IMAGE" && !list[splitContentType[1]] {
		return false
	}

	return true
}
