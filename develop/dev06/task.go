package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type flags struct {
	fields      string
	delimiter   string
	isSeparated bool
}

func parseFlags() *flags {
	f := flags{}

	flag.StringVar(&f.fields, "f", "0", "fields")
	flag.StringVar(&f.delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&f.isSeparated, "s", false, "separated")
	flag.Parse()

	return &f
}

func cut(input string, f *flags) string {
	if f.isSeparated && !strings.Contains(input, f.delimiter) {
		return ""
	}

	sb := strings.Builder{}
	splitted := strings.Split(input, f.delimiter)
	columns := strings.Split(f.fields, ",")
	for i := 0; i < len(columns); i++ {
		column, err := strconv.Atoi(columns[i])
		if err != nil {
			log.Fatalln("cannot parse column to int: ", err.Error())
		}

		sb.WriteString(splitted[column])
	}
	return sb.String()
}

func main() {
	f := parseFlags()
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println("CUT: ", cut(text, f))
	}
}