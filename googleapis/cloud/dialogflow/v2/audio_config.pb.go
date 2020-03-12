// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2/audio_config.proto

package dialogflow

import (
	fmt "fmt"
	math "math"

	proto "github.com/catper/protobuf/proto"
	duration "github.com/catper/protobuf/ptypes/duration"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Audio encoding of the audio content sent in the conversational query request.
// Refer to the
// [Cloud Speech API
// documentation](https://cloud.google.com/speech-to-text/docs/basics) for more
// details.
type AudioEncoding int32

const (
	// Not specified.
	AudioEncoding_AUDIO_ENCODING_UNSPECIFIED AudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	AudioEncoding_AUDIO_ENCODING_LINEAR_16 AudioEncoding = 1
	// [`FLAC`](https://xiph.org/flac/documentation.html) (Free Lossless Audio
	// Codec) is the recommended encoding because it is lossless (therefore
	// recognition is not compromised) and requires only about half the
	// bandwidth of `LINEAR16`. `FLAC` stream encoding supports 16-bit and
	// 24-bit samples, however, not all fields in `STREAMINFO` are supported.
	AudioEncoding_AUDIO_ENCODING_FLAC AudioEncoding = 2
	// 8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
	AudioEncoding_AUDIO_ENCODING_MULAW AudioEncoding = 3
	// Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz` must be 8000.
	AudioEncoding_AUDIO_ENCODING_AMR AudioEncoding = 4
	// Adaptive Multi-Rate Wideband codec. `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_AMR_WB AudioEncoding = 5
	// Opus encoded audio frames in Ogg container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_OGG_OPUS AudioEncoding = 6
	// Although the use of lossy encodings is not recommended, if a very low
	// bitrate encoding is required, `OGG_OPUS` is highly preferred over
	// Speex encoding. The [Speex](https://speex.org/) encoding supported by
	// Dialogflow API has a header byte in each block, as in MIME type
	// `audio/x-speex-with-header-byte`.
	// It is a variant of the RTP Speex encoding defined in
	// [RFC 5574](https://tools.ietf.org/html/rfc5574).
	// The stream is a sequence of blocks, one block per RTP packet. Each block
	// starts with a byte containing the length of the block, in bytes, followed
	// by one or more frames of Speex data, padded to an integral number of
	// bytes (octets) as specified in RFC 5574. In other words, each RTP header
	// is replaced with a single byte containing the block length. Only Speex
	// wideband is supported. `sample_rate_hertz` must be 16000.
	AudioEncoding_AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE AudioEncoding = 7
)

var AudioEncoding_name = map[int32]string{
	0: "AUDIO_ENCODING_UNSPECIFIED",
	1: "AUDIO_ENCODING_LINEAR_16",
	2: "AUDIO_ENCODING_FLAC",
	3: "AUDIO_ENCODING_MULAW",
	4: "AUDIO_ENCODING_AMR",
	5: "AUDIO_ENCODING_AMR_WB",
	6: "AUDIO_ENCODING_OGG_OPUS",
	7: "AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE",
}

var AudioEncoding_value = map[string]int32{
	"AUDIO_ENCODING_UNSPECIFIED":            0,
	"AUDIO_ENCODING_LINEAR_16":              1,
	"AUDIO_ENCODING_FLAC":                   2,
	"AUDIO_ENCODING_MULAW":                  3,
	"AUDIO_ENCODING_AMR":                    4,
	"AUDIO_ENCODING_AMR_WB":                 5,
	"AUDIO_ENCODING_OGG_OPUS":               6,
	"AUDIO_ENCODING_SPEEX_WITH_HEADER_BYTE": 7,
}

func (x AudioEncoding) String() string {
	return proto.EnumName(AudioEncoding_name, int32(x))
}

func (AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{0}
}

// Variant of the specified [Speech model][google.cloud.dialogflow.v2.InputAudioConfig.model] to use.
//
// See the [Cloud Speech
// documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
// for which models have different variants. For example, the "phone_call" model
// has both a standard and an enhanced variant. When you use an enhanced model,
// you will generally receive higher quality results than for a standard model.
type SpeechModelVariant int32

const (
	// No model variant specified. In this case Dialogflow defaults to
	// USE_BEST_AVAILABLE.
	SpeechModelVariant_SPEECH_MODEL_VARIANT_UNSPECIFIED SpeechModelVariant = 0
	// Use the best available variant of the [Speech
	// model][InputAudioConfig.model] that the caller is eligible for.
	//
	// Please see the [Dialogflow
	// docs](https://cloud.google.com/dialogflow/docs/data-logging) for
	// how to make your project eligible for enhanced models.
	SpeechModelVariant_USE_BEST_AVAILABLE SpeechModelVariant = 1
	// Use standard model variant even if an enhanced model is available.  See the
	// [Cloud Speech
	// documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
	// for details about enhanced models.
	SpeechModelVariant_USE_STANDARD SpeechModelVariant = 2
	// Use an enhanced model variant:
	//
	// * If an enhanced variant does not exist for the given
	//   [model][google.cloud.dialogflow.v2.InputAudioConfig.model] and request language, Dialogflow falls
	//   back to the standard variant.
	//
	//   The [Cloud Speech
	//   documentation](https://cloud.google.com/speech-to-text/docs/enhanced-models)
	//   describes which models have enhanced variants.
	//
	// * If the API caller isn't eligible for enhanced models, Dialogflow returns
	//   an error. Please see the [Dialogflow
	//   docs](https://cloud.google.com/dialogflow/docs/data-logging)
	//   for how to make your project eligible.
	SpeechModelVariant_USE_ENHANCED SpeechModelVariant = 3
)

var SpeechModelVariant_name = map[int32]string{
	0: "SPEECH_MODEL_VARIANT_UNSPECIFIED",
	1: "USE_BEST_AVAILABLE",
	2: "USE_STANDARD",
	3: "USE_ENHANCED",
}

var SpeechModelVariant_value = map[string]int32{
	"SPEECH_MODEL_VARIANT_UNSPECIFIED": 0,
	"USE_BEST_AVAILABLE":               1,
	"USE_STANDARD":                     2,
	"USE_ENHANCED":                     3,
}

func (x SpeechModelVariant) String() string {
	return proto.EnumName(SpeechModelVariant_name, int32(x))
}

func (SpeechModelVariant) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{1}
}

