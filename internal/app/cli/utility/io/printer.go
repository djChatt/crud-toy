package io

import (
	"os"

	"crud-toy/internal/app/service/infra/proc"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

type Printer interface {
	Println(string, ...color.Attribute)
	PrintTable(procLists []proc.Proc)
}

var PrinterInstance Printer

type printer struct{}

func InitPrinter() {
	if PrinterInstance == nil {
		PrinterInstance = &printer{}
	}
}

func (p *printer) Println(msg string, attr ...color.Attribute) {
	color.New(attr...).Println(msg)
}

func (p *printer) PrintTable(procsList []proc.Proc) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Author"})
	for _, proc := range procsList {
		t.AppendRows([]table.Row{
			{proc.ID, proc.Name, proc.Author},
		})
		t.AppendSeparator()
	}
	t.Render()
}
