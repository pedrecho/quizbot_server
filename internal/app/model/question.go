package model

type Question struct {
	ID      int64
	Package string
	Text    string
	Option1 string
	Option2 string
	Option3 string
	Option4 string
	Answer  int8
}
