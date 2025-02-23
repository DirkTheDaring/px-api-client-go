/*
ProxMox VE API

ProxMox VE API

API version: 8.3
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiflat

import (
	"encoding/json"
)

// checks if the CreateAccessTicketRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAccessTicketRequest{}

// CreateAccessTicketRequest struct for CreateAccessTicketRequest
type CreateAccessTicketRequest struct {
	// This parameter is now ignored and assumed to be 1.
	NewFormat *bool `json:"new-format,omitempty"`
	// One-time password for Two-factor authentication.
	Otp *string `json:"otp,omitempty"`
	// The secret password. This can also be a valid ticket.
	Password string `json:"password"`
	// Verify ticket, and check if user have access 'privs' on 'path'
	Path *string `json:"path,omitempty"`
	// Verify ticket, and check if user have access 'privs' on 'path'
	Privs *string `json:"privs,omitempty"`
	// You can optionally pass the realm using this parameter. Normally the realm is simply added to the username <username>@<relam>.
	Realm *string `json:"realm,omitempty"`
	// The signed TFA challenge string the user wants to respond to.
	TfaChallenge *string `json:"tfa-challenge,omitempty"`
	// User name
	Username string `json:"username"`
}

// NewCreateAccessTicketRequest instantiates a new CreateAccessTicketRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessTicketRequest(password string, username string) *CreateAccessTicketRequest {
	this := CreateAccessTicketRequest{}
	this.Password = password
	this.Username = username
	return &this
}

// NewCreateAccessTicketRequestWithDefaults instantiates a new CreateAccessTicketRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessTicketRequestWithDefaults() *CreateAccessTicketRequest {
	this := CreateAccessTicketRequest{}
	return &this
}

// GetNewFormat returns the NewFormat field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetNewFormat() bool {
	if o == nil || IsNil(o.NewFormat) {
		var ret bool
		return ret
	}
	return *o.NewFormat
}

// GetNewFormatOk returns a tuple with the NewFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetNewFormatOk() (*bool, bool) {
	if o == nil || IsNil(o.NewFormat) {
		return nil, false
	}
	return o.NewFormat, true
}

// HasNewFormat returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasNewFormat() bool {
	if o != nil && !IsNil(o.NewFormat) {
		return true
	}

	return false
}

// SetNewFormat gets a reference to the given bool and assigns it to the NewFormat field.
func (o *CreateAccessTicketRequest) SetNewFormat(v bool) {
	o.NewFormat = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetOtp() string {
	if o == nil || IsNil(o.Otp) {
		var ret string
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetOtpOk() (*string, bool) {
	if o == nil || IsNil(o.Otp) {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasOtp() bool {
	if o != nil && !IsNil(o.Otp) {
		return true
	}

	return false
}

// SetOtp gets a reference to the given string and assigns it to the Otp field.
func (o *CreateAccessTicketRequest) SetOtp(v string) {
	o.Otp = &v
}

// GetPassword returns the Password field value
func (o *CreateAccessTicketRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *CreateAccessTicketRequest) SetPassword(v string) {
	o.Password = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *CreateAccessTicketRequest) SetPath(v string) {
	o.Path = &v
}

// GetPrivs returns the Privs field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetPrivs() string {
	if o == nil || IsNil(o.Privs) {
		var ret string
		return ret
	}
	return *o.Privs
}

// GetPrivsOk returns a tuple with the Privs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetPrivsOk() (*string, bool) {
	if o == nil || IsNil(o.Privs) {
		return nil, false
	}
	return o.Privs, true
}

// HasPrivs returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasPrivs() bool {
	if o != nil && !IsNil(o.Privs) {
		return true
	}

	return false
}

// SetPrivs gets a reference to the given string and assigns it to the Privs field.
func (o *CreateAccessTicketRequest) SetPrivs(v string) {
	o.Privs = &v
}

// GetRealm returns the Realm field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetRealm() string {
	if o == nil || IsNil(o.Realm) {
		var ret string
		return ret
	}
	return *o.Realm
}

// GetRealmOk returns a tuple with the Realm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetRealmOk() (*string, bool) {
	if o == nil || IsNil(o.Realm) {
		return nil, false
	}
	return o.Realm, true
}

// HasRealm returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasRealm() bool {
	if o != nil && !IsNil(o.Realm) {
		return true
	}

	return false
}

// SetRealm gets a reference to the given string and assigns it to the Realm field.
func (o *CreateAccessTicketRequest) SetRealm(v string) {
	o.Realm = &v
}

// GetTfaChallenge returns the TfaChallenge field value if set, zero value otherwise.
func (o *CreateAccessTicketRequest) GetTfaChallenge() string {
	if o == nil || IsNil(o.TfaChallenge) {
		var ret string
		return ret
	}
	return *o.TfaChallenge
}

// GetTfaChallengeOk returns a tuple with the TfaChallenge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetTfaChallengeOk() (*string, bool) {
	if o == nil || IsNil(o.TfaChallenge) {
		return nil, false
	}
	return o.TfaChallenge, true
}

// HasTfaChallenge returns a boolean if a field has been set.
func (o *CreateAccessTicketRequest) HasTfaChallenge() bool {
	if o != nil && !IsNil(o.TfaChallenge) {
		return true
	}

	return false
}

// SetTfaChallenge gets a reference to the given string and assigns it to the TfaChallenge field.
func (o *CreateAccessTicketRequest) SetTfaChallenge(v string) {
	o.TfaChallenge = &v
}

// GetUsername returns the Username field value
func (o *CreateAccessTicketRequest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *CreateAccessTicketRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *CreateAccessTicketRequest) SetUsername(v string) {
	o.Username = v
}

func (o CreateAccessTicketRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAccessTicketRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewFormat) {
		toSerialize["new-format"] = o.NewFormat
	}
	if !IsNil(o.Otp) {
		toSerialize["otp"] = o.Otp
	}
	toSerialize["password"] = o.Password
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Privs) {
		toSerialize["privs"] = o.Privs
	}
	if !IsNil(o.Realm) {
		toSerialize["realm"] = o.Realm
	}
	if !IsNil(o.TfaChallenge) {
		toSerialize["tfa-challenge"] = o.TfaChallenge
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableCreateAccessTicketRequest struct {
	value *CreateAccessTicketRequest
	isSet bool
}

func (v NullableCreateAccessTicketRequest) Get() *CreateAccessTicketRequest {
	return v.value
}

func (v *NullableCreateAccessTicketRequest) Set(val *CreateAccessTicketRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessTicketRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessTicketRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessTicketRequest(val *CreateAccessTicketRequest) *NullableCreateAccessTicketRequest {
	return &NullableCreateAccessTicketRequest{value: val, isSet: true}
}

func (v NullableCreateAccessTicketRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessTicketRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


