package utils

import "path"

// fetch filepath
func GetFilepath(dayNumber string) string {
	return path.Join("..", "..", "data", "day"+dayNumber+".txt")
}
