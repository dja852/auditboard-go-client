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

// checks if the ControlsDataRequestZipPostRequestOptionsPdfSections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsDataRequestZipPostRequestOptionsPdfSections{}

// ControlsDataRequestZipPostRequestOptionsPdfSections struct for ControlsDataRequestZipPostRequestOptionsPdfSections
type ControlsDataRequestZipPostRequestOptionsPdfSections struct {
	PdfControlTestInformation *bool `json:"pdfControlTestInformation,omitempty"`
	PdfControlIssues *bool `json:"pdfControlIssues,omitempty"`
	PdfControlInformation *bool `json:"pdfControlInformation,omitempty"`
}

// NewControlsDataRequestZipPostRequestOptionsPdfSections instantiates a new ControlsDataRequestZipPostRequestOptionsPdfSections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsDataRequestZipPostRequestOptionsPdfSections() *ControlsDataRequestZipPostRequestOptionsPdfSections {
	this := ControlsDataRequestZipPostRequestOptionsPdfSections{}
	return &this
}

// NewControlsDataRequestZipPostRequestOptionsPdfSectionsWithDefaults instantiates a new ControlsDataRequestZipPostRequestOptionsPdfSections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsDataRequestZipPostRequestOptionsPdfSectionsWithDefaults() *ControlsDataRequestZipPostRequestOptionsPdfSections {
	this := ControlsDataRequestZipPostRequestOptionsPdfSections{}
	return &this
}

// GetPdfControlTestInformation returns the PdfControlTestInformation field value if set, zero value otherwise.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlTestInformation() bool {
	if o == nil || IsNil(o.PdfControlTestInformation) {
		var ret bool
		return ret
	}
	return *o.PdfControlTestInformation
}

// GetPdfControlTestInformationOk returns a tuple with the PdfControlTestInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlTestInformationOk() (*bool, bool) {
	if o == nil || IsNil(o.PdfControlTestInformation) {
		return nil, false
	}
	return o.PdfControlTestInformation, true
}

// HasPdfControlTestInformation returns a boolean if a field has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) HasPdfControlTestInformation() bool {
	if o != nil && !IsNil(o.PdfControlTestInformation) {
		return true
	}

	return false
}

// SetPdfControlTestInformation gets a reference to the given bool and assigns it to the PdfControlTestInformation field.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) SetPdfControlTestInformation(v bool) {
	o.PdfControlTestInformation = &v
}

// GetPdfControlIssues returns the PdfControlIssues field value if set, zero value otherwise.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlIssues() bool {
	if o == nil || IsNil(o.PdfControlIssues) {
		var ret bool
		return ret
	}
	return *o.PdfControlIssues
}

// GetPdfControlIssuesOk returns a tuple with the PdfControlIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlIssuesOk() (*bool, bool) {
	if o == nil || IsNil(o.PdfControlIssues) {
		return nil, false
	}
	return o.PdfControlIssues, true
}

// HasPdfControlIssues returns a boolean if a field has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) HasPdfControlIssues() bool {
	if o != nil && !IsNil(o.PdfControlIssues) {
		return true
	}

	return false
}

// SetPdfControlIssues gets a reference to the given bool and assigns it to the PdfControlIssues field.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) SetPdfControlIssues(v bool) {
	o.PdfControlIssues = &v
}

// GetPdfControlInformation returns the PdfControlInformation field value if set, zero value otherwise.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlInformation() bool {
	if o == nil || IsNil(o.PdfControlInformation) {
		var ret bool
		return ret
	}
	return *o.PdfControlInformation
}

// GetPdfControlInformationOk returns a tuple with the PdfControlInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) GetPdfControlInformationOk() (*bool, bool) {
	if o == nil || IsNil(o.PdfControlInformation) {
		return nil, false
	}
	return o.PdfControlInformation, true
}

// HasPdfControlInformation returns a boolean if a field has been set.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) HasPdfControlInformation() bool {
	if o != nil && !IsNil(o.PdfControlInformation) {
		return true
	}

	return false
}

// SetPdfControlInformation gets a reference to the given bool and assigns it to the PdfControlInformation field.
func (o *ControlsDataRequestZipPostRequestOptionsPdfSections) SetPdfControlInformation(v bool) {
	o.PdfControlInformation = &v
}

func (o ControlsDataRequestZipPostRequestOptionsPdfSections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsDataRequestZipPostRequestOptionsPdfSections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PdfControlTestInformation) {
		toSerialize["pdfControlTestInformation"] = o.PdfControlTestInformation
	}
	if !IsNil(o.PdfControlIssues) {
		toSerialize["pdfControlIssues"] = o.PdfControlIssues
	}
	if !IsNil(o.PdfControlInformation) {
		toSerialize["pdfControlInformation"] = o.PdfControlInformation
	}
	return toSerialize, nil
}

type NullableControlsDataRequestZipPostRequestOptionsPdfSections struct {
	value *ControlsDataRequestZipPostRequestOptionsPdfSections
	isSet bool
}

func (v NullableControlsDataRequestZipPostRequestOptionsPdfSections) Get() *ControlsDataRequestZipPostRequestOptionsPdfSections {
	return v.value
}

func (v *NullableControlsDataRequestZipPostRequestOptionsPdfSections) Set(val *ControlsDataRequestZipPostRequestOptionsPdfSections) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsDataRequestZipPostRequestOptionsPdfSections) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsDataRequestZipPostRequestOptionsPdfSections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsDataRequestZipPostRequestOptionsPdfSections(val *ControlsDataRequestZipPostRequestOptionsPdfSections) *NullableControlsDataRequestZipPostRequestOptionsPdfSections {
	return &NullableControlsDataRequestZipPostRequestOptionsPdfSections{value: val, isSet: true}
}

func (v NullableControlsDataRequestZipPostRequestOptionsPdfSections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsDataRequestZipPostRequestOptionsPdfSections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

