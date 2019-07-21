package main

type PreviewData struct {
	Columns []PdColumn
}

type PdColumn struct {
	Title string
	Position int
	ExcelPosition int
	ExcelColumn string
	Samples []string
}