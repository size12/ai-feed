package feeder

import (
	"ai-feed/internal/entity"
	"context"
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/tidwall/gjson"
	"net/http"
)

func VcRuFeeder(ctx context.Context) ([]*entity.Theme, error) {
	url := "https://api.vc.ru/v2.5/feed"

	resp, body, errs := gorequest.New().Get(url).End()

	if len(errs) > 0 {
		return nil, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200")
	}

	value := gjson.Get(body, "result.items.#(type==\"news\").data.news.#.title")

	if !value.IsArray() {
		return nil, errors.New("can't parse news")
	}

	array := value.Array()

	var themes []*entity.Theme

	for _, el := range array {
		themes = append(themes, &entity.Theme{
			Description: el.String(),
		})
	}

	return themes, nil
}