// Gender of the voice as described in
// [SSML voice element](https://www.w3.org/TR/speech-synthesis11/#edef_voice).
type SsmlVoiceGender int32

const (
	// An unspecified gender, which means that the client doesn't care which
	// gender the selected voice will have.
	SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED SsmlVoiceGender = 0
	// A male voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_MALE SsmlVoiceGender = 1
	// A female voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_FEMALE SsmlVoiceGender = 2
	// A gender-neutral voice.
	SsmlVoiceGender_SSML_VOICE_GENDER_NEUTRAL SsmlVoiceGender = 3
)

var SsmlVoiceGender_name = map[int32]string{
	0: "SSML_VOICE_GENDER_UNSPECIFIED",
	1: "SSML_VOICE_GENDER_MALE",
	2: "SSML_VOICE_GENDER_FEMALE",
	3: "SSML_VOICE_GENDER_NEUTRAL",
}

var SsmlVoiceGender_value = map[string]int32{
	"SSML_VOICE_GENDER_UNSPECIFIED": 0,
	"SSML_VOICE_GENDER_MALE":        1,
	"SSML_VOICE_GENDER_FEMALE":      2,
	"SSML_VOICE_GENDER_NEUTRAL":     3,
}

func (x SsmlVoiceGender) String() string {
	return proto.EnumName(SsmlVoiceGender_name, int32(x))
}

func (SsmlVoiceGender) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{2}
}

// Audio encoding of the output audio format in Text-To-Speech.
type OutputAudioEncoding int32

const (
	// Not specified.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_UNSPECIFIED OutputAudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	// Audio content returned as LINEAR16 also contains a WAV header.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_LINEAR_16 OutputAudioEncoding = 1
	// MP3 audio at 32kbps.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_MP3 OutputAudioEncoding = 2
	// Opus encoded audio wrapped in an ogg container. The result will be a
	// file which can be played natively on Android, and in browsers (at least
	// Chrome and Firefox). The quality of the encoding is considerably higher
	// than MP3 while using approximately the same bitrate.
	OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_OGG_OPUS OutputAudioEncoding = 3
)

var OutputAudioEncoding_name = map[int32]string{
	0: "OUTPUT_AUDIO_ENCODING_UNSPECIFIED",
	1: "OUTPUT_AUDIO_ENCODING_LINEAR_16",
	2: "OUTPUT_AUDIO_ENCODING_MP3",
	3: "OUTPUT_AUDIO_ENCODING_OGG_OPUS",
}

var OutputAudioEncoding_value = map[string]int32{
	"OUTPUT_AUDIO_ENCODING_UNSPECIFIED": 0,
	"OUTPUT_AUDIO_ENCODING_LINEAR_16":   1,
	"OUTPUT_AUDIO_ENCODING_MP3":         2,
	"OUTPUT_AUDIO_ENCODING_OGG_OPUS":    3,
}

func (x OutputAudioEncoding) String() string {
	return proto.EnumName(OutputAudioEncoding_name, int32(x))
}

func (OutputAudioEncoding) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{3}
}

// Hints for the speech recognizer to help with recognition in a specific
// conversation state.
type SpeechContext struct {
	// Optional. A list of strings containing words and phrases that the speech
	// recognizer should recognize with higher likelihood.
	//
	// This list can be used to:
	// * improve accuracy for words and phrases you expect the user to say,
	//   e.g. typical commands for your Dialogflow agent
	// * add additional words to the speech recognizer vocabulary
	// * ...
	//
	// See the [Cloud Speech
	// documentation](https://cloud.google.com/speech-to-text/quotas) for usage
	// limits.
	Phrases []string `protobuf:"bytes,1,rep,name=phrases,proto3" json:"phrases,omitempty"`
	// Optional. Boost for this context compared to other contexts:
	// * If the boost is positive, Dialogflow will increase the probability that
	//   the phrases in this context are recognized over similar sounding phrases.
	// * If the boost is unspecified or non-positive, Dialogflow will not apply
	//   any boost.
	//
	// Dialogflow recommends that you use boosts in the range (0, 20] and that you
	// find a value that fits your use case with binary search.
	Boost                float32  `protobuf:"fixed32,2,opt,name=boost,proto3" json:"boost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpeechContext) Reset()         { *m = SpeechContext{} }
func (m *SpeechContext) String() string { return proto.CompactTextString(m) }
func (*SpeechContext) ProtoMessage()    {}
func (*SpeechContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{0}
}

func (m *SpeechContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeechContext.Unmarshal(m, b)
}
func (m *SpeechContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeechContext.Marshal(b, m, deterministic)
}
func (m *SpeechContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeechContext.Merge(m, src)
}
func (m *SpeechContext) XXX_Size() int {
	return xxx_messageInfo_SpeechContext.Size(m)
}
func (m *SpeechContext) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeechContext.DiscardUnknown(m)
}

var xxx_messageInfo_SpeechContext proto.InternalMessageInfo

func (m *SpeechContext) GetPhrases() []string {
	if m != nil {
		return m.Phrases
	}
	return nil
}

func (m *SpeechContext) GetBoost() float32 {
	if m != nil {
		return m.Boost
	}
	return 0
}

// Information for a word recognized by the speech recognizer.
type SpeechWordInfo struct {
	// The word this info is for.
	Word string `protobuf:"bytes,3,opt,name=word,proto3" json:"word,omitempty"`
	// Time offset relative to the beginning of the audio that corresponds to the
	// start of the spoken word. This is an experimental feature and the accuracy
	// of the time offset can vary.
	StartOffset *duration.Duration `protobuf:"bytes,1,opt,name=start_offset,json=startOffset,proto3" json:"start_offset,omitempty"`
	// Time offset relative to the beginning of the audio that corresponds to the
	// end of the spoken word. This is an experimental feature and the accuracy of
	// the time offset can vary.
	EndOffset *duration.Duration `protobuf:"bytes,2,opt,name=end_offset,json=endOffset,proto3" json:"end_offset,omitempty"`
	// The Speech confidence between 0.0 and 1.0 for this word. A higher number
	// indicates an estimated greater likelihood that the recognized word is
	// correct. The default of 0.0 is a sentinel value indicating that confidence
	// was not set.
	//
	// This field is not guaranteed to be fully stable over time for the same
	// audio input. Users should also not rely on it to always be provided.
	Confidence           float32  `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SpeechWordInfo) Reset()         { *m = SpeechWordInfo{} }
