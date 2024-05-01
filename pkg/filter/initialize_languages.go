package filter

import (
	"encoding/json"
	"github.com/DavydenkoGr/gitfame/configs"
)

type language struct {
	name string `json:"name"`
    languageType string `json:"type"`
    extensions []string `json:"extensions"`
}

func (f *Filter) InitializeLanguages(permittedLanguagesList []string) {
	permittedLanguagesSet := map[string]struct{}{}

	for _, language := range permittedLanguagesList {
		permittedLanguagesSet[language] = struct{}{}
	}

	var languages []language
	_ = json.Unmarshal(configs.LanguagesData, &languages)
	
	for _, language := range languages {
		if _, ok := permittedLanguagesSet[language.name]; !ok {
			continue
		}

		for _, extension := range language.extensions {
			f.LanguagesSet[extension] = struct{}{}
		}
	}
}
