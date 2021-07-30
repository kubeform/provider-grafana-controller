/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

func GetEncoder() map[string]jsoniter.ValEncoder {
	return map[string]jsoniter.ValEncoder{
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettings{}).Type1()):                        MonitoringCheckSpecSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDns{}).Type1()):                     MonitoringCheckSpecSettingsDnsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}).Type1()):    MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}).Type1()): MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttp{}).Type1()):                    MonitoringCheckSpecSettingsHttpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpBasicAuth{}).Type1()):           MonitoringCheckSpecSettingsHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpTlsConfig{}).Type1()):           MonitoringCheckSpecSettingsHttpTlsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsPing{}).Type1()):                    MonitoringCheckSpecSettingsPingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcp{}).Type1()):                     MonitoringCheckSpecSettingsTcpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcpTlsConfig{}).Type1()):            MonitoringCheckSpecSettingsTcpTlsConfigCodec{},
	}
}

func GetDecoder() map[string]jsoniter.ValDecoder {
	return map[string]jsoniter.ValDecoder{
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettings{}).Type1()):                        MonitoringCheckSpecSettingsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDns{}).Type1()):                     MonitoringCheckSpecSettingsDnsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}).Type1()):    MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}).Type1()): MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttp{}).Type1()):                    MonitoringCheckSpecSettingsHttpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpBasicAuth{}).Type1()):           MonitoringCheckSpecSettingsHttpBasicAuthCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpTlsConfig{}).Type1()):           MonitoringCheckSpecSettingsHttpTlsConfigCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsPing{}).Type1()):                    MonitoringCheckSpecSettingsPingCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcp{}).Type1()):                     MonitoringCheckSpecSettingsTcpCodec{},
		jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcpTlsConfig{}).Type1()):            MonitoringCheckSpecSettingsTcpTlsConfigCodec{},
	}
}

func getEncodersWithout(typ string) map[string]jsoniter.ValEncoder {
	origMap := GetEncoder()
	delete(origMap, typ)
	return origMap
}

func getDecodersWithout(typ string) map[string]jsoniter.ValDecoder {
	origMap := GetDecoder()
	delete(origMap, typ)
	return origMap
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsCodec struct {
}

func (MonitoringCheckSpecSettingsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettings)(ptr) == nil
}

func (MonitoringCheckSpecSettingsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettings)(ptr)
	var objs []MonitoringCheckSpecSettings
	if obj != nil {
		objs = []MonitoringCheckSpecSettings{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettings{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettings)(ptr) = MonitoringCheckSpecSettings{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettings)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettings)(ptr) = MonitoringCheckSpecSettings{}
			}
		} else {
			*(*MonitoringCheckSpecSettings)(ptr) = MonitoringCheckSpecSettings{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettings

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettings{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettings)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettings)(ptr) = MonitoringCheckSpecSettings{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettings", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsDnsCodec struct {
}

func (MonitoringCheckSpecSettingsDnsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsDns)(ptr) == nil
}

func (MonitoringCheckSpecSettingsDnsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsDns)(ptr)
	var objs []MonitoringCheckSpecSettingsDns
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsDns{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDns{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsDnsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsDns)(ptr) = MonitoringCheckSpecSettingsDns{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsDns)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsDns)(ptr) = MonitoringCheckSpecSettingsDns{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsDns)(ptr) = MonitoringCheckSpecSettingsDns{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsDns

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDns{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsDns)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsDns)(ptr) = MonitoringCheckSpecSettingsDns{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsDns", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec struct {
}

func (MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) == nil
}

func (MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr)
	var objs []MonitoringCheckSpecSettingsDnsValidateAnswerRrs
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsDnsValidateAnswerRrs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsDnsValidateAnswerRrsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsDnsValidateAnswerRrs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsDnsValidateAnswerRrs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsDnsValidateAnswerRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAnswerRrs{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsDnsValidateAnswerRrs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec struct {
}

func (MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) == nil
}

func (MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr)
	var objs []MonitoringCheckSpecSettingsDnsValidateAuthorityRrs
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsDnsValidateAuthorityRrsCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsDnsValidateAuthorityRrs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsDnsValidateAuthorityRrs

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsDnsValidateAuthorityRrs)(ptr) = MonitoringCheckSpecSettingsDnsValidateAuthorityRrs{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsDnsValidateAuthorityRrs", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsHttpCodec struct {
}

func (MonitoringCheckSpecSettingsHttpCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsHttp)(ptr) == nil
}

func (MonitoringCheckSpecSettingsHttpCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsHttp)(ptr)
	var objs []MonitoringCheckSpecSettingsHttp
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsHttp{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttp{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsHttpCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsHttp)(ptr) = MonitoringCheckSpecSettingsHttp{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsHttp)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsHttp)(ptr) = MonitoringCheckSpecSettingsHttp{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsHttp)(ptr) = MonitoringCheckSpecSettingsHttp{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsHttp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsHttp)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsHttp)(ptr) = MonitoringCheckSpecSettingsHttp{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsHttp", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsHttpBasicAuthCodec struct {
}

func (MonitoringCheckSpecSettingsHttpBasicAuthCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) == nil
}

