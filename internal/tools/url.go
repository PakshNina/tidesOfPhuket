package tools

import "fmt"

func CreateFullUrl(url, methodPath string) string {
	return fmt.Sprintf("%s?%s", url, methodPath)
}

func AddToUrlParameter(url, key, value string) string {
	return fmt.Sprintf("%s&%s=%s", url, key, value)
}