func (m *SpeechWordInfo) String() string { return proto.CompactTextString(m) }
func (*SpeechWordInfo) ProtoMessage()    {}
func (*SpeechWordInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{1}
}

func (m *SpeechWordInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpeechWordInfo.Unmarshal(m, b)
}
func (m *SpeechWordInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpeechWordInfo.Marshal(b, m, deterministic)
}
func (m *SpeechWordInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpeechWordInfo.Merge(m, src)
}
func (m *SpeechWordInfo) XXX_Size() int {
	return xxx_messageInfo_SpeechWordInfo.Size(m)
}
func (m *SpeechWordInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpeechWordInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpeechWordInfo proto.InternalMessageInfo

func (m *SpeechWordInfo) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *SpeechWordInfo) GetStartOffset() *duration.Duration {
	if m != nil {
		return m.StartOffset
	}
	return nil
}

func (m *SpeechWordInfo) GetEndOffset() *duration.Duration {
	if m != nil {
		return m.EndOffset
	}
	return nil
}

func (m *SpeechWordInfo) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

// Instructs the speech recognizer how to process the audio content.
type InputAudioConfig struct {
	// Required. Audio encoding of the audio content to process.
	AudioEncoding AudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=google.cloud.dialogflow.v2.AudioEncoding" json:"audio_encoding,omitempty"`
	// Required. Sample rate (in Hertz) of the audio content sent in the query.
	// Refer to
	// [Cloud Speech API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics) for
	// more details.
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// Required. The language of the supplied audio. Dialogflow does not do
	// translations. See [Language
	// Support](https://cloud.google.com/dialogflow/docs/reference/language)
	// for a list of the currently supported language codes. Note that queries in
	// the same session do not necessarily need to specify the same language.
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. If `true`, Dialogflow returns [SpeechWordInfo][google.cloud.dialogflow.v2.SpeechWordInfo] in
	// [StreamingRecognitionResult][google.cloud.dialogflow.v2.StreamingRecognitionResult] with information about the recognized speech
	// words, e.g. start and end time offsets. If false or unspecified, Speech
	// doesn't return any word-level information.
	EnableWordInfo bool `protobuf:"varint,13,opt,name=enable_word_info,json=enableWordInfo,proto3" json:"enable_word_info,omitempty"`
	// Optional. A list of strings containing words and phrases that the speech
	// recognizer should recognize with higher likelihood.
	//
	// See [the Cloud Speech
	// documentation](https://cloud.google.com/speech-to-text/docs/basics#phrase-hints)
	// for more details.
	//
	// This field is deprecated. Please use [speech_contexts]() instead. If you
	// specify both [phrase_hints]() and [speech_contexts](), Dialogflow will
	// treat the [phrase_hints]() as a single additional [SpeechContext]().
	PhraseHints []string `protobuf:"bytes,4,rep,name=phrase_hints,json=phraseHints,proto3" json:"phrase_hints,omitempty"` // Deprecated: Do not use.
	// Optional. Context information to assist speech recognition.
	//
	// See [the Cloud Speech
	// documentation](https://cloud.google.com/speech-to-text/docs/basics#phrase-hints)
	// for more details.
	SpeechContexts []*SpeechContext `protobuf:"bytes,11,rep,name=speech_contexts,json=speechContexts,proto3" json:"speech_contexts,omitempty"`
	// Optional. Which Speech model to select for the given request. Select the
	// model best suited to your domain to get best results. If a model is not
	// explicitly specified, then we auto-select a model based on the parameters
	// in the InputAudioConfig.
	// If enhanced speech model is enabled for the agent and an enhanced
	// version of the specified model for the language does not exist, then the
	// speech is recognized using the standard version of the specified model.
	// Refer to
	// [Cloud Speech API
	// documentation](https://cloud.google.com/speech-to-text/docs/basics#select-model)
	// for more details.
	Model string `protobuf:"bytes,7,opt,name=model,proto3" json:"model,omitempty"`
	// Optional. Which variant of the [Speech model][google.cloud.dialogflow.v2.InputAudioConfig.model] to use.
	ModelVariant SpeechModelVariant `protobuf:"varint,10,opt,name=model_variant,json=modelVariant,proto3,enum=google.cloud.dialogflow.v2.SpeechModelVariant" json:"model_variant,omitempty"`
	// Optional. If `false` (default), recognition does not cease until the
	// client closes the stream.
	// If `true`, the recognizer will detect a single spoken utterance in input
	// audio. Recognition ceases when it detects the audio's voice has
	// stopped or paused. In this case, once a detected intent is received, the
	// client should close the stream and start a new request with a new stream as
	// needed.
	// Note: This setting is relevant only for streaming methods.
	// Note: When specified, InputAudioConfig.single_utterance takes precedence
	// over StreamingDetectIntentRequest.single_utterance.
	SingleUtterance      bool     `protobuf:"varint,8,opt,name=single_utterance,json=singleUtterance,proto3" json:"single_utterance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputAudioConfig) Reset()         { *m = InputAudioConfig{} }
func (m *InputAudioConfig) String() string { return proto.CompactTextString(m) }
func (*InputAudioConfig) ProtoMessage()    {}
func (*InputAudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{2}
}

func (m *InputAudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputAudioConfig.Unmarshal(m, b)
}
func (m *InputAudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputAudioConfig.Marshal(b, m, deterministic)
}
func (m *InputAudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputAudioConfig.Merge(m, src)
}
func (m *InputAudioConfig) XXX_Size() int {
	return xxx_messageInfo_InputAudioConfig.Size(m)
}
func (m *InputAudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_InputAudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_InputAudioConfig proto.InternalMessageInfo

func (m *InputAudioConfig) GetAudioEncoding() AudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return AudioEncoding_AUDIO_ENCODING_UNSPECIFIED
}

func (m *InputAudioConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *InputAudioConfig) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *InputAudioConfig) GetEnableWordInfo() bool {
	if m != nil {
		return m.EnableWordInfo
	}
	return false
}

// Deprecated: Do not use.
func (m *InputAudioConfig) GetPhraseHints() []string {
	if m != nil {
		return m.PhraseHints
	}
	return nil
}

func (m *InputAudioConfig) GetSpeechContexts() []*SpeechContext {
	if m != nil {
		return m.SpeechContexts
	}
	return nil
}

func (m *InputAudioConfig) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *InputAudioConfig) GetModelVariant() SpeechModelVariant {
	if m != nil {
		return m.ModelVariant
	}
	return SpeechModelVariant_SPEECH_MODEL_VARIANT_UNSPECIFIED
}

func (m *InputAudioConfig) GetSingleUtterance() bool {
	if m != nil {
		return m.SingleUtterance
	}
	return false
}

// Description of which voice to use for speech synthesis.
type VoiceSelectionParams struct {
	// Optional. The name of the voice. If not set, the service will choose a
	// voice based on the other parameters such as language_code and
	// [ssml_gender][google.cloud.dialogflow.v2.VoiceSelectionParams.ssml_gender].
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. The preferred gender of the voice. If not set, the service will
	// choose a voice based on the other parameters such as language_code and
	// [name][google.cloud.dialogflow.v2.VoiceSelectionParams.name]. Note that this is only a preference, not requirement. If a
	// voice of the appropriate gender is not available, the synthesizer should
	// substitute a voice with a different gender rather than failing the request.
	SsmlGender           SsmlVoiceGender `protobuf:"varint,2,opt,name=ssml_gender,json=ssmlGender,proto3,enum=google.cloud.dialogflow.v2.SsmlVoiceGender" json:"ssml_gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *VoiceSelectionParams) Reset()         { *m = VoiceSelectionParams{} }
func (m *VoiceSelectionParams) String() string { return proto.CompactTextString(m) }
func (*VoiceSelectionParams) ProtoMessage()    {}
func (*VoiceSelectionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{3}
}

func (m *VoiceSelectionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoiceSelectionParams.Unmarshal(m, b)
}
func (m *VoiceSelectionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoiceSelectionParams.Marshal(b, m, deterministic)
}
func (m *VoiceSelectionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoiceSelectionParams.Merge(m, src)
}
func (m *VoiceSelectionParams) XXX_Size() int {
	return xxx_messageInfo_VoiceSelectionParams.Size(m)
}
func (m *VoiceSelectionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_VoiceSelectionParams.DiscardUnknown(m)
}

var xxx_messageInfo_VoiceSelectionParams proto.InternalMessageInfo

func (m *VoiceSelectionParams) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VoiceSelectionParams) GetSsmlGender() SsmlVoiceGender {
	if m != nil {
		return m.SsmlGender
	}
	return SsmlVoiceGender_SSML_VOICE_GENDER_UNSPECIFIED
}

