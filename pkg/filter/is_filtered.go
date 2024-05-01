package filter

import "path/filepath"

func (f *Filter) IsFiltered(filename string) bool {
	if len(f.ExtensionsSet) != 0 && !f.isFromExtensionSet(filename) {
		return false
	}

	if len(f.LanguagesSet) != 0 && !f.isFromLanguageSet(filename) {
		return false
	}

	if len(f.RestrictTo) != 0 && !f.isFromRestrictToPattern(filename) {
		return false
	}

	if len(f.Exclude) != 0 && f.isFromExcludePattern(filename) {
		return false
	}

	return true
}

func (f *Filter) isFromExtensionSet(filename string) bool {
	_, ok := f.ExtensionsSet[filepath.Ext(filename)]

	return ok
}

func (f *Filter) isFromLanguageSet(filename string) bool {
	_, ok := f.LanguagesSet[filepath.Ext(filename)]

	return ok
}

func (f *Filter) isFromRestrictToPattern(filename string) bool {
	for _, pattern := range f.RestrictTo {
		if match, _ := filepath.Match(pattern, filename); match {
			return true
		}
	}

	return false
}

func (f *Filter) isFromExcludePattern(filename string) bool {
	for _, pattern := range f.Exclude {
		if match, _ := filepath.Match(pattern, filename); match {
			return true
		}
	}

	return false
}
