package utils

import "path"

func GetFilepath(dayNumber string) string {
	return path.Join("..", "data", "day"+dayNumber+".txt")
}
