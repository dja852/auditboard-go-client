/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
)

// checks if the KeyPerformanceIndicatorAddValuePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyPerformanceIndicatorAddValuePost{}

// KeyPerformanceIndicatorAddValuePost struct for KeyPerformanceIndicatorAddValuePost
type KeyPerformanceIndicatorAddValuePost struct {
	Value float32 `json:"value"`
	Target float32 `json:"target"`
	HistoricalDate string `json:"historical_date"`
	Comments *string `json:"comments,omitempty"`
}

type _KeyPerformanceIndicatorAddValuePost KeyPerformanceIndicatorAddValuePost

// NewKeyPerformanceIndicatorAddValuePost instantiates a new KeyPerformanceIndicatorAddValuePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyPerformanceIndicatorAddValuePost(value float32, target float32, historicalDate string) *KeyPerformanceIndicatorAddValuePost {
	this := KeyPerformanceIndicatorAddValuePost{}
	this.Value = value
	this.Target = target
	this.HistoricalDate = historicalDate
	return &this
}

// NewKeyPerformanceIndicatorAddValuePostWithDefaults instantiates a new KeyPerformanceIndicatorAddValuePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyPerformanceIndicatorAddValuePostWithDefaults() *KeyPerformanceIndicatorAddValuePost {
	this := KeyPerformanceIndicatorAddValuePost{}
	return &this
}

// GetValue returns the Value field value
func (o *KeyPerformanceIndicatorAddValuePost) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KeyPerformanceIndicatorAddValuePost) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KeyPerformanceIndicatorAddValuePost) SetValue(v float32) {
	o.Value = v
}

// GetTarget returns the Target field value
func (o *KeyPerformanceIndicatorAddValuePost) GetTarget() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *KeyPerformanceIndicatorAddValuePost) GetTargetOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *KeyPerformanceIndicatorAddValuePost) SetTarget(v float32) {
	o.Target = v
}

// GetHistoricalDate returns the HistoricalDate field value
func (o *KeyPerformanceIndicatorAddValuePost) GetHistoricalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HistoricalDate
}

// GetHistoricalDateOk returns a tuple with the HistoricalDate field value
// and a boolean to check if the value has been set.
func (o *KeyPerformanceIndicatorAddValuePost) GetHistoricalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoricalDate, true
}

// SetHistoricalDate sets field value
func (o *KeyPerformanceIndicatorAddValuePost) SetHistoricalDate(v string) {
	o.HistoricalDate = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *KeyPerformanceIndicatorAddValuePost) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyPerformanceIndicatorAddValuePost) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *KeyPerformanceIndicatorAddValuePost) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *KeyPerformanceIndicatorAddValuePost) SetComments(v string) {
	o.Comments = &v
}

func (o KeyPerformanceIndicatorAddValuePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyPerformanceIndicatorAddValuePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["target"] = o.Target
	toSerialize["historical_date"] = o.HistoricalDate
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	return toSerialize, nil
}

type NullableKeyPerformanceIndicatorAddValuePost struct {
	value *KeyPerformanceIndicatorAddValuePost
	isSet bool
}

func (v NullableKeyPerformanceIndicatorAddValuePost) Get() *KeyPerformanceIndicatorAddValuePost {
	return v.value
}

func (v *NullableKeyPerformanceIndicatorAddValuePost) Set(val *KeyPerformanceIndicatorAddValuePost) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyPerformanceIndicatorAddValuePost) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyPerformanceIndicatorAddValuePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyPerformanceIndicatorAddValuePost(val *KeyPerformanceIndicatorAddValuePost) *NullableKeyPerformanceIndicatorAddValuePost {
	return &NullableKeyPerformanceIndicatorAddValuePost{value: val, isSet: true}
}

func (v NullableKeyPerformanceIndicatorAddValuePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyPerformanceIndicatorAddValuePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


