// Code generated by protoc-gen-microweb. DO NOT EDIT.
// source: proto.proto

package proto

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/ptypes/empty"
)

type webAccountsServiceHandler struct {
	r chi.Router
	h AccountsServiceHandler
}

func (h *webAccountsServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webAccountsServiceHandler) ListAccounts(w http.ResponseWriter, r *http.Request) {

	req := &ListAccountsRequest{}

	resp := &ListAccountsResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.ListAccounts(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webAccountsServiceHandler) GetAccount(w http.ResponseWriter, r *http.Request) {

	req := &GetAccountRequest{}

	resp := &Account{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.GetAccount(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webAccountsServiceHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {

	req := &CreateAccountRequest{}

	resp := &Account{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.CreateAccount(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webAccountsServiceHandler) UpdateAccount(w http.ResponseWriter, r *http.Request) {

	req := &UpdateAccountRequest{}

	resp := &Account{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.UpdateAccount(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webAccountsServiceHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {

	req := &DeleteAccountRequest{}
	resp := &empty.Empty{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.DeleteAccount(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

func RegisterAccountsServiceWeb(r chi.Router, i AccountsServiceHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webAccountsServiceHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("POST", "/api/v0/accounts/accounts-list", handler.ListAccounts)
	r.MethodFunc("GET", "/v0/accounts/{id=*}", handler.GetAccount)
	r.MethodFunc("POST", "/v0/accounts", handler.CreateAccount)
	r.MethodFunc("PATCH", "/v0/accounts/{account.id=*}", handler.UpdateAccount)
	r.MethodFunc("DELETE", "/v0/accounts/{id=*}", handler.DeleteAccount)
}

type webGroupsServiceHandler struct {
	r chi.Router
	h GroupsServiceHandler
}

func (h *webGroupsServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func (h *webGroupsServiceHandler) ListGroups(w http.ResponseWriter, r *http.Request) {

	req := &ListGroupsRequest{}

	resp := &ListGroupsResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.ListGroups(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) GetGroup(w http.ResponseWriter, r *http.Request) {

	req := &GetGroupRequest{}

	resp := &Group{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.GetGroup(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {

	req := &CreateGroupRequest{}

	resp := &Group{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.CreateGroup(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {

	req := &UpdateGroupRequest{}

	resp := &Group{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.UpdateGroup(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) DeleteGroup(w http.ResponseWriter, r *http.Request) {

	req := &DeleteGroupRequest{}
	resp := &empty.Empty{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.DeleteGroup(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusNoContent)
	render.NoContent(w, r)
}

func (h *webGroupsServiceHandler) AddMember(w http.ResponseWriter, r *http.Request) {

	req := &AddMemberRequest{}

	resp := &Group{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.AddMember(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) RemoveMember(w http.ResponseWriter, r *http.Request) {

	req := &RemoveMemberRequest{}

	resp := &Group{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.RemoveMember(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *webGroupsServiceHandler) ListMembers(w http.ResponseWriter, r *http.Request) {

	req := &ListMembersRequest{}

	resp := &ListMembersResponse{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if err := h.h.ListMembers(
		r.Context(),
		req,
		resp,
	); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func RegisterGroupsServiceWeb(r chi.Router, i GroupsServiceHandler, middlewares ...func(http.Handler) http.Handler) {
	handler := &webGroupsServiceHandler{
		r: r,
		h: i,
	}

	r.MethodFunc("GET", "/v0/groups", handler.ListGroups)
	r.MethodFunc("GET", "/v0/groups/{id=*}", handler.GetGroup)
	r.MethodFunc("POST", "/v0/groups", handler.CreateGroup)
	r.MethodFunc("PATCH", "/v0/groups/{group.id=*}", handler.UpdateGroup)
	r.MethodFunc("DELETE", "/v0/groups/{id=*}", handler.DeleteGroup)
	r.MethodFunc("POST", "/v0/groups/{id=*}/members/$ref", handler.AddMember)
	r.MethodFunc("DELETE", "/v0/groups/{id=*}/members/{account_id}/$ref", handler.RemoveMember)
	r.MethodFunc("GET", "/v0/groups/{id=*}/members/$ref", handler.ListMembers)
}

// ListAccountsRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListAccountsRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListAccountsRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListAccountsRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListAccountsRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListAccountsRequest)(nil)

// ListAccountsRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListAccountsRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListAccountsRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListAccountsRequest) UnmarshalJSON(b []byte) error {
	return ListAccountsRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListAccountsRequest)(nil)

// ListAccountsResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListAccountsResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListAccountsResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListAccountsResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListAccountsResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListAccountsResponse)(nil)

// ListAccountsResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListAccountsResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListAccountsResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListAccountsResponse) UnmarshalJSON(b []byte) error {
	return ListAccountsResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListAccountsResponse)(nil)

// GetAccountRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetAccountRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetAccountRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetAccountRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetAccountRequest)(nil)

// GetAccountRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetAccountRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetAccountRequest) UnmarshalJSON(b []byte) error {
	return GetAccountRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetAccountRequest)(nil)

// CreateAccountRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of CreateAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CreateAccountRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *CreateAccountRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := CreateAccountRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*CreateAccountRequest)(nil)

// CreateAccountRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of CreateAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CreateAccountRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *CreateAccountRequest) UnmarshalJSON(b []byte) error {
	return CreateAccountRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*CreateAccountRequest)(nil)

// UpdateAccountRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UpdateAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateAccountRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *UpdateAccountRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := UpdateAccountRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*UpdateAccountRequest)(nil)

// UpdateAccountRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of UpdateAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateAccountRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *UpdateAccountRequest) UnmarshalJSON(b []byte) error {
	return UpdateAccountRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*UpdateAccountRequest)(nil)

// DeleteAccountRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of DeleteAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var DeleteAccountRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *DeleteAccountRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := DeleteAccountRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*DeleteAccountRequest)(nil)

// DeleteAccountRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of DeleteAccountRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var DeleteAccountRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *DeleteAccountRequest) UnmarshalJSON(b []byte) error {
	return DeleteAccountRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*DeleteAccountRequest)(nil)

// AccountJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Account. This struct is safe to replace or modify but
// should not be done so concurrently.
var AccountJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Account) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := AccountJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Account)(nil)

// AccountJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Account. This struct is safe to replace or modify but
// should not be done so concurrently.
var AccountJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Account) UnmarshalJSON(b []byte) error {
	return AccountJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Account)(nil)

// IdentitiesJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Identities. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentitiesJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Identities) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := IdentitiesJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Identities)(nil)

// IdentitiesJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Identities. This struct is safe to replace or modify but
// should not be done so concurrently.
var IdentitiesJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Identities) UnmarshalJSON(b []byte) error {
	return IdentitiesJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Identities)(nil)

// PasswordProfileJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of PasswordProfile. This struct is safe to replace or modify but
// should not be done so concurrently.
var PasswordProfileJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *PasswordProfile) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := PasswordProfileJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*PasswordProfile)(nil)

// PasswordProfileJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of PasswordProfile. This struct is safe to replace or modify but
// should not be done so concurrently.
var PasswordProfileJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *PasswordProfile) UnmarshalJSON(b []byte) error {
	return PasswordProfileJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*PasswordProfile)(nil)

// ListGroupsRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListGroupsRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListGroupsRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListGroupsRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListGroupsRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListGroupsRequest)(nil)

// ListGroupsRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListGroupsRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListGroupsRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListGroupsRequest) UnmarshalJSON(b []byte) error {
	return ListGroupsRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListGroupsRequest)(nil)

// ListGroupsResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListGroupsResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListGroupsResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListGroupsResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListGroupsResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListGroupsResponse)(nil)

// ListGroupsResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListGroupsResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListGroupsResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListGroupsResponse) UnmarshalJSON(b []byte) error {
	return ListGroupsResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListGroupsResponse)(nil)

// GetGroupRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of GetGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetGroupRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *GetGroupRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GetGroupRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*GetGroupRequest)(nil)

// GetGroupRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of GetGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var GetGroupRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *GetGroupRequest) UnmarshalJSON(b []byte) error {
	return GetGroupRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*GetGroupRequest)(nil)

// CreateGroupRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of CreateGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CreateGroupRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *CreateGroupRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := CreateGroupRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*CreateGroupRequest)(nil)

// CreateGroupRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of CreateGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var CreateGroupRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *CreateGroupRequest) UnmarshalJSON(b []byte) error {
	return CreateGroupRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*CreateGroupRequest)(nil)

// UpdateGroupRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of UpdateGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateGroupRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *UpdateGroupRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := UpdateGroupRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*UpdateGroupRequest)(nil)

// UpdateGroupRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of UpdateGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var UpdateGroupRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *UpdateGroupRequest) UnmarshalJSON(b []byte) error {
	return UpdateGroupRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*UpdateGroupRequest)(nil)

// DeleteGroupRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of DeleteGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var DeleteGroupRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *DeleteGroupRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := DeleteGroupRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*DeleteGroupRequest)(nil)

// DeleteGroupRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of DeleteGroupRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var DeleteGroupRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *DeleteGroupRequest) UnmarshalJSON(b []byte) error {
	return DeleteGroupRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*DeleteGroupRequest)(nil)

// AddMemberRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of AddMemberRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var AddMemberRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *AddMemberRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := AddMemberRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*AddMemberRequest)(nil)

// AddMemberRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of AddMemberRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var AddMemberRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *AddMemberRequest) UnmarshalJSON(b []byte) error {
	return AddMemberRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*AddMemberRequest)(nil)

// RemoveMemberRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of RemoveMemberRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var RemoveMemberRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *RemoveMemberRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := RemoveMemberRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*RemoveMemberRequest)(nil)

// RemoveMemberRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of RemoveMemberRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var RemoveMemberRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *RemoveMemberRequest) UnmarshalJSON(b []byte) error {
	return RemoveMemberRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*RemoveMemberRequest)(nil)

// ListMembersRequestJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListMembersRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListMembersRequestJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListMembersRequest) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListMembersRequestJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListMembersRequest)(nil)

// ListMembersRequestJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListMembersRequest. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListMembersRequestJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListMembersRequest) UnmarshalJSON(b []byte) error {
	return ListMembersRequestJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListMembersRequest)(nil)

// ListMembersResponseJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of ListMembersResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListMembersResponseJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *ListMembersResponse) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := ListMembersResponseJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*ListMembersResponse)(nil)

// ListMembersResponseJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of ListMembersResponse. This struct is safe to replace or modify but
// should not be done so concurrently.
var ListMembersResponseJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *ListMembersResponse) UnmarshalJSON(b []byte) error {
	return ListMembersResponseJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*ListMembersResponse)(nil)

// GroupJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of Group. This struct is safe to replace or modify but
// should not be done so concurrently.
var GroupJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *Group) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := GroupJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*Group)(nil)

// GroupJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of Group. This struct is safe to replace or modify but
// should not be done so concurrently.
var GroupJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *Group) UnmarshalJSON(b []byte) error {
	return GroupJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*Group)(nil)

// OnPremisesProvisioningErrorJSONMarshaler describes the default jsonpb.Marshaler used by all
// instances of OnPremisesProvisioningError. This struct is safe to replace or modify but
// should not be done so concurrently.
var OnPremisesProvisioningErrorJSONMarshaler = new(jsonpb.Marshaler)

// MarshalJSON satisfies the encoding/json Marshaler interface. This method
// uses the more correct jsonpb package to correctly marshal the message.
func (m *OnPremisesProvisioningError) MarshalJSON() ([]byte, error) {
	if m == nil {
		return json.Marshal(nil)
	}

	buf := &bytes.Buffer{}

	if err := OnPremisesProvisioningErrorJSONMarshaler.Marshal(buf, m); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

var _ json.Marshaler = (*OnPremisesProvisioningError)(nil)

// OnPremisesProvisioningErrorJSONUnmarshaler describes the default jsonpb.Unmarshaler used by all
// instances of OnPremisesProvisioningError. This struct is safe to replace or modify but
// should not be done so concurrently.
var OnPremisesProvisioningErrorJSONUnmarshaler = new(jsonpb.Unmarshaler)

// UnmarshalJSON satisfies the encoding/json Unmarshaler interface. This method
// uses the more correct jsonpb package to correctly unmarshal the message.
func (m *OnPremisesProvisioningError) UnmarshalJSON(b []byte) error {
	return OnPremisesProvisioningErrorJSONUnmarshaler.Unmarshal(bytes.NewReader(b), m)
}

var _ json.Unmarshaler = (*OnPremisesProvisioningError)(nil)
