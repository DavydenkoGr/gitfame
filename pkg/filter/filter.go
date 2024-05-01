package filter

type Filter struct {
	ExtensionsSet map[string]struct{}
	LanguagesSet map[string]struct{}
	Exclude []string
	RestrictTo []string
}
