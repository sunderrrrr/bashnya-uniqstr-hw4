package cli

import (
	"flag"
	"fmt"
	"uniqstr/internal/domain"
)

func argParse() (*domain.Options, error) {
	count := flag.Bool("c", false, "Вывести перед строкой количество вхождений")
	repeated := flag.Bool("d", false, "Вывести строки, которые встречаются более одного раза")
	uniq := flag.Bool("u", false, "Вывести строки, которые встречаются только 1 раз")
	firstWords := flag.Int("f", 0, "Не учитывать первые N слов")
	firstChars := flag.Int("s", 0, "Не учитывать первые N символов")
	charCase := flag.Bool("i", false, "Не учитывать регистр")
	flag.Parse()
	if (*count && *repeated) || (*count && *uniq) || (*repeated && *uniq) {
		return nil, fmt.Errorf("флаги -c, -d, -u взаимоисключающие")
	}
	args := flag.Args()
	var input_file, output_file string
	if len(args) > 0 {
		input_file = args[0]
	}
	if len(args) > 1 {
		output_file = args[1]
	}
	return &domain.Options{
		C:      *count,
		D:      *repeated,
		U:      *uniq,
		F:      *firstWords,
		S:      *firstChars,
		I:      *charCase,
		Input:  input_file,
		Output: output_file,
	}, nil
}
