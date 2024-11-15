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

// checks if the ReportsGeneratePostRequestQueryDataQueryFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportsGeneratePostRequestQueryDataQueryFields{}

// ReportsGeneratePostRequestQueryDataQueryFields struct for ReportsGeneratePostRequestQueryDataQueryFields
type ReportsGeneratePostRequestQueryDataQueryFields struct {
	Control []ReportsGeneratePostRequestQueryDataQueryFieldsControlInner `json:"control,omitempty"`
	ControlTest []ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner `json:"controlTest,omitempty"`
	ControlIssue []interface{} `json:"controlIssue,omitempty"`
	Test1 []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner `json:"test_1,omitempty"`
	Test2 []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner `json:"test_2,omitempty"`
	Test3 []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner `json:"test_3,omitempty"`
}

// NewReportsGeneratePostRequestQueryDataQueryFields instantiates a new ReportsGeneratePostRequestQueryDataQueryFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportsGeneratePostRequestQueryDataQueryFields() *ReportsGeneratePostRequestQueryDataQueryFields {
	this := ReportsGeneratePostRequestQueryDataQueryFields{}
	return &this
}

// NewReportsGeneratePostRequestQueryDataQueryFieldsWithDefaults instantiates a new ReportsGeneratePostRequestQueryDataQueryFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportsGeneratePostRequestQueryDataQueryFieldsWithDefaults() *ReportsGeneratePostRequestQueryDataQueryFields {
	this := ReportsGeneratePostRequestQueryDataQueryFields{}
	return &this
}

// GetControl returns the Control field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControl() []ReportsGeneratePostRequestQueryDataQueryFieldsControlInner {
	if o == nil || IsNil(o.Control) {
		var ret []ReportsGeneratePostRequestQueryDataQueryFieldsControlInner
		return ret
	}
	return o.Control
}

// GetControlOk returns a tuple with the Control field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControlOk() ([]ReportsGeneratePostRequestQueryDataQueryFieldsControlInner, bool) {
	if o == nil || IsNil(o.Control) {
		return nil, false
	}
	return o.Control, true
}

// HasControl returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasControl() bool {
	if o != nil && !IsNil(o.Control) {
		return true
	}

	return false
}

// SetControl gets a reference to the given []ReportsGeneratePostRequestQueryDataQueryFieldsControlInner and assigns it to the Control field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetControl(v []ReportsGeneratePostRequestQueryDataQueryFieldsControlInner) {
	o.Control = v
}

// GetControlTest returns the ControlTest field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControlTest() []ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner {
	if o == nil || IsNil(o.ControlTest) {
		var ret []ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner
		return ret
	}
	return o.ControlTest
}

// GetControlTestOk returns a tuple with the ControlTest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControlTestOk() ([]ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner, bool) {
	if o == nil || IsNil(o.ControlTest) {
		return nil, false
	}
	return o.ControlTest, true
}

// HasControlTest returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasControlTest() bool {
	if o != nil && !IsNil(o.ControlTest) {
		return true
	}

	return false
}

// SetControlTest gets a reference to the given []ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner and assigns it to the ControlTest field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetControlTest(v []ReportsGeneratePostRequestQueryDataQueryFieldsControlTestInner) {
	o.ControlTest = v
}

// GetControlIssue returns the ControlIssue field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControlIssue() []interface{} {
	if o == nil || IsNil(o.ControlIssue) {
		var ret []interface{}
		return ret
	}
	return o.ControlIssue
}

// GetControlIssueOk returns a tuple with the ControlIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetControlIssueOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ControlIssue) {
		return nil, false
	}
	return o.ControlIssue, true
}

// HasControlIssue returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasControlIssue() bool {
	if o != nil && !IsNil(o.ControlIssue) {
		return true
	}

	return false
}

// SetControlIssue gets a reference to the given []interface{} and assigns it to the ControlIssue field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetControlIssue(v []interface{}) {
	o.ControlIssue = v
}

