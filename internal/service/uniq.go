package service

import (
	"fmt"
	"strings"
	"uniqstr/internal/domain"
)

type UniqService struct {
}

func NewUniqService() *UniqService {
	return &UniqService{}
}

func (s *UniqService) Process(lines []string, options *domain.Options) (string, error) {
	counts := make(map[string]int)
	printed := make(map[string]bool)
	keys := make(map[string]string)

	for _, line := range lines {
		k := prepareLine(line, options)
		counts[k]++
		keys[k] = line
	}
	result := ""
	for _, line := range lines {
		k := prepareLine(line, options)
		switch {
		case options.C:
			if !printed[k] {
				result += fmt.Sprintf("%d %s\n", counts[k], line)
				printed[k] = true
			}
		case options.D:
			if counts[k] > 1 && !printed[k] {
				result += line + "\n"
				printed[k] = true
			}
		case options.U:
			if counts[k] == 1 && !printed[k] {
				result += line + "\n"
				printed[k] = true
			}
		default:
			if !printed[k] {
				result += line + "\n"
				printed[k] = true
			}
		}
	}
	return result, nil
}

func skipFields(line string, skip int) string {
	fields := strings.Fields(line)
	if skip >= len(fields) {
		return ""
	} else {
		return strings.Join(fields[skip:], " ")
	}
}

func skipChars(line string, skip int) string {
	if skip >= len(line) {
		return ""
	}
	return line[skip:]
}

func prepareLine(line string, options *domain.Options) string {
	res := line
	if options.I {
		return strings.ToLower(res)
	}
	if options.F > 0 {
		return skipFields(res, options.F)
	}
	if options.S > 0 {
		return skipChars(res, options.S)
	}
	return res
}