// Configuration of how speech should be synthesized.
type SynthesizeSpeechConfig struct {
	// Optional. Speaking rate/speed, in the range [0.25, 4.0]. 1.0 is the normal
	// native speed supported by the specific voice. 2.0 is twice as fast, and
	// 0.5 is half as fast. If unset(0.0), defaults to the native 1.0 speed. Any
	// other values < 0.25 or > 4.0 will return an error.
	SpeakingRate float64 `protobuf:"fixed64,1,opt,name=speaking_rate,json=speakingRate,proto3" json:"speaking_rate,omitempty"`
	// Optional. Speaking pitch, in the range [-20.0, 20.0]. 20 means increase 20
	// semitones from the original pitch. -20 means decrease 20 semitones from the
	// original pitch.
	Pitch float64 `protobuf:"fixed64,2,opt,name=pitch,proto3" json:"pitch,omitempty"`
	// Optional. Volume gain (in dB) of the normal native volume supported by the
	// specific voice, in the range [-96.0, 16.0]. If unset, or set to a value of
	// 0.0 (dB), will play at normal native signal amplitude. A value of -6.0 (dB)
	// will play at approximately half the amplitude of the normal native signal
	// amplitude. A value of +6.0 (dB) will play at approximately twice the
	// amplitude of the normal native signal amplitude. We strongly recommend not
	// to exceed +10 (dB) as there's usually no effective increase in loudness for
	// any value greater than that.
	VolumeGainDb float64 `protobuf:"fixed64,3,opt,name=volume_gain_db,json=volumeGainDb,proto3" json:"volume_gain_db,omitempty"`
	// Optional. An identifier which selects 'audio effects' profiles that are
	// applied on (post synthesized) text to speech. Effects are applied on top of
	// each other in the order they are given.
	EffectsProfileId []string `protobuf:"bytes,5,rep,name=effects_profile_id,json=effectsProfileId,proto3" json:"effects_profile_id,omitempty"`
	// Optional. The desired voice of the synthesized audio.
	Voice                *VoiceSelectionParams `protobuf:"bytes,4,opt,name=voice,proto3" json:"voice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SynthesizeSpeechConfig) Reset()         { *m = SynthesizeSpeechConfig{} }
func (m *SynthesizeSpeechConfig) String() string { return proto.CompactTextString(m) }
func (*SynthesizeSpeechConfig) ProtoMessage()    {}
func (*SynthesizeSpeechConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{4}
}

func (m *SynthesizeSpeechConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynthesizeSpeechConfig.Unmarshal(m, b)
}
func (m *SynthesizeSpeechConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynthesizeSpeechConfig.Marshal(b, m, deterministic)
}
func (m *SynthesizeSpeechConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynthesizeSpeechConfig.Merge(m, src)
}
func (m *SynthesizeSpeechConfig) XXX_Size() int {
	return xxx_messageInfo_SynthesizeSpeechConfig.Size(m)
}
func (m *SynthesizeSpeechConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SynthesizeSpeechConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SynthesizeSpeechConfig proto.InternalMessageInfo

func (m *SynthesizeSpeechConfig) GetSpeakingRate() float64 {
	if m != nil {
		return m.SpeakingRate
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetPitch() float64 {
	if m != nil {
		return m.Pitch
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetVolumeGainDb() float64 {
	if m != nil {
		return m.VolumeGainDb
	}
	return 0
}

func (m *SynthesizeSpeechConfig) GetEffectsProfileId() []string {
	if m != nil {
		return m.EffectsProfileId
	}
	return nil
}

func (m *SynthesizeSpeechConfig) GetVoice() *VoiceSelectionParams {
	if m != nil {
		return m.Voice
	}
	return nil
}

// Instructs the speech synthesizer on how to generate the output audio content.
// If this audio config is supplied in a request, it overrides all existing
// text-to-speech settings applied to the agent.
type OutputAudioConfig struct {
	// Required. Audio encoding of the synthesized audio content.
	AudioEncoding OutputAudioEncoding `protobuf:"varint,1,opt,name=audio_encoding,json=audioEncoding,proto3,enum=google.cloud.dialogflow.v2.OutputAudioEncoding" json:"audio_encoding,omitempty"`
	// The synthesis sample rate (in hertz) for this audio. If not
	// provided, then the synthesizer will use the default sample rate based on
	// the audio encoding. If this is different from the voice's natural sample
	// rate, then the synthesizer will honor this request by converting to the
	// desired sample rate (which might result in worse audio quality).
	SampleRateHertz int32 `protobuf:"varint,2,opt,name=sample_rate_hertz,json=sampleRateHertz,proto3" json:"sample_rate_hertz,omitempty"`
	// Configuration of how speech should be synthesized.
	SynthesizeSpeechConfig *SynthesizeSpeechConfig `protobuf:"bytes,3,opt,name=synthesize_speech_config,json=synthesizeSpeechConfig,proto3" json:"synthesize_speech_config,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *OutputAudioConfig) Reset()         { *m = OutputAudioConfig{} }
func (m *OutputAudioConfig) String() string { return proto.CompactTextString(m) }
func (*OutputAudioConfig) ProtoMessage()    {}
func (*OutputAudioConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ff9be2146363af6, []int{5}
}

func (m *OutputAudioConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputAudioConfig.Unmarshal(m, b)
}
func (m *OutputAudioConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputAudioConfig.Marshal(b, m, deterministic)
}
func (m *OutputAudioConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputAudioConfig.Merge(m, src)
}
func (m *OutputAudioConfig) XXX_Size() int {
	return xxx_messageInfo_OutputAudioConfig.Size(m)
}
func (m *OutputAudioConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputAudioConfig.DiscardUnknown(m)
}

var xxx_messageInfo_OutputAudioConfig proto.InternalMessageInfo

func (m *OutputAudioConfig) GetAudioEncoding() OutputAudioEncoding {
	if m != nil {
		return m.AudioEncoding
	}
	return OutputAudioEncoding_OUTPUT_AUDIO_ENCODING_UNSPECIFIED
}

func (m *OutputAudioConfig) GetSampleRateHertz() int32 {
	if m != nil {
		return m.SampleRateHertz
	}
	return 0
}

func (m *OutputAudioConfig) GetSynthesizeSpeechConfig() *SynthesizeSpeechConfig {
	if m != nil {
		return m.SynthesizeSpeechConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.dialogflow.v2.AudioEncoding", AudioEncoding_name, AudioEncoding_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.SpeechModelVariant", SpeechModelVariant_name, SpeechModelVariant_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.SsmlVoiceGender", SsmlVoiceGender_name, SsmlVoiceGender_value)
	proto.RegisterEnum("google.cloud.dialogflow.v2.OutputAudioEncoding", OutputAudioEncoding_name, OutputAudioEncoding_value)
	proto.RegisterType((*SpeechContext)(nil), "google.cloud.dialogflow.v2.SpeechContext")
	proto.RegisterType((*SpeechWordInfo)(nil), "google.cloud.dialogflow.v2.SpeechWordInfo")
	proto.RegisterType((*InputAudioConfig)(nil), "google.cloud.dialogflow.v2.InputAudioConfig")
	proto.RegisterType((*VoiceSelectionParams)(nil), "google.cloud.dialogflow.v2.VoiceSelectionParams")
	proto.RegisterType((*SynthesizeSpeechConfig)(nil), "google.cloud.dialogflow.v2.SynthesizeSpeechConfig")
	proto.RegisterType((*OutputAudioConfig)(nil), "google.cloud.dialogflow.v2.OutputAudioConfig")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2/audio_config.proto", fileDescriptor_3ff9be2146363af6)
}

var fileDescriptor_3ff9be2146363af6 = []byte{
	// 1163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdf, 0x6e, 0x1a, 0xc7,
	0x17, 0xfe, 0x2d, 0x98, 0x24, 0x3e, 0x80, 0xbd, 0x99, 0xe4, 0xe7, 0xac, 0xdd, 0xc4, 0x21, 0x24,
	0x91, 0x88, 0xdb, 0x40, 0x4b, 0xa4, 0xaa, 0x52, 0x2b, 0x55, 0x0b, 0xbb, 0xc6, 0x48, 0xfc, 0xd3,
	0x02, 0x76, 0xff, 0x5c, 0x8c, 0x06, 0x76, 0x58, 0x56, 0x5d, 0x66, 0xd0, 0xce, 0x42, 0xd2, 0x5c,
	0xf6, 0xb2, 0x52, 0x9f, 0xa1, 0x52, 0xd5, 0xab, 0x3e, 0x43, 0x1f, 0x26, 0x6f, 0xd0, 0xdb, 0x4a,
	0xbd, 0xa9, 0x66, 0x16, 0x6c, 0x8c, 0x09, 0xb9, 0xe8, 0xdd, 0x9c, 0xf3, 0x9d, 0x73, 0xe6, 0xcc,
	0xf9, 0xbe, 0xa3, 0x5d, 0x78, 0xe9, 0x71, 0xee, 0x05, 0xb4, 0x34, 0x0c, 0xf8, 0xcc, 0x2d, 0xb9,
	0x3e, 0x09, 0xb8, 0x37, 0x0a, 0xf8, 0xeb, 0xd2, 0xbc, 0x5c, 0x22, 0x33, 0xd7, 0xe7, 0x78, 0xc8,
	0xd9, 0xc8, 0xf7, 0x8a, 0xd3, 0x90, 0x47, 0x1c, 0x1d, 0xc5, 0xe1, 0x45, 0x15, 0x5e, 0xbc, 0x0a,
	0x2f, 0xce, 0xcb, 0x47, 0x8f, 0x17, 0xa5, 0xc8, 0xd4, 0x2f, 0x8d, 0x7c, 0x1a, 0xb8, 0x78, 0x40,
	0xc7, 0x64, 0xee, 0xf3, 0x30, 0x4e, 0x3e, 0x3a, 0x5c, 0x09, 0x08, 0xa9, 0xe0, 0xb3, 0x70, 0x48,
	0x17, 0xd0, 0xf1, 0x02, 0x52, 0xd6, 0x60, 0x36, 0x2a, 0xb9, 0xb3, 0x90, 0x44, 0x3e, 0x67, 0x0b,
	0xfc, 0xe1, 0x4a, 0x2a, 0x61, 0x8c, 0x47, 0x0a, 0x14, 0x31, 0x9a, 0xff, 0x1a, 0xb2, 0xdd, 0x29,
	0xa5, 0xc3, 0x71, 0x95, 0xb3, 0x88, 0xbe, 0x89, 0x90, 0x01, 0xb7, 0xa7, 0xe3, 0x90, 0x08, 0x2a,
	0x0c, 0x2d, 0x97, 0x2c, 0xec, 0x3a, 0x4b, 0x13, 0xdd, 0x87, 0xd4, 0x80, 0x73, 0x11, 0x19, 0x89,
	0x9c, 0x56, 0x48, 0x38, 0xb1, 0x91, 0xff, 0x53, 0x83, 0xbd, 0xb8, 0xc2, 0x05, 0x0f, 0xdd, 0x3a,
	0x1b, 0x71, 0x84, 0x60, 0xe7, 0x35, 0x0f, 0x5d, 0x23, 0x99, 0xd3, 0x0a, 0xbb, 0x8e, 0x3a, 0xa3,
	0xaf, 0x20, 0x23, 0x22, 0x12, 0x46, 0x98, 0x8f, 0x46, 0x82, 0x46, 0x86, 0x96, 0xd3, 0x0a, 0xe9,
	0xf2, 0x61, 0x71, 0x31, 0x94, 0x65, 0xf3, 0x45, 0x6b, 0xd1, 0xbc, 0x93, 0x56, 0xe1, 0x6d, 0x15,
	0x8d, 0xbe, 0x00, 0xa0, 0xcc, 0x5d, 0xe6, 0x26, 0x3e, 0x94, 0xbb, 0x4b, 0x99, 0xbb, 0xc8, 0x3c,
	0x06, 0x50, 0x2c, 0xb8, 0x94, 0x0d, 0xa9, 0xb1, 0xa3, 0x3a, 0x5f, 0xf1, 0xe4, 0x7f, 0xda, 0x01,
	0xbd, 0xce, 0xa6, 0xb3, 0xc8, 0x94, 0x8c, 0x55, 0x15, 0x61, 0xa8, 0x03, 0x7b, 0x31, 0x81, 0x94,
	0x0d, 0xb9, 0xeb, 0x33, 0x4f, 0xb5, 0xbb, 0x57, 0x7e, 0x51, 0x7c, 0x3f, 0x87, 0x45, 0x55, 0xc0,
	0x5e, 0x24, 0x38, 0x59, 0xb2, 0x6a, 0xa2, 0x13, 0xb8, 0x2b, 0xc8, 0x64, 0x1a, 0x50, 0x1c, 0x92,
	0x88, 0xe2, 0x31, 0x0d, 0xa3, 0xb7, 0xea, 0x1d, 0x29, 0x67, 0x3f, 0x06, 0x1c, 0x12, 0xd1, 0x33,
	0xe9, 0x46, 0x4f, 0x21, 0x1b, 0x10, 0xe6, 0xcd, 0x88, 0x47, 0xf1, 0x90, 0xbb, 0x74, 0x31, 0xc7,
	0xcc, 0xd2, 0x59, 0xe5, 0x2e, 0x45, 0x2f, 0x41, 0xa7, 0x8c, 0x0c, 0x02, 0x8a, 0xe5, 0x78, 0xb1,
	0xcf, 0x46, 0xdc, 0xc8, 0xe6, 0xb4, 0xc2, 0x9d, 0x4a, 0xf2, 0x9d, 0xa9, 0x39, 0x7b, 0x31, 0x78,
	0x49, 0x49, 0x01, 0x32, 0x31, 0x8d, 0x78, 0xec, 0xb3, 0x48, 0x18, 0x3b, 0x92, 0xda, 0x4a, 0xea,
	0x9d, 0xa9, 0x19, 0x9a, 0x93, 0x8e, 0xa1, 0x33, 0x89, 0xa0, 0x0b, 0xd8, 0x17, 0x8a, 0x4e, 0xa9,
	0x5e, 0xa9, 0x08, 0x61, 0xa4, 0x73, 0xc9, 0x42, 0x7a, 0xfb, 0xe3, 0xaf, 0x69, 0x68, 0xd1, 0x82,
	0x58, 0xf5, 0x09, 0x74, 0x08, 0xa9, 0x09, 0x77, 0x69, 0x60, 0xdc, 0x96, 0xcf, 0x89, 0x63, 0x62,
	0x0f, 0xea, 0x42, 0x56, 0x1d, 0xf0, 0x9c, 0x84, 0x3e, 0x61, 0x91, 0x01, 0x6a, 0xdc, 0xc5, 0x0f,
	0xdf, 0xd8, 0x94, 0x69, 0xe7, 0x71, 0x96, 0x93, 0x99, 0xac, 0x58, 0xe8, 0x05, 0xe8, 0xc2, 0x67,
	0x5e, 0x40, 0xf1, 0x2c, 0x8a, 0x68, 0x48, 0x24, 0xff, 0x77, 0xe4, 0x84, 0x9c, 0xfd, 0xd8, 0xdf,
	0x5f, 0xba, 0xf3, 0x6f, 0xe0, 0xfe, 0x39, 0xf7, 0x87, 0xb4, 0x4b, 0x03, 0x3a, 0x94, 0x0a, 0xea,
	0x90, 0x90, 0x4c, 0x84, 0x14, 0x32, 0x23, 0x13, 0xaa, 0xd8, 0xdf, 0x75, 0xd4, 0x19, 0x35, 0x20,
	0x2d, 0xc4, 0x24, 0xc0, 0x1e, 0x65, 0x2e, 0x0d, 0x15, 0x87, 0x7b, 0xe5, 0x8f, 0xb7, 0x76, 0x2a,
	0x26, 0x81, 0x2a, 0x5f, 0x53, 0x29, 0x0e, 0xc8, 0xfc, 0xf8, 0x9c, 0xff, 0x4b, 0x83, 0x83, 0xee,
	0x8f, 0x2c, 0x1a, 0x53, 0xe1, 0xbf, 0xa5, 0x97, 0x53, 0x94, 0x22, 0x7c, 0x0a, 0x59, 0x31, 0xa5,
	0xe4, 0x07, 0x9f, 0x79, 0x4a, 0x34, 0xaa, 0x0b, 0xcd, 0xc9, 0x2c, 0x9d, 0x52, 0x30, 0x72, 0x27,
	0xa7, 0x7e, 0x34, 0x1c, 0xab, 0x3e, 0x34, 0x27, 0x36, 0xd0, 0x33, 0xd8, 0x9b, 0xf3, 0x60, 0x36,
	0xa1, 0xd8, 0x23, 0x3e, 0xc3, 0xee, 0x40, 0x49, 0x48, 0x73, 0x32, 0xb1, 0xb7, 0x46, 0x7c, 0x66,
	0x0d, 0xd0, 0x27, 0x80, 0xe8, 0x68, 0x44, 0x87, 0x91, 0xc0, 0xd3, 0x90, 0x8f, 0xfc, 0x80, 0x62,
	0xdf, 0x35, 0x52, 0x6a, 0xe9, 0xf5, 0x05, 0xd2, 0x89, 0x81, 0xba, 0x8b, 0x4e, 0x21, 0x35, 0x97,
	0x8f, 0x50, 0x3b, 0x94, 0x2e, 0x7f, 0xba, 0xed, 0xc5, 0x9b, 0x86, 0xe9, 0xc4, 0xe9, 0xf9, 0x9f,
	0x13, 0x70, 0xb7, 0x3d, 0x8b, 0xd6, 0x36, 0xee, 0xfb, 0xf7, 0x6c, 0x5c, 0x69, 0xdb, 0x35, 0x2b,
	0x65, 0x96, 0x8b, 0x26, 0x65, 0x95, 0xf8, 0x2f, 0xcb, 0x17, 0x80, 0x21, 0x2e, 0xf9, 0xc0, 0x57,
	0x9b, 0x30, 0xf2, 0x3d, 0x35, 0xc4, 0x74, 0xb9, 0xbc, 0x95, 0xeb, 0x8d, 0x5c, 0x3a, 0x07, 0x62,
	0xa3, 0xff, 0xe4, 0x1f, 0x0d, 0xb2, 0xd7, 0xfa, 0x47, 0xc7, 0x70, 0x64, 0xf6, 0xad, 0x7a, 0x1b,
	0xdb, 0xad, 0x6a, 0xdb, 0xaa, 0xb7, 0x6a, 0xb8, 0xdf, 0xea, 0x76, 0xec, 0x6a, 0xfd, 0xb4, 0x6e,
	0x5b, 0xfa, 0xff, 0xd0, 0x43, 0x30, 0xd6, 0xf0, 0x46, 0xbd, 0x65, 0x9b, 0x0e, 0xfe, 0xec, 0x73,
	0x5d, 0x43, 0x0f, 0xe0, 0xde, 0x1a, 0x7a, 0xda, 0x30, 0xab, 0x7a, 0x02, 0x19, 0x70, 0x7f, 0x0d,
	0x68, 0xf6, 0x1b, 0xe6, 0x85, 0x9e, 0x44, 0x07, 0x80, 0xd6, 0x10, 0xb3, 0xe9, 0xe8, 0x3b, 0xe8,
	0x10, 0xfe, 0x7f, 0xd3, 0x8f, 0x2f, 0x2a, 0x7a, 0x0a, 0x7d, 0x04, 0x0f, 0xd6, 0xa0, 0x76, 0xad,
	0x86, 0xdb, 0x9d, 0x7e, 0x57, 0xbf, 0x85, 0x5e, 0xc0, 0xf3, 0x35, 0xb0, 0xdb, 0xb1, 0xed, 0x6f,
	0xf0, 0x45, 0xbd, 0x77, 0x86, 0xcf, 0x6c, 0xd3, 0xb2, 0x1d, 0x5c, 0xf9, 0xb6, 0x67, 0xeb, 0xb7,
	0x4f, 0xe6, 0x80, 0x6e, 0x6e, 0x31, 0x7a, 0x06, 0x39, 0x99, 0x51, 0x3d, 0xc3, 0xcd, 0xb6, 0x65,
	0x37, 0xf0, 0xb9, 0xe9, 0xd4, 0xcd, 0x56, 0x6f, 0x6d, 0x0e, 0x07, 0x80, 0xfa, 0x5d, 0x1b, 0x57,
	0xec, 0x6e, 0x0f, 0x9b, 0xe7, 0x66, 0xbd, 0x61, 0x56, 0x1a, 0xb6, 0xae, 0x21, 0x1d, 0x32, 0xd2,
	0xdf, 0xed, 0x99, 0x2d, 0xcb, 0x74, 0x2c, 0x3d, 0xb1, 0xf4, 0xd8, 0xad, 0x33, 0xb3, 0x55, 0xb5,
	0x2d, 0x3d, 0x79, 0xf2, 0x8b, 0x06, 0xfb, 0x6b, 0x4b, 0x89, 0x9e, 0xc0, 0xa3, 0x6e, 0xb7, 0xd9,
	0xc0, 0xe7, 0xed, 0x7a, 0xd5, 0xc6, 0x35, 0xbb, 0x25, 0xfb, 0xbc, 0x7e, 0xe5, 0x11, 0x1c, 0xdc,
	0x0c, 0x69, 0x9a, 0xea, 0xda, 0x87, 0x60, 0xdc, 0xc4, 0x4e, 0x6d, 0x85, 0x26, 0xd0, 0x23, 0x38,
	0xbc, 0x89, 0xb6, 0xec, 0x7e, 0xcf, 0x31, 0x1b, 0x7a, 0xf2, 0xe4, 0x77, 0x0d, 0xee, 0x6d, 0xd0,
	0x32, 0x7a, 0x0e, 0x4f, 0xda, 0xfd, 0x5e, 0xa7, 0xdf, 0xc3, 0x5b, 0x25, 0xf1, 0x14, 0x1e, 0x6f,
	0x0e, 0x5b, 0x55, 0xc6, 0x23, 0x38, 0xdc, 0x1c, 0xd4, 0xec, 0xbc, 0xd2, 0x13, 0x28, 0x0f, 0xc7,
	0x9b, 0xe1, 0x4b, 0x66, 0x93, 0x95, 0x5f, 0x35, 0x38, 0x1e, 0xf2, 0xc9, 0x16, 0xf9, 0x57, 0xf4,
	0x95, 0x9d, 0xee, 0xc8, 0xcf, 0x72, 0x47, 0xfb, 0xce, 0x5a, 0xc4, 0x7b, 0x5c, 0x7e, 0xc0, 0x8a,
	0x3c, 0xf4, 0x4a, 0x1e, 0x65, 0xea, 0xa3, 0x5d, 0x8a, 0x21, 0x32, 0xf5, 0xc5, 0xa6, 0xbf, 0xa8,
	0x2f, 0xaf, 0xac, 0xbf, 0x35, 0xed, 0xb7, 0x44, 0xc2, 0x3a, 0xfd, 0x23, 0x71, 0x54, 0x8b, 0xcb,
	0x55, 0xd5, 0xf5, 0xd6, 0xd5, 0xf5, 0xe7, 0xe5, 0xc1, 0x2d, 0x55, 0xf5, 0xd5, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5f, 0x98, 0x4d, 0xe2, 0x9a, 0x09, 0x00, 0x00,
}
