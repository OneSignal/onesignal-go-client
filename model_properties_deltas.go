/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.2.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// PropertiesDeltas struct for PropertiesDeltas
type PropertiesDeltas struct {
	SessionTime *float32 `json:"session_time,omitempty"`
	SessionCount *float32 `json:"session_count,omitempty"`
	Purchases []Purchase `json:"purchases,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PropertiesDeltas PropertiesDeltas

// NewPropertiesDeltas instantiates a new PropertiesDeltas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertiesDeltas() *PropertiesDeltas {
	this := PropertiesDeltas{}
	return &this
}

// NewPropertiesDeltasWithDefaults instantiates a new PropertiesDeltas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertiesDeltasWithDefaults() *PropertiesDeltas {
	this := PropertiesDeltas{}
	return &this
}

// GetSessionTime returns the SessionTime field value if set, zero value otherwise.
func (o *PropertiesDeltas) GetSessionTime() float32 {
	if o == nil || o.SessionTime == nil {
		var ret float32
		return ret
	}
	return *o.SessionTime
}

// GetSessionTimeOk returns a tuple with the SessionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertiesDeltas) GetSessionTimeOk() (*float32, bool) {
	if o == nil || o.SessionTime == nil {
		return nil, false
	}
	return o.SessionTime, true
}

// HasSessionTime returns a boolean if a field has been set.
func (o *PropertiesDeltas) HasSessionTime() bool {
	if o != nil && o.SessionTime != nil {
		return true
	}

	return false
}

// SetSessionTime gets a reference to the given float32 and assigns it to the SessionTime field.
func (o *PropertiesDeltas) SetSessionTime(v float32) {
	o.SessionTime = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *PropertiesDeltas) GetSessionCount() float32 {
	if o == nil || o.SessionCount == nil {
		var ret float32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertiesDeltas) GetSessionCountOk() (*float32, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *PropertiesDeltas) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given float32 and assigns it to the SessionCount field.
func (o *PropertiesDeltas) SetSessionCount(v float32) {
	o.SessionCount = &v
}

// GetPurchases returns the Purchases field value if set, zero value otherwise.
func (o *PropertiesDeltas) GetPurchases() []Purchase {
	if o == nil || o.Purchases == nil {
		var ret []Purchase
		return ret
	}
	return o.Purchases
}

// GetPurchasesOk returns a tuple with the Purchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertiesDeltas) GetPurchasesOk() ([]Purchase, bool) {
	if o == nil || o.Purchases == nil {
		return nil, false
	}
	return o.Purchases, true
}

// HasPurchases returns a boolean if a field has been set.
func (o *PropertiesDeltas) HasPurchases() bool {
	if o != nil && o.Purchases != nil {
		return true
	}

	return false
}

// SetPurchases gets a reference to the given []Purchase and assigns it to the Purchases field.
func (o *PropertiesDeltas) SetPurchases(v []Purchase) {
	o.Purchases = v
}

func (o PropertiesDeltas) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SessionTime != nil {
		toSerialize["session_time"] = o.SessionTime
	}
	if o.SessionCount != nil {
		toSerialize["session_count"] = o.SessionCount
	}
	if o.Purchases != nil {
		toSerialize["purchases"] = o.Purchases
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PropertiesDeltas) UnmarshalJSON(bytes []byte) (err error) {
	varPropertiesDeltas := _PropertiesDeltas{}

	if err = json.Unmarshal(bytes, &varPropertiesDeltas); err == nil {
		*o = PropertiesDeltas(varPropertiesDeltas)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "session_time")
		delete(additionalProperties, "session_count")
		delete(additionalProperties, "purchases")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePropertiesDeltas struct {
	value *PropertiesDeltas
	isSet bool
}

func (v NullablePropertiesDeltas) Get() *PropertiesDeltas {
	return v.value
}

func (v *NullablePropertiesDeltas) Set(val *PropertiesDeltas) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertiesDeltas) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertiesDeltas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertiesDeltas(val *PropertiesDeltas) *NullablePropertiesDeltas {
	return &NullablePropertiesDeltas{value: val, isSet: true}
}

func (v NullablePropertiesDeltas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertiesDeltas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


