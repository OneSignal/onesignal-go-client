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

// PlatformDeliveryDataSmsAllOf struct for PlatformDeliveryDataSmsAllOf
type PlatformDeliveryDataSmsAllOf struct {
	// Number of messages reported as delivered successfully by the SMS service provider.
	ProviderSuccessful NullableInt32 `json:"provider_successful,omitempty"`
	// Number of recipients who didn't receive your message as reported by the SMS service provider.
	ProviderFailed NullableInt32 `json:"provider_failed,omitempty"`
	// Number of errors reported by the SMS service provider.
	ProviderErrored NullableInt32 `json:"provider_errored,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformDeliveryDataSmsAllOf PlatformDeliveryDataSmsAllOf

// NewPlatformDeliveryDataSmsAllOf instantiates a new PlatformDeliveryDataSmsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDeliveryDataSmsAllOf() *PlatformDeliveryDataSmsAllOf {
	this := PlatformDeliveryDataSmsAllOf{}
	return &this
}

// NewPlatformDeliveryDataSmsAllOfWithDefaults instantiates a new PlatformDeliveryDataSmsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDeliveryDataSmsAllOfWithDefaults() *PlatformDeliveryDataSmsAllOf {
	this := PlatformDeliveryDataSmsAllOf{}
	return &this
}

// GetProviderSuccessful returns the ProviderSuccessful field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataSmsAllOf) GetProviderSuccessful() int32 {
	if o == nil || o.ProviderSuccessful.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ProviderSuccessful.Get()
}

// GetProviderSuccessfulOk returns a tuple with the ProviderSuccessful field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataSmsAllOf) GetProviderSuccessfulOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderSuccessful.Get(), o.ProviderSuccessful.IsSet()
}

// HasProviderSuccessful returns a boolean if a field has been set.
func (o *PlatformDeliveryDataSmsAllOf) HasProviderSuccessful() bool {
	if o != nil && o.ProviderSuccessful.IsSet() {
		return true
	}

	return false
}

// SetProviderSuccessful gets a reference to the given NullableInt32 and assigns it to the ProviderSuccessful field.
func (o *PlatformDeliveryDataSmsAllOf) SetProviderSuccessful(v int32) {
	o.ProviderSuccessful.Set(&v)
}
// SetProviderSuccessfulNil sets the value for ProviderSuccessful to be an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) SetProviderSuccessfulNil() {
	o.ProviderSuccessful.Set(nil)
}

// UnsetProviderSuccessful ensures that no value is present for ProviderSuccessful, not even an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) UnsetProviderSuccessful() {
	o.ProviderSuccessful.Unset()
}

// GetProviderFailed returns the ProviderFailed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataSmsAllOf) GetProviderFailed() int32 {
	if o == nil || o.ProviderFailed.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ProviderFailed.Get()
}

// GetProviderFailedOk returns a tuple with the ProviderFailed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataSmsAllOf) GetProviderFailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderFailed.Get(), o.ProviderFailed.IsSet()
}

// HasProviderFailed returns a boolean if a field has been set.
func (o *PlatformDeliveryDataSmsAllOf) HasProviderFailed() bool {
	if o != nil && o.ProviderFailed.IsSet() {
		return true
	}

	return false
}

// SetProviderFailed gets a reference to the given NullableInt32 and assigns it to the ProviderFailed field.
func (o *PlatformDeliveryDataSmsAllOf) SetProviderFailed(v int32) {
	o.ProviderFailed.Set(&v)
}
// SetProviderFailedNil sets the value for ProviderFailed to be an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) SetProviderFailedNil() {
	o.ProviderFailed.Set(nil)
}

// UnsetProviderFailed ensures that no value is present for ProviderFailed, not even an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) UnsetProviderFailed() {
	o.ProviderFailed.Unset()
}

// GetProviderErrored returns the ProviderErrored field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDeliveryDataSmsAllOf) GetProviderErrored() int32 {
	if o == nil || o.ProviderErrored.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ProviderErrored.Get()
}

// GetProviderErroredOk returns a tuple with the ProviderErrored field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDeliveryDataSmsAllOf) GetProviderErroredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderErrored.Get(), o.ProviderErrored.IsSet()
}

// HasProviderErrored returns a boolean if a field has been set.
func (o *PlatformDeliveryDataSmsAllOf) HasProviderErrored() bool {
	if o != nil && o.ProviderErrored.IsSet() {
		return true
	}

	return false
}

// SetProviderErrored gets a reference to the given NullableInt32 and assigns it to the ProviderErrored field.
func (o *PlatformDeliveryDataSmsAllOf) SetProviderErrored(v int32) {
	o.ProviderErrored.Set(&v)
}
// SetProviderErroredNil sets the value for ProviderErrored to be an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) SetProviderErroredNil() {
	o.ProviderErrored.Set(nil)
}

// UnsetProviderErrored ensures that no value is present for ProviderErrored, not even an explicit nil
func (o *PlatformDeliveryDataSmsAllOf) UnsetProviderErrored() {
	o.ProviderErrored.Unset()
}

func (o PlatformDeliveryDataSmsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProviderSuccessful.IsSet() {
		toSerialize["provider_successful"] = o.ProviderSuccessful.Get()
	}
	if o.ProviderFailed.IsSet() {
		toSerialize["provider_failed"] = o.ProviderFailed.Get()
	}
	if o.ProviderErrored.IsSet() {
		toSerialize["provider_errored"] = o.ProviderErrored.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PlatformDeliveryDataSmsAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformDeliveryDataSmsAllOf := _PlatformDeliveryDataSmsAllOf{}

	if err = json.Unmarshal(bytes, &varPlatformDeliveryDataSmsAllOf); err == nil {
		*o = PlatformDeliveryDataSmsAllOf(varPlatformDeliveryDataSmsAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "provider_successful")
		delete(additionalProperties, "provider_failed")
		delete(additionalProperties, "provider_errored")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformDeliveryDataSmsAllOf struct {
	value *PlatformDeliveryDataSmsAllOf
	isSet bool
}

func (v NullablePlatformDeliveryDataSmsAllOf) Get() *PlatformDeliveryDataSmsAllOf {
	return v.value
}

func (v *NullablePlatformDeliveryDataSmsAllOf) Set(val *PlatformDeliveryDataSmsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformDeliveryDataSmsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformDeliveryDataSmsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformDeliveryDataSmsAllOf(val *PlatformDeliveryDataSmsAllOf) *NullablePlatformDeliveryDataSmsAllOf {
	return &NullablePlatformDeliveryDataSmsAllOf{value: val, isSet: true}
}

func (v NullablePlatformDeliveryDataSmsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformDeliveryDataSmsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


