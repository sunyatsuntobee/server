package controllers

import "html/template"

type layout struct {
	Title   string
	Content template.HTML
}
