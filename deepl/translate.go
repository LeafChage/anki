package deepl

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// ref: https://www.deepl.com/ja/docs-api/translate-text/translate-text

type translateRequest struct {
	Text       []string    `json:"text"`
	SourceLang *SourceLang `json:"source_lang"`
	TargetLang TargetLang  `json:"target_lang"`
}

type Translate struct {
	DetectedSourceLanguage SourceLang `json:"detected_source_language"`
	Text                   string     `json:"text"`
}

type TranslateResponse struct {
	Translations []Translate `json:"translations"`
}

type TranslateRequest interface {
	Apply(*translateRequest)
}

type translateText string

func TranslateText(v string) translateText { return translateText(v) }

var _ TranslateRequest = translateText("")
var _ TranslateRequest = SourceLang("")
var _ TranslateRequest = TargetLang("")

func (self translateText) Apply(r *translateRequest) { r.Text = []string{string(self)} }
func (self SourceLang) Apply(r *translateRequest)    { r.SourceLang = &self }
func (self TargetLang) Apply(r *translateRequest)    { r.TargetLang = self }

func (self deepLClient) Translate(
	text translateText,
	targetLang TargetLang,
	optional ...TranslateRequest,
) (*TranslateResponse, error) {
	body := &translateRequest{}
	text.Apply(body)
	targetLang.Apply(body)
	for _, opt := range optional {
		opt.Apply(body)
	}

	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	res, err := requestDeepL(
		"POST",
		deepLURL("v2/translate"),
		self.apiKey,
		bytes.NewReader(data),
	)
	if err != nil {
		return nil, err
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	v := &TranslateResponse{}
	err = json.Unmarshal(resData, v)
	if err != nil {
		return nil, err
	}
	return v, nil
}
