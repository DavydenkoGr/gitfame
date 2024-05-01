package configs

import "os"

// language config content
var	LanguagesData, _ = os.ReadFile("language_extensions.json")
