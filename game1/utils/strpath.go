package utils

import (
	"path/filepath"
	"strings"
)

func ParsePath1(path string) string {
	path = filepath.Clean(path)
	path = strings.ReplaceAll(path, "\\", "/")
	path = strings.TrimPrefix(path, "../")
	path = strings.TrimPrefix(path, "../")
	return path
}
