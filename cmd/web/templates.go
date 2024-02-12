package main

import (
	"github.com/Caps1d/Lets-Go/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
