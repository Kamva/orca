package orca

import (
	"net/url"

	"github.com/Kamva/nautilus"
)

// GetTranslationKey get translation key for given rule on given field
func GetTranslationKey(taggable nautilus.Taggable, field string, ruleName string) string {
	query := taggable.GetTag(taggable, field, "translation")
	mapping, _ := url.ParseQuery(query)

	return mapping[ruleName][0]
}
