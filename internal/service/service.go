package service

import (
	"errors"
	"strings"
	"morse_converter/pkg/morse"
)

var (
	ErrEmptyInput = errors.New("input cannot be empty")
)

func Convert(input string) (string, error) {
	// Проверка на пустой ввод
	if len(input) == 0 {
		return "", ErrEmptyInput
	}

	// Проверка, содержит ли строка символы Морзе
	if strings.ContainsAny(input, ".-") {
		return morse.ToText(input), nil
	}
	
	// Если не Морзе, конвертируем текст в Морзе
	return morse.ToMorse(input), nil
}