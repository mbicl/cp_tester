package utils

import "github.com/mehanizm/iuliia-go"

func Cyrill2Latin(input string) string {
	return iuliia.Wikipedia.Translate(input)
}
