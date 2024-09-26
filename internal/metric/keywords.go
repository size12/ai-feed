package metric

import "ai-feed/internal/entity"

func Keywords(text string) entity.Keywords {
	return []*entity.Keyword{
		{
			Name:  "test1",
			Count: 12,
		},
		{
			Name:  "test2",
			Count: 100,
		},
	}
}
