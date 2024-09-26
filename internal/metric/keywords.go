package metric

import (
	"ai-feed/internal/entity"
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/aaaton/golem/v4/dicts/ru"
	"github.com/parnurzeal/gorequest"
	"github.com/pemistahl/lingua-go"
	"github.com/rs/zerolog/log"
	"regexp"
	"slices"
	"strings"
)

var detectLanguages = []lingua.Language{
	lingua.English,
	lingua.Russian,
}

var letterPattern = regexp.MustCompile("[a-zA-Zа-яА-Я]+")

var detector = lingua.NewLanguageDetectorBuilder().
	FromLanguages(detectLanguages...).
	Build()

var (
	russianStopWords []string
	englishStopWords []string
)

// fetching stop words on start
func init() {
	// russian
	_, body, errs := gorequest.New().
		Get("https://raw.githubusercontent.com/stopwords-iso/stopwords-ru/master/stopwords-ru.txt").End()

	if len(errs) > 0 {
		log.Fatal().Interface("errors", errs)
	}

	words := strings.Fields(body)
	russianStopWords = words

	// english
	_, body, errs = gorequest.New().
		Get("https://raw.githubusercontent.com/stopwords-iso/stopwords-ru/master/stopwords-ru.txt").End()

	if len(errs) > 0 {
		log.Fatal().Interface("errors", errs)
	}

	words = strings.Fields(body)
	englishStopWords = words
}

func Keywords(text string) entity.Keywords {
	// detect text language
	language, exists := detector.DetectLanguageOf(text)

	if !exists {
		return nil
	}

	var (
		pack      golem.LanguagePack
		stopWords []string
	)

	switch language.IsoCode639_1().String() {
	case "EN":
		pack = en.New()
		stopWords = englishStopWords
	case "RU":
		pack = ru.New()
		stopWords = russianStopWords

	default:
		return nil
	}

	// lemmatize all words in text (aligning -> align)
	// and put them in lowercase

	lemmatizer, err := golem.New(pack)
	if err != nil {
		panic(err)
	}

	words := strings.Fields(text)

	for i := 0; i < len(words); i++ {
		word := strings.ToLower(words[i])
		word = lemmatizer.Lemma(word)
		word = removeNonLetters(word)
		words[i] = word
	}

	// delete words which are in stop words list

	filteredWords := make([]string, 0, len(words))

	for _, word := range words {
		if !slices.Contains(stopWords, word) {
			filteredWords = append(filteredWords, word)
		}
	}

	words = filteredWords

	counts := make(map[string]int)

	for _, word := range words {
		counts[word] += 1
	}

	result := make([]*entity.Keyword, 0, len(counts))

	for key, value := range counts {
		result = append(result, &entity.Keyword{
			Name:  key,
			Count: value,
		})
	}

	slices.SortFunc(result, func(a, b *entity.Keyword) int {
		return b.Count - a.Count
	})

	return result
}

func removeNonLetters(str string) string {
	return string(letterPattern.Find([]byte(str)))
}
