package common

import (
	newError "bwa-startup/internal/common/errors"
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"net/url"
	"strconv"
	"strings"
)

func IsSupportedImageType(list map[string]bool, data string) error {
	splitContentType := strings.Split(data, "/")
	if strings.ToUpper(splitContentType[0]) != "IMAGE" && !list[splitContentType[1]] {
		return newError.ErrUnsupportedMediaType
	}

	return nil
}

func GetUserId(value any) int {
	userIdString := fmt.Sprint(value)
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return -1
	}
	return userId
}

func GetUrlImage(bucket, image, token string) string {
	return fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media&token=%s", bucket, url.PathEscape(image), token)
}
