package libs

import "github.com/bregydoc/gtranslate"

func TranslateEnToEs(text string) string {
	translation, _ := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: "en",
			To:   "es",
		},
	)
	return translation
}