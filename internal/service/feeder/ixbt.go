package feeder

import (
	"ai-feed/internal/entity"
	"context"
	"github.com/antchfx/htmlquery"
	"strings"
)

func IxbtFeeder(ctx context.Context) ([]*entity.Theme, error) {
	url := "https://www.ixbt.com/news/?show=tape"

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}

	var themes []*entity.Theme

	list := htmlquery.Find(doc, "//h2[contains(@class, \"no-margin\")]")
	for _, el := range list {
		themes = append(themes, &entity.Theme{
			Description: strings.TrimSpace(htmlquery.InnerText(el)),
		})
	}

	return themes, nil
}
