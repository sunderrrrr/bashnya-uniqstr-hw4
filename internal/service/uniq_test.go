package service

import (
	"testing"
	"uniqstr/internal/domain"
)

type uniqTest struct {
	name     string
	lines    []string
	options  *domain.Options
	expected string
}

func TestUniqService_Process(t *testing.T) {
	tests := []uniqTest{
		uniqTest{
			"без параметров - уникальные строки",
			[]string{"aab", "bba", "abc", "aab"},
			&domain.Options{},
			"aab\nbba\nabc\n",
		},
		{
			name:  "флаг -c - подсчет вхождений",
			lines: []string{"a", "b", "a", "c", "b"},
			options: &domain.Options{
				C: true, D: false, U: false, F: 0, S: 0, I: false,
			},
			expected: "2 a\n2 b\n1 c\n",
		},
		{
			name:  "флаг -d - только повторяющиеся",
			lines: []string{"a", "b", "a", "c", "b", "d"},
			options: &domain.Options{
				C: false, D: true, U: false, F: 0, S: 0, I: false,
			},
			expected: "a\nb\n",
		},
		{
			name:  "флаг -u - только уникальные",
			lines: []string{"a", "b", "a", "c", "b", "d"},
			options: &domain.Options{
				C: false, D: false, U: true, F: 0, S: 0, I: false,
			},
			expected: "c\nd\n",
		},
		{
			name:  "флаг -i - игнорирование регистра",
			lines: []string{"Apple", "apple", "BANANA", "banana", "Cherry"},
			options: &domain.Options{
				C: false, D: false, U: false, F: 0, S: 0, I: true,
			},
			expected: "Apple\nBANANA\nCherry\n",
		},
		{
			name:  "флаг -f 1 - игнорирование первого поля",
			lines: []string{"1 apple", "2 apple", "3 banana", "1 cherry"},
			options: &domain.Options{
				C: false, D: false, U: false, F: 1, S: 0, I: false,
			},
			expected: "1 apple\n3 banana\n1 cherry\n",
		},
		{
			name:  "флаг -s 2 - игнорирование первых 2 символов",
			lines: []string{"aaapple", "bbapple", "ccbanana", "ddcherry"},
			options: &domain.Options{
				C: false, D: false, U: false, F: 0, S: 2, I: false,
			},
			expected: "aaapple\nccbanana\nddcherry\n",
		},
		{
			name:  "комбинация -i -f 1",
			lines: []string{"Apple", "apple", "BANANA", "banana"},
			options: &domain.Options{
				C: false, D: false, U: false, F: 1, S: 0, I: true,
			},
			expected: "Apple\nBANANA\n",
		},
		{
			name:  "пустой ввод",
			lines: []string{},
			options: &domain.Options{
				C: false, D: false, U: false, F: 0, S: 0, I: false,
			},
			expected: "",
		},
		{
			name:  "одна строка",
			lines: []string{"single"},
			options: &domain.Options{
				C: false, D: false, U: false, F: 0, S: 0, I: false,
			},
			expected: "single\n",
		},
		{
			name:  "все строки уникальные с -c",
			lines: []string{"a", "b", "c"},
			options: &domain.Options{
				C: true, D: false, U: false, F: 0, S: 0, I: false,
			},
			expected: "1 a\n1 b\n1 c\n",
		},
		{
			name:  "все строки одинаковые с -u",
			lines: []string{"a", "a", "a"},
			options: &domain.Options{
				C: false, D: false, U: true, F: 0, S: 0, I: false,
			},
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewUniqService()
			res, err := service.Process(test.lines, test.options)
			if err != nil {
				t.Error(err)
				return
			}
			if res != test.expected {
				t.Errorf("got %s, want %s", res, test.expected)
			}
		})
	}
}
