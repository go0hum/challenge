package domain

type TemplateEmail struct {
	Logo        string
	Total       float32
	TotalDebit  float32
	TotalCredit float32
	ConteoMes   map[string]int8
}
