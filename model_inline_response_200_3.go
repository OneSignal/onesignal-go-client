/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.3.0
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	InAppMessages []map[string]interface{} `json:"in_app_messages,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InlineResponse2003 InlineResponse2003

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetInAppMessages returns the InAppMessages field value if set, zero value otherwise.
func (o *InlineResponse2003) GetInAppMessages() []map[string]interface{} {
	if o == nil || o.InAppMessages == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.InAppMessages
}

// GetInAppMessagesOk returns a tuple with the InAppMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetInAppMessagesOk() ([]map[string]interface{}, bool) {
	if o == nil || o.InAppMessages == nil {
		return nil, false
	}
	return o.InAppMessages, true
}

// HasInAppMessages returns a boolean if a field has been set.
func (o *InlineResponse2003) HasInAppMessages() bool {
	if o != nil && o.InAppMessages != nil {
		return true
	}

	return false
}

// SetInAppMessages gets a reference to the given []map[string]interface{} and assigns it to the InAppMessages field.
func (o *InlineResponse2003) SetInAppMessages(v []map[string]interface{}) {
	o.InAppMessages = v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InAppMessages != nil {
		toSerialize["in_app_messages"] = o.InAppMessages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InlineResponse2003) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2003 := _InlineResponse2003{}

	if err = json.Unmarshal(bytes, &varInlineResponse2003); err == nil {
		*o = InlineResponse2003(varInlineResponse2003)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "in_app_messages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


