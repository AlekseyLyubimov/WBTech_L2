package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
Задача на распаковку
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func main() {
	println(unpack("a4bc2d5e"))
}

func unpack(s string) (string, error) {
	if len(s) == 0 || unicode.IsDigit(rune(s[0])) {
		return "", fmt.Errorf("bad string")
	}

	runes := []rune(s)
	outstr := ""

	for i := 0; i < len(runes); i++ {
		if (unicode.IsDigit(runes[i])) {
			num, _ := strconv.Atoi(string(runes[i]))
			outstr += strings.Repeat(string(runes[i-1]), num-1)
		} else {
			outstr += string(runes[i])
		}
	}

	return outstr, nil
}