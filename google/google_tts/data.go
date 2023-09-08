package google_tts

type SSMGender string

const (
	GenderMale      SSMGender = "MALE"
	GenderFemale    SSMGender = "FEMALE"
	GenderNeutral   SSMGender = "NEUTRAL"
	GenderUndefined SSMGender = "_"
)

type AudioEncoding string

const (
	AudioEncodingAUDIO_ENCODING_UNSPECIFIED AudioEncoding = "AUDIO_ENCODING_UNSPECIFIED"
	AudioEncodingMP3                        AudioEncoding = "MP3"
	AudioEncodingOGG_OPUS                   AudioEncoding = "OGG_OPUS"
	AudioEncodingMULAW                      AudioEncoding = "MULAW"
	AudioEncodingALAW                       AudioEncoding = "ALAW"
)
