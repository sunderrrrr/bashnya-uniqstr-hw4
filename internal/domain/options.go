package domain

type Options struct {
	C      bool // написать количество на выходе у строки
	D      bool // вывести только которые повторялись
	U      bool // вывести только уникальные
	F      int  // не учитывать первые N слов
	S      int  // не учитывать первые N символов
	I      bool // не учитывать регистр
	Input  string
	Output string
}
