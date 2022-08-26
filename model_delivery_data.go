/*
OneSignal

A powerful way to send personalized messages at scale and build effective customer engagement strategies. Learn more at onesignal.com

API version: 1.0.1
Contact: devrel@onesignal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onesignal

import (
	"encoding/json"
)

// DeliveryData struct for DeliveryData
type DeliveryData struct {
	// Number of messages delivered to push servers, mobile carriers, or email service providers.
	Successful NullableInt32 `json:"successful,omitempty"`
	// Number of messages sent to unsubscribed devices.
	Failed NullableInt32 `json:"failed,omitempty"`
	// Number of errors reported.
	Errored NullableInt32 `json:"errored,omitempty"`
	// Number of messages that were clicked.
	Converted NullableInt32 `json:"converted,omitempty"`
	// Number of devices that received the message.
	Received NullableInt32 `json:"received,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeliveryData DeliveryData

// NewDeliveryData instantiates a new DeliveryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryData() *DeliveryData {
	this := DeliveryData{}
	return &this
}

// NewDeliveryDataWithDefaults instantiates a new DeliveryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryDataWithDefaults() *DeliveryData {
	this := DeliveryData{}
	return &this
}

// GetSuccessful returns the Successful field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryData) GetSuccessful() int32 {
	if o == nil || o.Successful.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Successful.Get()
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryData) GetSuccessfulOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Successful.Get(), o.Successful.IsSet()
}

// HasSuccessful returns a boolean if a field has been set.
func (o *DeliveryData) HasSuccessful() bool {
	if o != nil && o.Successful.IsSet() {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given NullableInt32 and assigns it to the Successful field.
func (o *DeliveryData) SetSuccessful(v int32) {
	o.Successful.Set(&v)
}
// SetSuccessfulNil sets the value for Successful to be an explicit nil
func (o *DeliveryData) SetSuccessfulNil() {
	o.Successful.Set(nil)
}

// UnsetSuccessful ensures that no value is present for Successful, not even an explicit nil
func (o *DeliveryData) UnsetSuccessful() {
	o.Successful.Unset()
}

// GetFailed returns the Failed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryData) GetFailed() int32 {
	if o == nil || o.Failed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Failed.Get()
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryData) GetFailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Failed.Get(), o.Failed.IsSet()
}

// HasFailed returns a boolean if a field has been set.
func (o *DeliveryData) HasFailed() bool {
	if o != nil && o.Failed.IsSet() {
		return true
	}

	return false
}

// SetFailed gets a reference to the given NullableInt32 and assigns it to the Failed field.
func (o *DeliveryData) SetFailed(v int32) {
	o.Failed.Set(&v)
}
// SetFailedNil sets the value for Failed to be an explicit nil
func (o *DeliveryData) SetFailedNil() {
	o.Failed.Set(nil)
}

// UnsetFailed ensures that no value is present for Failed, not even an explicit nil
func (o *DeliveryData) UnsetFailed() {
	o.Failed.Unset()
}

// GetErrored returns the Errored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryData) GetErrored() int32 {
	if o == nil || o.Errored.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Errored.Get()
}

// GetErroredOk returns a tuple with the Errored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryData) GetErroredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errored.Get(), o.Errored.IsSet()
}

// HasErrored returns a boolean if a field has been set.
func (o *DeliveryData) HasErrored() bool {
	if o != nil && o.Errored.IsSet() {
		return true
	}

	return false
}

// SetErrored gets a reference to the given NullableInt32 and assigns it to the Errored field.
func (o *DeliveryData) SetErrored(v int32) {
	o.Errored.Set(&v)
}
// SetErroredNil sets the value for Errored to be an explicit nil
func (o *DeliveryData) SetErroredNil() {
	o.Errored.Set(nil)
}

// UnsetErrored ensures that no value is present for Errored, not even an explicit nil
func (o *DeliveryData) UnsetErrored() {
	o.Errored.Unset()
}

// GetConverted returns the Converted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryData) GetConverted() int32 {
	if o == nil || o.Converted.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Converted.Get()
}

// GetConvertedOk returns a tuple with the Converted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryData) GetConvertedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Converted.Get(), o.Converted.IsSet()
}

// HasConverted returns a boolean if a field has been set.
func (o *DeliveryData) HasConverted() bool {
	if o != nil && o.Converted.IsSet() {
		return true
	}

	return false
}

// SetConverted gets a reference to the given NullableInt32 and assigns it to the Converted field.
func (o *DeliveryData) SetConverted(v int32) {
	o.Converted.Set(&v)
}
// SetConvertedNil sets the value for Converted to be an explicit nil
func (o *DeliveryData) SetConvertedNil() {
	o.Converted.Set(nil)
}

// UnsetConverted ensures that no value is present for Converted, not even an explicit nil
func (o *DeliveryData) UnsetConverted() {
	o.Converted.Unset()
}

// GetReceived returns the Received field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DeliveryData) GetReceived() int32 {
	if o == nil || o.Received.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Received.Get()
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DeliveryData) GetReceivedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Received.Get(), o.Received.IsSet()
}

// HasReceived returns a boolean if a field has been set.
func (o *DeliveryData) HasReceived() bool {
	if o != nil && o.Received.IsSet() {
		return true
	}

	return false
}

// SetReceived gets a reference to the given NullableInt32 and assigns it to the Received field.
func (o *DeliveryData) SetReceived(v int32) {
	o.Received.Set(&v)
}
// SetReceivedNil sets the value for Received to be an explicit nil
func (o *DeliveryData) SetReceivedNil() {
	o.Received.Set(nil)
}

// UnsetReceived ensures that no value is present for Received, not even an explicit nil
func (o *DeliveryData) UnsetReceived() {
	o.Received.Unset()
}

func (o DeliveryData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Successful.IsSet() {
		toSerialize["successful"] = o.Successful.Get()
	}
	if o.Failed.IsSet() {
		toSerialize["failed"] = o.Failed.Get()
	}
	if o.Errored.IsSet() {
		toSerialize["errored"] = o.Errored.Get()
	}
	if o.Converted.IsSet() {
		toSerialize["converted"] = o.Converted.Get()
	}
	if o.Received.IsSet() {
		toSerialize["received"] = o.Received.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DeliveryData) UnmarshalJSON(bytes []byte) (err error) {
	varDeliveryData := _DeliveryData{}

	if err = json.Unmarshal(bytes, &varDeliveryData); err == nil {
		*o = DeliveryData(varDeliveryData)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "successful")
		delete(additionalProperties, "failed")
		delete(additionalProperties, "errored")
		delete(additionalProperties, "converted")
		delete(additionalProperties, "received")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeliveryData struct {
	value *DeliveryData
	isSet bool
}

func (v NullableDeliveryData) Get() *DeliveryData {
	return v.value
}

func (v *NullableDeliveryData) Set(val *DeliveryData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryData(val *DeliveryData) *NullableDeliveryData {
	return &NullableDeliveryData{value: val, isSet: true}
}

func (v NullableDeliveryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliveryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


