/*
NuoDB Control Plane REST API

NuoDB Control Plane (CP) allows users to create and manage NuoDB databases remotely using a Database as a Service (DBaaS) model.

API version: 2.3.0
Contact: NuoDB.Support@3ds.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nuodbaas

import (
	"encoding/json"
)

// checks if the Scope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scope{}

// Scope struct for Scope
type Scope struct {
	FieldSelector *FieldSelector `json:"fieldSelector,omitempty"`
	// The label keys on which the selected databases are divided into groups.
	GroupByLabels []string `json:"groupByLabels,omitempty"`
	LabelSelector *LabelSelector `json:"labelSelector,omitempty"`
}

// NewScope instantiates a new Scope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScope() *Scope {
	this := Scope{}
	return &this
}

// NewScopeWithDefaults instantiates a new Scope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopeWithDefaults() *Scope {
	this := Scope{}
	return &this
}

// GetFieldSelector returns the FieldSelector field value if set, zero value otherwise.
func (o *Scope) GetFieldSelector() FieldSelector {
	if o == nil || IsNil(o.FieldSelector) {
		var ret FieldSelector
		return ret
	}
	return *o.FieldSelector
}

// GetFieldSelectorOk returns a tuple with the FieldSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetFieldSelectorOk() (*FieldSelector, bool) {
	if o == nil || IsNil(o.FieldSelector) {
		return nil, false
	}
	return o.FieldSelector, true
}

// HasFieldSelector returns a boolean if a field has been set.
func (o *Scope) HasFieldSelector() bool {
	if o != nil && !IsNil(o.FieldSelector) {
		return true
	}

	return false
}

// SetFieldSelector gets a reference to the given FieldSelector and assigns it to the FieldSelector field.
func (o *Scope) SetFieldSelector(v FieldSelector) {
	o.FieldSelector = &v
}

// GetGroupByLabels returns the GroupByLabels field value if set, zero value otherwise.
func (o *Scope) GetGroupByLabels() []string {
	if o == nil || IsNil(o.GroupByLabels) {
		var ret []string
		return ret
	}
	return o.GroupByLabels
}

// GetGroupByLabelsOk returns a tuple with the GroupByLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetGroupByLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupByLabels) {
		return nil, false
	}
	return o.GroupByLabels, true
}

// HasGroupByLabels returns a boolean if a field has been set.
func (o *Scope) HasGroupByLabels() bool {
	if o != nil && !IsNil(o.GroupByLabels) {
		return true
	}

	return false
}

// SetGroupByLabels gets a reference to the given []string and assigns it to the GroupByLabels field.
func (o *Scope) SetGroupByLabels(v []string) {
	o.GroupByLabels = v
}

// GetLabelSelector returns the LabelSelector field value if set, zero value otherwise.
func (o *Scope) GetLabelSelector() LabelSelector {
	if o == nil || IsNil(o.LabelSelector) {
		var ret LabelSelector
		return ret
	}
	return *o.LabelSelector
}

// GetLabelSelectorOk returns a tuple with the LabelSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scope) GetLabelSelectorOk() (*LabelSelector, bool) {
	if o == nil || IsNil(o.LabelSelector) {
		return nil, false
	}
	return o.LabelSelector, true
}

// HasLabelSelector returns a boolean if a field has been set.
func (o *Scope) HasLabelSelector() bool {
	if o != nil && !IsNil(o.LabelSelector) {
		return true
	}

	return false
}

// SetLabelSelector gets a reference to the given LabelSelector and assigns it to the LabelSelector field.
func (o *Scope) SetLabelSelector(v LabelSelector) {
	o.LabelSelector = &v
}

func (o Scope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldSelector) {
		toSerialize["fieldSelector"] = o.FieldSelector
	}
	if !IsNil(o.GroupByLabels) {
		toSerialize["groupByLabels"] = o.GroupByLabels
	}
	if !IsNil(o.LabelSelector) {
		toSerialize["labelSelector"] = o.LabelSelector
	}
	return toSerialize, nil
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


