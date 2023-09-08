package google_tts

import (
	"anki/lib/xstring"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"google.golang.org/api/texttospeech/v1"
)

type synthesizeOptions interface {
	Apply(*texttospeech.SynthesizeSpeechRequest) bool
}

type synthesisInputText string
type voiceSelectionParamsLanguageCode string

func SynthesisText(v string) synthesisInputText { return synthesisInputText(v) }
func SynthesisLanguageCode(v string) voiceSelectionParamsLanguageCode {
	return voiceSelectionParamsLanguageCode(v)
}

var _ synthesizeOptions = synthesisInputText("")
var _ synthesizeOptions = voiceSelectionParamsLanguageCode("")

func (self synthesisInputText) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Input == nil {
		t.Input = &texttospeech.SynthesisInput{}
	}
	t.Input.Text = string(self)
	return true
}

func (self voiceSelectionParamsLanguageCode) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Voice == nil {
		t.Voice = &texttospeech.VoiceSelectionParams{}
	}
	t.Voice.LanguageCode = string(self)
	return true
}

func (self googleTTS) Synthesize(
	ctx context.Context,
	options ...synthesizeOptions,
) error {
	opt := &texttospeech.SynthesizeSpeechRequest{
		AudioConfig: &texttospeech.AudioConfig{
			AudioEncoding: string(AudioEncodingMP3),
		},
	}
	for _, option := range options {
		option.Apply(opt)
	}

	rq, err := self.service.Text.Synthesize(opt).Do()
	if err != nil {
		return err
	}

	data, err := base64.StdEncoding.DecodeString(rq.AudioContent)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(
		"./%s_%s.mp3",
		opt.Voice.LanguageCode,
		xstring.ExcludeWhiteSpaces(opt.Input.Text),
	)
	err = ioutil.WriteFile(path, data, 777)
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}
