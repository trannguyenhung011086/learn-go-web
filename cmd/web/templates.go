package main

import "trannguyenhung011086/learn-go-web/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
