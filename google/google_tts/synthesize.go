package google_tts

import (
	"anki/lib/xstring"
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"google.golang.org/api/texttospeech/v1"
)

type SynthesizeOptions interface {
	Apply(*texttospeech.SynthesizeSpeechRequest) bool
}

type synthesisText string
type synthesisLanguageCode string
type synthesisVoiceName string

func SynthesisText(v string) synthesisText                 { return synthesisText(v) }
func SynthesisLanguageCode(v string) synthesisLanguageCode { return synthesisLanguageCode(v) }
func SynthesisVoiceName(v string) synthesisVoiceName       { return synthesisVoiceName(v) }

var _ SynthesizeOptions = synthesisText("")
var _ SynthesizeOptions = synthesisLanguageCode("")
var _ SynthesizeOptions = SSMGender("")
var _ SynthesizeOptions = synthesisVoiceName("")

func (self synthesisText) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Input == nil {
		t.Input = &texttospeech.SynthesisInput{}
	}
	t.Input.Text = string(self)
	return true
}

func (self synthesisLanguageCode) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Voice == nil {
		t.Voice = &texttospeech.VoiceSelectionParams{}
	}
	t.Voice.LanguageCode = string(self)
	return true
}

func (self SSMGender) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Voice == nil {
		t.Voice = &texttospeech.VoiceSelectionParams{}
	}
	t.Voice.SsmlGender = string(self)
	return true
}

func (self synthesisVoiceName) Apply(t *texttospeech.SynthesizeSpeechRequest) bool {
	if t.Voice == nil {
		t.Voice = &texttospeech.VoiceSelectionParams{}
	}
	t.Voice.Name = string(self)
	return true
}

func (self googleTTS) Synthesize(
	ctx context.Context,
	options ...SynthesizeOptions,
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
		"./%s_%s_%s.mp3",
		opt.Voice.LanguageCode,
		xstring.ExcludeWhiteSpaces(opt.Input.Text),
		time.Now().Format("060102_150405"),
	)
	err = ioutil.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Println(path)
	return nil
}
