package configs

import "os"

var (
	LanguagesData, _ = os.ReadFile("language_extensions.json")
)
