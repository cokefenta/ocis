// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: proto.proto

package proto

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"
)

type webBundleServiceHandler struct {
	r chi.Router
	h BundleServiceHandler
}

func (h *webBundleServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webBundleServiceHandler) SaveSettingsBundle(w http.ResponseWriter, r *http.Request) {

	req := &SaveSettingsBundleRequest{}

	resp := &SaveSettingsBundleResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.SaveSettingsBundle(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webBundleServiceHandler) GetSettingsBundle(w http.ResponseWriter, r *http.Request) {

	req := &GetSettingsBundleRequest{}

	resp := &GetSettingsBundleResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.GetSettingsBundle(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webBundleServiceHandler) ListSettingsBundles(w http.ResponseWriter, r *http.Request) {

	req := &ListSettingsBundlesRequest{}

	resp := &ListSettingsBundlesResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.ListSettingsBundles(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func RegisterBundleServiceWeb(r chi.Router, i BundleServiceHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webBundleServiceHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/v0/settings/bundle-save", handler.SaveSettingsBundle)
	r.MethodFunc("POST", "/api/v0/settings/bundle-get", handler.GetSettingsBundle)
	r.MethodFunc("POST", "/api/v0/settings/bundles-list", handler.ListSettingsBundles)
}

type webValueServiceHandler struct {
	r chi.Router
	h ValueServiceHandler
}

func (h *webValueServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webValueServiceHandler) SaveSettingsValue(w http.ResponseWriter, r *http.Request) {

	req := &SaveSettingsValueRequest{}

	resp := &SaveSettingsValueResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.SaveSettingsValue(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webValueServiceHandler) GetSettingsValue(w http.ResponseWriter, r *http.Request) {

	req := &GetSettingsValueRequest{}

	resp := &GetSettingsValueResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.GetSettingsValue(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webValueServiceHandler) ListSettingsValues(w http.ResponseWriter, r *http.Request) {

	req := &ListSettingsValuesRequest{}

	resp := &ListSettingsValuesResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.ListSettingsValues(
		context.Background(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func RegisterValueServiceWeb(r chi.Router, i ValueServiceHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webValueServiceHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/v0/settings/value-save", handler.SaveSettingsValue)
	r.MethodFunc("POST", "/api/v0/settings/value-get", handler.GetSettingsValue)
	r.MethodFunc("POST", "/api/v0/settings/values-list", handler.ListSettingsValues)
}

// SaveSettingsBundleRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SaveSettingsBundleRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsBundleRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SaveSettingsBundleRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SaveSettingsBundleRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SaveSettingsBundleRequest)(nil)

// SaveSettingsBundleRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SaveSettingsBundleRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsBundleRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SaveSettingsBundleRequest) UnmarshalJSON(b []byte) error {
	return SaveSettingsBundleRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SaveSettingsBundleRequest)(nil)

// SaveSettingsBundleResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SaveSettingsBundleResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsBundleResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SaveSettingsBundleResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SaveSettingsBundleResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SaveSettingsBundleResponse)(nil)

// SaveSettingsBundleResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SaveSettingsBundleResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsBundleResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SaveSettingsBundleResponse) UnmarshalJSON(b []byte) error {
	return SaveSettingsBundleResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SaveSettingsBundleResponse)(nil)

// GetSettingsBundleRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetSettingsBundleRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsBundleRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetSettingsBundleRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetSettingsBundleRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetSettingsBundleRequest)(nil)

// GetSettingsBundleRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetSettingsBundleRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsBundleRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetSettingsBundleRequest) UnmarshalJSON(b []byte) error {
	return GetSettingsBundleRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetSettingsBundleRequest)(nil)

// GetSettingsBundleResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetSettingsBundleResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsBundleResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetSettingsBundleResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetSettingsBundleResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetSettingsBundleResponse)(nil)

// GetSettingsBundleResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetSettingsBundleResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsBundleResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetSettingsBundleResponse) UnmarshalJSON(b []byte) error {
	return GetSettingsBundleResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetSettingsBundleResponse)(nil)

// ListSettingsBundlesRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSettingsBundlesRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsBundlesRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSettingsBundlesRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListSettingsBundlesRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSettingsBundlesRequest)(nil)

// ListSettingsBundlesRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSettingsBundlesRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsBundlesRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSettingsBundlesRequest) UnmarshalJSON(b []byte) error {
	return ListSettingsBundlesRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSettingsBundlesRequest)(nil)

// ListSettingsBundlesResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSettingsBundlesResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsBundlesResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSettingsBundlesResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListSettingsBundlesResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSettingsBundlesResponse)(nil)

// ListSettingsBundlesResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSettingsBundlesResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsBundlesResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSettingsBundlesResponse) UnmarshalJSON(b []byte) error {
	return ListSettingsBundlesResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSettingsBundlesResponse)(nil)

// SaveSettingsValueRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SaveSettingsValueRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsValueRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SaveSettingsValueRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SaveSettingsValueRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SaveSettingsValueRequest)(nil)

// SaveSettingsValueRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SaveSettingsValueRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsValueRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SaveSettingsValueRequest) UnmarshalJSON(b []byte) error {
	return SaveSettingsValueRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SaveSettingsValueRequest)(nil)

// SaveSettingsValueResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SaveSettingsValueResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsValueResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SaveSettingsValueResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SaveSettingsValueResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SaveSettingsValueResponse)(nil)

// SaveSettingsValueResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SaveSettingsValueResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var SaveSettingsValueResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SaveSettingsValueResponse) UnmarshalJSON(b []byte) error {
	return SaveSettingsValueResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SaveSettingsValueResponse)(nil)

// GetSettingsValueRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetSettingsValueRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsValueRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetSettingsValueRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetSettingsValueRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetSettingsValueRequest)(nil)

// GetSettingsValueRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetSettingsValueRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsValueRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetSettingsValueRequest) UnmarshalJSON(b []byte) error {
	return GetSettingsValueRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetSettingsValueRequest)(nil)

// GetSettingsValueResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetSettingsValueResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsValueResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetSettingsValueResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetSettingsValueResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetSettingsValueResponse)(nil)

// GetSettingsValueResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetSettingsValueResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetSettingsValueResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetSettingsValueResponse) UnmarshalJSON(b []byte) error {
	return GetSettingsValueResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetSettingsValueResponse)(nil)

// ListSettingsValuesRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSettingsValuesRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsValuesRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSettingsValuesRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListSettingsValuesRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSettingsValuesRequest)(nil)

// ListSettingsValuesRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSettingsValuesRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsValuesRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSettingsValuesRequest) UnmarshalJSON(b []byte) error {
	return ListSettingsValuesRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSettingsValuesRequest)(nil)

// ListSettingsValuesResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListSettingsValuesResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsValuesResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListSettingsValuesResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListSettingsValuesResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListSettingsValuesResponse)(nil)

// ListSettingsValuesResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListSettingsValuesResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListSettingsValuesResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListSettingsValuesResponse) UnmarshalJSON(b []byte) error {
	return ListSettingsValuesResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListSettingsValuesResponse)(nil)

// IdentifierJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Identifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentifierJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Identifier) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := IdentifierJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Identifier)(nil)

// IdentifierJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Identifier. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentifierJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Identifier) UnmarshalJSON(b []byte) error {
	return IdentifierJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Identifier)(nil)

// SettingsBundleJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SettingsBundle. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsBundleJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SettingsBundle) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SettingsBundleJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SettingsBundle)(nil)

// SettingsBundleJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SettingsBundle. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsBundleJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SettingsBundle) UnmarshalJSON(b []byte) error {
	return SettingsBundleJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SettingsBundle)(nil)

// SettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Setting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Setting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Setting)(nil)

// SettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Setting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Setting) UnmarshalJSON(b []byte) error {
	return SettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Setting)(nil)

// IntSettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of IntSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var IntSettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *IntSetting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := IntSettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*IntSetting)(nil)

// IntSettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of IntSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var IntSettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *IntSetting) UnmarshalJSON(b []byte) error {
	return IntSettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*IntSetting)(nil)

// StringSettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of StringSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var StringSettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *StringSetting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := StringSettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*StringSetting)(nil)

// StringSettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of StringSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var StringSettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *StringSetting) UnmarshalJSON(b []byte) error {
	return StringSettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*StringSetting)(nil)

// BoolSettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of BoolSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var BoolSettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *BoolSetting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := BoolSettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*BoolSetting)(nil)

// BoolSettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of BoolSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var BoolSettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *BoolSetting) UnmarshalJSON(b []byte) error {
	return BoolSettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*BoolSetting)(nil)

// SingleChoiceListSettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SingleChoiceListSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SingleChoiceListSettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SingleChoiceListSetting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SingleChoiceListSettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SingleChoiceListSetting)(nil)

// SingleChoiceListSettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SingleChoiceListSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var SingleChoiceListSettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SingleChoiceListSetting) UnmarshalJSON(b []byte) error {
	return SingleChoiceListSettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SingleChoiceListSetting)(nil)

// MultiChoiceListSettingJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of MultiChoiceListSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var MultiChoiceListSettingJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *MultiChoiceListSetting) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := MultiChoiceListSettingJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*MultiChoiceListSetting)(nil)

// MultiChoiceListSettingJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of MultiChoiceListSetting. This struct is safe to replace or modify but
// should not be done so concurrently.
var MultiChoiceListSettingJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *MultiChoiceListSetting) UnmarshalJSON(b []byte) error {
	return MultiChoiceListSettingJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*MultiChoiceListSetting)(nil)

// ListOptionJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListOption) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListOptionJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListOption)(nil)

// ListOptionJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListOption. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListOption) UnmarshalJSON(b []byte) error {
	return ListOptionJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListOption)(nil)

// SettingsValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SettingsValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SettingsValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SettingsValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SettingsValue)(nil)

// SettingsValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SettingsValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SettingsValue) UnmarshalJSON(b []byte) error {
	return SettingsValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SettingsValue)(nil)

// ListValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListValue)(nil)

// ListValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListValue) UnmarshalJSON(b []byte) error {
	return ListValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListValue)(nil)

// ListOptionValueJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListOptionValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionValueJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListOptionValue) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListOptionValueJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListOptionValue)(nil)

// ListOptionValueJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListOptionValue. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListOptionValueJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListOptionValue) UnmarshalJSON(b []byte) error {
	return ListOptionValueJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListOptionValue)(nil)

// SettingsValuesJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of SettingsValues. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsValuesJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *SettingsValues) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := SettingsValuesJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*SettingsValues)(nil)

// SettingsValuesJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of SettingsValues. This struct is safe to replace or modify but
// should not be done so concurrently.
var SettingsValuesJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *SettingsValues) UnmarshalJSON(b []byte) error {
	return SettingsValuesJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*SettingsValues)(nil)