// GetTest1 returns the Test1 field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest1() []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner {
	if o == nil || IsNil(o.Test1) {
		var ret []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner
		return ret
	}
	return o.Test1
}

// GetTest1Ok returns a tuple with the Test1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest1Ok() ([]ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner, bool) {
	if o == nil || IsNil(o.Test1) {
		return nil, false
	}
	return o.Test1, true
}

// HasTest1 returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasTest1() bool {
	if o != nil && !IsNil(o.Test1) {
		return true
	}

	return false
}

// SetTest1 gets a reference to the given []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner and assigns it to the Test1 field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetTest1(v []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner) {
	o.Test1 = v
}

// GetTest2 returns the Test2 field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest2() []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner {
	if o == nil || IsNil(o.Test2) {
		var ret []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner
		return ret
	}
	return o.Test2
}

// GetTest2Ok returns a tuple with the Test2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest2Ok() ([]ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner, bool) {
	if o == nil || IsNil(o.Test2) {
		return nil, false
	}
	return o.Test2, true
}

// HasTest2 returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasTest2() bool {
	if o != nil && !IsNil(o.Test2) {
		return true
	}

	return false
}

// SetTest2 gets a reference to the given []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner and assigns it to the Test2 field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetTest2(v []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner) {
	o.Test2 = v
}

// GetTest3 returns the Test3 field value if set, zero value otherwise.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest3() []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner {
	if o == nil || IsNil(o.Test3) {
		var ret []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner
		return ret
	}
	return o.Test3
}

// GetTest3Ok returns a tuple with the Test3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) GetTest3Ok() ([]ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner, bool) {
	if o == nil || IsNil(o.Test3) {
		return nil, false
	}
	return o.Test3, true
}

// HasTest3 returns a boolean if a field has been set.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) HasTest3() bool {
	if o != nil && !IsNil(o.Test3) {
		return true
	}

	return false
}

// SetTest3 gets a reference to the given []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner and assigns it to the Test3 field.
func (o *ReportsGeneratePostRequestQueryDataQueryFields) SetTest3(v []ReportsGeneratePostRequestQueryDataQueryFieldsTest1Inner) {
	o.Test3 = v
}

func (o ReportsGeneratePostRequestQueryDataQueryFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportsGeneratePostRequestQueryDataQueryFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Control) {
		toSerialize["control"] = o.Control
	}
	if !IsNil(o.ControlTest) {
		toSerialize["controlTest"] = o.ControlTest
	}
	if !IsNil(o.ControlIssue) {
		toSerialize["controlIssue"] = o.ControlIssue
	}
	if !IsNil(o.Test1) {
		toSerialize["test_1"] = o.Test1
	}
	if !IsNil(o.Test2) {
		toSerialize["test_2"] = o.Test2
	}
	if !IsNil(o.Test3) {
		toSerialize["test_3"] = o.Test3
	}
	return toSerialize, nil
}

type NullableReportsGeneratePostRequestQueryDataQueryFields struct {
	value *ReportsGeneratePostRequestQueryDataQueryFields
	isSet bool
}

func (v NullableReportsGeneratePostRequestQueryDataQueryFields) Get() *ReportsGeneratePostRequestQueryDataQueryFields {
	return v.value
}

func (v *NullableReportsGeneratePostRequestQueryDataQueryFields) Set(val *ReportsGeneratePostRequestQueryDataQueryFields) {
	v.value = val
	v.isSet = true
}

func (v NullableReportsGeneratePostRequestQueryDataQueryFields) IsSet() bool {
	return v.isSet
}

func (v *NullableReportsGeneratePostRequestQueryDataQueryFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportsGeneratePostRequestQueryDataQueryFields(val *ReportsGeneratePostRequestQueryDataQueryFields) *NullableReportsGeneratePostRequestQueryDataQueryFields {
	return &NullableReportsGeneratePostRequestQueryDataQueryFields{value: val, isSet: true}
}

func (v NullableReportsGeneratePostRequestQueryDataQueryFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportsGeneratePostRequestQueryDataQueryFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


