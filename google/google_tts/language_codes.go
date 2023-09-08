package google_tts

import (
	"context"
	"github.com/deckarep/golang-set/v2"
)

func (self googleTTS) LanguageCodes(ctx context.Context) ([]string, error) {
	rq, err := self.service.Voices.List().Do()
	if err != nil {
		return []string{}, err
	}

	set := mapset.NewSet[string]()
	for _, v := range rq.Voices {
		for _, l := range v.LanguageCodes {
			set.Add(l)
		}
	}
	return set.ToSlice(), err
}