func (MonitoringCheckSpecSettingsHttpBasicAuthCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr)
	var objs []MonitoringCheckSpecSettingsHttpBasicAuth
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsHttpBasicAuth{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpBasicAuth{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsHttpBasicAuthCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = MonitoringCheckSpecSettingsHttpBasicAuth{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = MonitoringCheckSpecSettingsHttpBasicAuth{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = MonitoringCheckSpecSettingsHttpBasicAuth{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsHttpBasicAuth

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpBasicAuth{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsHttpBasicAuth)(ptr) = MonitoringCheckSpecSettingsHttpBasicAuth{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsHttpBasicAuth", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsHttpTlsConfigCodec struct {
}

func (MonitoringCheckSpecSettingsHttpTlsConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) == nil
}

func (MonitoringCheckSpecSettingsHttpTlsConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr)
	var objs []MonitoringCheckSpecSettingsHttpTlsConfig
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsHttpTlsConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpTlsConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsHttpTlsConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = MonitoringCheckSpecSettingsHttpTlsConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsHttpTlsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpTlsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = MonitoringCheckSpecSettingsHttpTlsConfig{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = MonitoringCheckSpecSettingsHttpTlsConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsHttpTlsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsHttpTlsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsHttpTlsConfig)(ptr) = MonitoringCheckSpecSettingsHttpTlsConfig{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsHttpTlsConfig", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsPingCodec struct {
}

func (MonitoringCheckSpecSettingsPingCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsPing)(ptr) == nil
}

func (MonitoringCheckSpecSettingsPingCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsPing)(ptr)
	var objs []MonitoringCheckSpecSettingsPing
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsPing{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsPing{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsPingCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsPing)(ptr) = MonitoringCheckSpecSettingsPing{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsPing

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsPing{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsPing)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsPing)(ptr) = MonitoringCheckSpecSettingsPing{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsPing)(ptr) = MonitoringCheckSpecSettingsPing{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsPing

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsPing{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsPing)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsPing)(ptr) = MonitoringCheckSpecSettingsPing{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsPing", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsTcpCodec struct {
}

func (MonitoringCheckSpecSettingsTcpCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsTcp)(ptr) == nil
}

func (MonitoringCheckSpecSettingsTcpCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsTcp)(ptr)
	var objs []MonitoringCheckSpecSettingsTcp
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsTcp{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcp{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsTcpCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsTcp)(ptr) = MonitoringCheckSpecSettingsTcp{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsTcp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsTcp)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsTcp)(ptr) = MonitoringCheckSpecSettingsTcp{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsTcp)(ptr) = MonitoringCheckSpecSettingsTcp{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsTcp

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcp{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsTcp)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsTcp)(ptr) = MonitoringCheckSpecSettingsTcp{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsTcp", "unexpected JSON type")
	}
}

// +k8s:deepcopy-gen=false
type MonitoringCheckSpecSettingsTcpTlsConfigCodec struct {
}

func (MonitoringCheckSpecSettingsTcpTlsConfigCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) == nil
}

func (MonitoringCheckSpecSettingsTcpTlsConfigCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	obj := (*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr)
	var objs []MonitoringCheckSpecSettingsTcpTlsConfig
	if obj != nil {
		objs = []MonitoringCheckSpecSettingsTcpTlsConfig{*obj}
	}

	jsonit := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		TagKey:                 "tf",
		TypeEncoders:           getEncodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcpTlsConfig{}).Type1())),
	}.Froze()

	byt, _ := jsonit.Marshal(objs)

	stream.Write(byt)
}

func (MonitoringCheckSpecSettingsTcpTlsConfigCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = MonitoringCheckSpecSettingsTcpTlsConfig{}
		return
	case jsoniter.ArrayValue:
		objsByte := iter.SkipAndReturnBytes()
		if len(objsByte) > 0 {
			var objs []MonitoringCheckSpecSettingsTcpTlsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcpTlsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objsByte, &objs)

			if len(objs) > 0 {
				*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = objs[0]
			} else {
				*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = MonitoringCheckSpecSettingsTcpTlsConfig{}
			}
		} else {
			*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = MonitoringCheckSpecSettingsTcpTlsConfig{}
		}
	case jsoniter.ObjectValue:
		objByte := iter.SkipAndReturnBytes()
		if len(objByte) > 0 {
			var obj MonitoringCheckSpecSettingsTcpTlsConfig

			jsonit := jsoniter.Config{
				EscapeHTML:             true,
				SortMapKeys:            true,
				ValidateJsonRawMessage: true,
				TagKey:                 "tf",
				TypeDecoders:           getDecodersWithout(jsoniter.MustGetKind(reflect2.TypeOf(MonitoringCheckSpecSettingsTcpTlsConfig{}).Type1())),
			}.Froze()
			jsonit.Unmarshal(objByte, &obj)

			*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = obj
		} else {
			*(*MonitoringCheckSpecSettingsTcpTlsConfig)(ptr) = MonitoringCheckSpecSettingsTcpTlsConfig{}
		}
	default:
		iter.ReportError("decode MonitoringCheckSpecSettingsTcpTlsConfig", "unexpected JSON type")
	}
}
