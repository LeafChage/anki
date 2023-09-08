package google_tts

import (
	"context"

	"google.golang.org/api/texttospeech/v1"
)

type filterableVoiceName interface {
	Pass(*texttospeech.Voice) bool
}

// options
type languageCode string

var _ filterableVoiceName = languageCode("")
var _ filterableVoiceName = SSMGender("")

func LanguageCode(v string) languageCode { return languageCode(v) }

func (l languageCode) Pass(voice *texttospeech.Voice) bool {
	if l == "_" {
		return true
	}

	for _, v := range voice.LanguageCodes {
		if v == string(l) {
			return true
		}
	}
	return false
}
func (g SSMGender) Pass(voice *texttospeech.Voice) bool {
	return g == GenderUndefined || string(g) == voice.SsmlGender
}

func (self googleTTS) VoiceName(ctx context.Context, filters ...filterableVoiceName) ([]string, error) {
	rq, err := self.service.Voices.List().Do()
	if err != nil {
		return []string{}, err
	}

	result := []string{}
	for _, v := range rq.Voices {
		for _, f := range filters {
			if f.Pass(v) && f.Pass(v) {
				result = append(result, v.Name)
			}
		}
	}
	return result, nil
}
