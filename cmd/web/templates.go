package main

import "kailashgautham.com/snippetbox/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
