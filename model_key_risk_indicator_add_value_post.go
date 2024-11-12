/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the KeyRiskIndicatorAddValuePost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyRiskIndicatorAddValuePost{}

// KeyRiskIndicatorAddValuePost struct for KeyRiskIndicatorAddValuePost
type KeyRiskIndicatorAddValuePost struct {
	Value float32 `json:"value"`
	HistoricalDate string `json:"historical_date"`
	Comments *string `json:"comments,omitempty"`
}

type _KeyRiskIndicatorAddValuePost KeyRiskIndicatorAddValuePost

// NewKeyRiskIndicatorAddValuePost instantiates a new KeyRiskIndicatorAddValuePost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyRiskIndicatorAddValuePost(value float32, historicalDate string) *KeyRiskIndicatorAddValuePost {
	this := KeyRiskIndicatorAddValuePost{}
	this.Value = value
	this.HistoricalDate = historicalDate
	return &this
}

// NewKeyRiskIndicatorAddValuePostWithDefaults instantiates a new KeyRiskIndicatorAddValuePost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyRiskIndicatorAddValuePostWithDefaults() *KeyRiskIndicatorAddValuePost {
	this := KeyRiskIndicatorAddValuePost{}
	return &this
}

// GetValue returns the Value field value
func (o *KeyRiskIndicatorAddValuePost) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorAddValuePost) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *KeyRiskIndicatorAddValuePost) SetValue(v float32) {
	o.Value = v
}

// GetHistoricalDate returns the HistoricalDate field value
func (o *KeyRiskIndicatorAddValuePost) GetHistoricalDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HistoricalDate
}

// GetHistoricalDateOk returns a tuple with the HistoricalDate field value
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorAddValuePost) GetHistoricalDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HistoricalDate, true
}

// SetHistoricalDate sets field value
func (o *KeyRiskIndicatorAddValuePost) SetHistoricalDate(v string) {
	o.HistoricalDate = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *KeyRiskIndicatorAddValuePost) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyRiskIndicatorAddValuePost) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *KeyRiskIndicatorAddValuePost) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *KeyRiskIndicatorAddValuePost) SetComments(v string) {
	o.Comments = &v
}

func (o KeyRiskIndicatorAddValuePost) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyRiskIndicatorAddValuePost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	toSerialize["historical_date"] = o.HistoricalDate
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	return toSerialize, nil
}

func (o *KeyRiskIndicatorAddValuePost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
		"historical_date",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varKeyRiskIndicatorAddValuePost := _KeyRiskIndicatorAddValuePost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varKeyRiskIndicatorAddValuePost)

	if err != nil {
		return err
	}

	*o = KeyRiskIndicatorAddValuePost(varKeyRiskIndicatorAddValuePost)

	return err
}

type NullableKeyRiskIndicatorAddValuePost struct {
	value *KeyRiskIndicatorAddValuePost
	isSet bool
}

func (v NullableKeyRiskIndicatorAddValuePost) Get() *KeyRiskIndicatorAddValuePost {
	return v.value
}

func (v *NullableKeyRiskIndicatorAddValuePost) Set(val *KeyRiskIndicatorAddValuePost) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyRiskIndicatorAddValuePost) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyRiskIndicatorAddValuePost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyRiskIndicatorAddValuePost(val *KeyRiskIndicatorAddValuePost) *NullableKeyRiskIndicatorAddValuePost {
	return &NullableKeyRiskIndicatorAddValuePost{value: val, isSet: true}
}

func (v NullableKeyRiskIndicatorAddValuePost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyRiskIndicatorAddValuePost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

