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

// BasicNotificationAllOfAndroidBackgroundLayout Channel: Push Notifications Platform: Android Allowing setting a background image for the notification. This is a JSON object containing the following keys. See our Background Image documentation for image sizes. 
type BasicNotificationAllOfAndroidBackgroundLayout struct {
	// Asset file, android resource name, or URL to remote image.
	Image *string `json:"image,omitempty"`
	// Title text color ARGB Hex format. Example(Blue) \"FF0000FF\".
	HeadingsColor *string `json:"headings_color,omitempty"`
	// Body text color ARGB Hex format. Example(Red) \"FFFF0000\".
	ContentsColor *string `json:"contents_color,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BasicNotificationAllOfAndroidBackgroundLayout BasicNotificationAllOfAndroidBackgroundLayout

// NewBasicNotificationAllOfAndroidBackgroundLayout instantiates a new BasicNotificationAllOfAndroidBackgroundLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicNotificationAllOfAndroidBackgroundLayout() *BasicNotificationAllOfAndroidBackgroundLayout {
	this := BasicNotificationAllOfAndroidBackgroundLayout{}
	return &this
}

// NewBasicNotificationAllOfAndroidBackgroundLayoutWithDefaults instantiates a new BasicNotificationAllOfAndroidBackgroundLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicNotificationAllOfAndroidBackgroundLayoutWithDefaults() *BasicNotificationAllOfAndroidBackgroundLayout {
	this := BasicNotificationAllOfAndroidBackgroundLayout{}
	return &this
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) SetImage(v string) {
	o.Image = &v
}

// GetHeadingsColor returns the HeadingsColor field value if set, zero value otherwise.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetHeadingsColor() string {
	if o == nil || o.HeadingsColor == nil {
		var ret string
		return ret
	}
	return *o.HeadingsColor
}

// GetHeadingsColorOk returns a tuple with the HeadingsColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetHeadingsColorOk() (*string, bool) {
	if o == nil || o.HeadingsColor == nil {
		return nil, false
	}
	return o.HeadingsColor, true
}

// HasHeadingsColor returns a boolean if a field has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) HasHeadingsColor() bool {
	if o != nil && o.HeadingsColor != nil {
		return true
	}

	return false
}

// SetHeadingsColor gets a reference to the given string and assigns it to the HeadingsColor field.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) SetHeadingsColor(v string) {
	o.HeadingsColor = &v
}

// GetContentsColor returns the ContentsColor field value if set, zero value otherwise.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetContentsColor() string {
	if o == nil || o.ContentsColor == nil {
		var ret string
		return ret
	}
	return *o.ContentsColor
}

// GetContentsColorOk returns a tuple with the ContentsColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) GetContentsColorOk() (*string, bool) {
	if o == nil || o.ContentsColor == nil {
		return nil, false
	}
	return o.ContentsColor, true
}

// HasContentsColor returns a boolean if a field has been set.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) HasContentsColor() bool {
	if o != nil && o.ContentsColor != nil {
		return true
	}

	return false
}

// SetContentsColor gets a reference to the given string and assigns it to the ContentsColor field.
func (o *BasicNotificationAllOfAndroidBackgroundLayout) SetContentsColor(v string) {
	o.ContentsColor = &v
}

func (o BasicNotificationAllOfAndroidBackgroundLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.HeadingsColor != nil {
		toSerialize["headings_color"] = o.HeadingsColor
	}
	if o.ContentsColor != nil {
		toSerialize["contents_color"] = o.ContentsColor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BasicNotificationAllOfAndroidBackgroundLayout) UnmarshalJSON(bytes []byte) (err error) {
	varBasicNotificationAllOfAndroidBackgroundLayout := _BasicNotificationAllOfAndroidBackgroundLayout{}

	if err = json.Unmarshal(bytes, &varBasicNotificationAllOfAndroidBackgroundLayout); err == nil {
		*o = BasicNotificationAllOfAndroidBackgroundLayout(varBasicNotificationAllOfAndroidBackgroundLayout)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "image")
		delete(additionalProperties, "headings_color")
		delete(additionalProperties, "contents_color")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBasicNotificationAllOfAndroidBackgroundLayout struct {
	value *BasicNotificationAllOfAndroidBackgroundLayout
	isSet bool
}

func (v NullableBasicNotificationAllOfAndroidBackgroundLayout) Get() *BasicNotificationAllOfAndroidBackgroundLayout {
	return v.value
}

func (v *NullableBasicNotificationAllOfAndroidBackgroundLayout) Set(val *BasicNotificationAllOfAndroidBackgroundLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicNotificationAllOfAndroidBackgroundLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicNotificationAllOfAndroidBackgroundLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicNotificationAllOfAndroidBackgroundLayout(val *BasicNotificationAllOfAndroidBackgroundLayout) *NullableBasicNotificationAllOfAndroidBackgroundLayout {
	return &NullableBasicNotificationAllOfAndroidBackgroundLayout{value: val, isSet: true}
}

func (v NullableBasicNotificationAllOfAndroidBackgroundLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicNotificationAllOfAndroidBackgroundLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


