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

// Player struct for Player
type Player struct {
	// The device's OneSignal ID
	Id string `json:"id"`
	// If true, this is the equivalent of a user being Unsubscribed
	InvalidIdentifier *bool `json:"invalid_identifier,omitempty"`
	AppId *string `json:"app_id,omitempty"`
	// Required The device's platform:   0 = iOS   1 = Android   2 = Amazon   3 = WindowsPhone (MPNS)   4 = Chrome Apps / Extensions   5 = Chrome Web Push   6 = Windows (WNS)   7 = Safari   8 = Firefox   9 = MacOS   10 = Alexa   11 = Email   13 = For Huawei App Gallery Builds SDK Setup. Not for Huawei Devices using FCM   14 = SMS 
	DeviceType int32 `json:"device_type"`
	// a custom user ID
	ExternalUserId NullableString `json:"external_user_id,omitempty"`
	// Only required if you have enabled Identity Verification and device_type is NOT 11 email.
	ExternalUserIdAuthHash *string `json:"external_user_id_auth_hash,omitempty"`
	// Email - Only required if you have enabled Identity Verification and device_type is email (11).
	EmailAuthHash *string `json:"email_auth_hash,omitempty"`
	// Recommended: For Push Notifications, this is the Push Token Identifier from Google or Apple. For Apple Push identifiers, you must strip all non alphanumeric characters. Examples: iOS: 7abcd558f29d0b1f048083e2834ad8ea4b3d87d8ad9c088b33c132706ff445f0 Android: APA91bHbYHk7aq-Uam_2pyJ2qbZvqllyyh2wjfPRaw5gLEX2SUlQBRvOc6sck1sa7H7nGeLNlDco8lXj83HWWwzV... For Email Addresses, set the full email address email@email.com and make sure to set device_type to 11. 
	Identifier NullableString `json:"identifier,omitempty"`
	// Language code. Typically lower case two letters, except for Chinese where it must be one of zh-Hans or zh-Hant. Example: en 
	Language *string `json:"language,omitempty"`
	// Number of seconds away from UTC. Example: -28800 
	Timezone NullableInt32 `json:"timezone,omitempty"`
	// Version of your app. Example: 1.1 
	GameVersion NullableString `json:"game_version,omitempty"`
	// Device make and model. Example: iPhone5,1 
	DeviceModel NullableString `json:"device_model,omitempty"`
	// Device operating system version. Example: 7.0.4 
	DeviceOs NullableString `json:"device_os,omitempty"`
	// The ad id for the device's platform: Android = Advertising Id iOS = identifierForVendor WP8.0 = DeviceUniqueId WP8.1 = AdvertisingId 
	AdId NullableString `json:"ad_id,omitempty"`
	// Name and version of the sdk/plugin that's calling this API method (if any)
	Sdk NullableString `json:"sdk,omitempty"`
	// Number of times the user has played the game, defaults to 1
	SessionCount *int32 `json:"session_count,omitempty"`
	// Custom tags for the player. Only support string and integer key value pairs. Does not support arrays or other nested objects. Setting a tag value to null or an empty string will remove the tag. Example: {\"foo\":\"bar\",\"this\":\"that\"} Limitations: - 100 tags per call - Android SDK users: tags cannot be removed or changed via API if set through SDK sendTag methods. Recommended to only tag devices with 1 kilobyte of data Please consider using your own Database to save more than 1 kilobyte of data. See: Internal Database & CRM 
	Tags map[string]interface{} `json:"tags,omitempty"`
	// Amount the user has spent in USD, up to two decimal places
	AmountSpent *float32 `json:"amount_spent,omitempty"`
	// Unixtime when the player joined the game
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Seconds player was running your app.
	Playtime *int64 `json:"playtime,omitempty"`
	// Current iOS badge count displayed on the app icon NOTE: Not supported for apps created after June 2018, since badge count for apps created after this date are handled on the client. 
	BadgeCount *int32 `json:"badge_count,omitempty"`
	// Unixtime when the player was last active
	LastActive *int32 `json:"last_active,omitempty"`
	// 1 = subscribed -2 = unsubscribed iOS - These values are set each time the user opens the app from the SDK. Use the SDK function set Subscription instead. Android - You may set this but you can no longer use the SDK method setSubscription later in your app as it will create synchronization issues. 
	NotificationTypes *int32 `json:"notification_types,omitempty"`
	// This is used in deciding whether to use your iOS Sandbox or Production push certificate when sending a push when both have been uploaded. Set to the iOS provisioning profile that was used to build your app. 1 = Development 2 = Ad-Hoc Omit this field for App Store builds. 
	TestType NullableInt32 `json:"test_type,omitempty"`
	// Longitude of the device, used for geotagging to segment on.
	Long *float32 `json:"long,omitempty"`
	// Latitude of the device, used for geotagging to segment on.
	Lat *float32 `json:"lat,omitempty"`
	// Country code in the ISO 3166-1 Alpha 2 format
	Country *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Player Player

// NewPlayer instantiates a new Player object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlayer(id string, deviceType int32) *Player {
	this := Player{}
	this.Id = id
	this.DeviceType = deviceType
	return &this
}

// NewPlayerWithDefaults instantiates a new Player object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlayerWithDefaults() *Player {
	this := Player{}
	return &this
}

// GetId returns the Id field value
func (o *Player) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Player) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Player) SetId(v string) {
	o.Id = v
}

// GetInvalidIdentifier returns the InvalidIdentifier field value if set, zero value otherwise.
func (o *Player) GetInvalidIdentifier() bool {
	if o == nil || o.InvalidIdentifier == nil {
		var ret bool
		return ret
	}
	return *o.InvalidIdentifier
}

// GetInvalidIdentifierOk returns a tuple with the InvalidIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetInvalidIdentifierOk() (*bool, bool) {
	if o == nil || o.InvalidIdentifier == nil {
		return nil, false
	}
	return o.InvalidIdentifier, true
}

// HasInvalidIdentifier returns a boolean if a field has been set.
func (o *Player) HasInvalidIdentifier() bool {
	if o != nil && o.InvalidIdentifier != nil {
		return true
	}

	return false
}

// SetInvalidIdentifier gets a reference to the given bool and assigns it to the InvalidIdentifier field.
func (o *Player) SetInvalidIdentifier(v bool) {
	o.InvalidIdentifier = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *Player) GetAppId() string {
	if o == nil || o.AppId == nil {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetAppIdOk() (*string, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *Player) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *Player) SetAppId(v string) {
	o.AppId = &v
}

// GetDeviceType returns the DeviceType field value
func (o *Player) GetDeviceType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value
// and a boolean to check if the value has been set.
func (o *Player) GetDeviceTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceType, true
}

// SetDeviceType sets field value
func (o *Player) SetDeviceType(v int32) {
	o.DeviceType = v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetExternalUserId() string {
	if o == nil || o.ExternalUserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExternalUserId.Get()
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetExternalUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExternalUserId.Get(), o.ExternalUserId.IsSet()
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *Player) HasExternalUserId() bool {
	if o != nil && o.ExternalUserId.IsSet() {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given NullableString and assigns it to the ExternalUserId field.
func (o *Player) SetExternalUserId(v string) {
	o.ExternalUserId.Set(&v)
}
// SetExternalUserIdNil sets the value for ExternalUserId to be an explicit nil
func (o *Player) SetExternalUserIdNil() {
	o.ExternalUserId.Set(nil)
}

// UnsetExternalUserId ensures that no value is present for ExternalUserId, not even an explicit nil
func (o *Player) UnsetExternalUserId() {
	o.ExternalUserId.Unset()
}

// GetExternalUserIdAuthHash returns the ExternalUserIdAuthHash field value if set, zero value otherwise.
func (o *Player) GetExternalUserIdAuthHash() string {
	if o == nil || o.ExternalUserIdAuthHash == nil {
		var ret string
		return ret
	}
	return *o.ExternalUserIdAuthHash
}

// GetExternalUserIdAuthHashOk returns a tuple with the ExternalUserIdAuthHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetExternalUserIdAuthHashOk() (*string, bool) {
	if o == nil || o.ExternalUserIdAuthHash == nil {
		return nil, false
	}
	return o.ExternalUserIdAuthHash, true
}

// HasExternalUserIdAuthHash returns a boolean if a field has been set.
func (o *Player) HasExternalUserIdAuthHash() bool {
	if o != nil && o.ExternalUserIdAuthHash != nil {
		return true
	}

	return false
}

// SetExternalUserIdAuthHash gets a reference to the given string and assigns it to the ExternalUserIdAuthHash field.
func (o *Player) SetExternalUserIdAuthHash(v string) {
	o.ExternalUserIdAuthHash = &v
}

// GetEmailAuthHash returns the EmailAuthHash field value if set, zero value otherwise.
func (o *Player) GetEmailAuthHash() string {
	if o == nil || o.EmailAuthHash == nil {
		var ret string
		return ret
	}
	return *o.EmailAuthHash
}

// GetEmailAuthHashOk returns a tuple with the EmailAuthHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetEmailAuthHashOk() (*string, bool) {
	if o == nil || o.EmailAuthHash == nil {
		return nil, false
	}
	return o.EmailAuthHash, true
}

// HasEmailAuthHash returns a boolean if a field has been set.
func (o *Player) HasEmailAuthHash() bool {
	if o != nil && o.EmailAuthHash != nil {
		return true
	}

	return false
}

// SetEmailAuthHash gets a reference to the given string and assigns it to the EmailAuthHash field.
func (o *Player) SetEmailAuthHash(v string) {
	o.EmailAuthHash = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetIdentifier() string {
	if o == nil || o.Identifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.Identifier.Get()
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identifier.Get(), o.Identifier.IsSet()
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Player) HasIdentifier() bool {
	if o != nil && o.Identifier.IsSet() {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given NullableString and assigns it to the Identifier field.
func (o *Player) SetIdentifier(v string) {
	o.Identifier.Set(&v)
}
// SetIdentifierNil sets the value for Identifier to be an explicit nil
func (o *Player) SetIdentifierNil() {
	o.Identifier.Set(nil)
}

// UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
func (o *Player) UnsetIdentifier() {
	o.Identifier.Unset()
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Player) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Player) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Player) SetLanguage(v string) {
	o.Language = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetTimezone() int32 {
	if o == nil || o.Timezone.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Timezone.Get()
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetTimezoneOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timezone.Get(), o.Timezone.IsSet()
}

// HasTimezone returns a boolean if a field has been set.
func (o *Player) HasTimezone() bool {
	if o != nil && o.Timezone.IsSet() {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given NullableInt32 and assigns it to the Timezone field.
func (o *Player) SetTimezone(v int32) {
	o.Timezone.Set(&v)
}
// SetTimezoneNil sets the value for Timezone to be an explicit nil
func (o *Player) SetTimezoneNil() {
	o.Timezone.Set(nil)
}

// UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil
func (o *Player) UnsetTimezone() {
	o.Timezone.Unset()
}

// GetGameVersion returns the GameVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetGameVersion() string {
	if o == nil || o.GameVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.GameVersion.Get()
}

// GetGameVersionOk returns a tuple with the GameVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetGameVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GameVersion.Get(), o.GameVersion.IsSet()
}

// HasGameVersion returns a boolean if a field has been set.
func (o *Player) HasGameVersion() bool {
	if o != nil && o.GameVersion.IsSet() {
		return true
	}

	return false
}

// SetGameVersion gets a reference to the given NullableString and assigns it to the GameVersion field.
func (o *Player) SetGameVersion(v string) {
	o.GameVersion.Set(&v)
}
// SetGameVersionNil sets the value for GameVersion to be an explicit nil
func (o *Player) SetGameVersionNil() {
	o.GameVersion.Set(nil)
}

// UnsetGameVersion ensures that no value is present for GameVersion, not even an explicit nil
func (o *Player) UnsetGameVersion() {
	o.GameVersion.Unset()
}

// GetDeviceModel returns the DeviceModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetDeviceModel() string {
	if o == nil || o.DeviceModel.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceModel.Get()
}

// GetDeviceModelOk returns a tuple with the DeviceModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetDeviceModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceModel.Get(), o.DeviceModel.IsSet()
}

// HasDeviceModel returns a boolean if a field has been set.
func (o *Player) HasDeviceModel() bool {
	if o != nil && o.DeviceModel.IsSet() {
		return true
	}

	return false
}

// SetDeviceModel gets a reference to the given NullableString and assigns it to the DeviceModel field.
func (o *Player) SetDeviceModel(v string) {
	o.DeviceModel.Set(&v)
}
// SetDeviceModelNil sets the value for DeviceModel to be an explicit nil
func (o *Player) SetDeviceModelNil() {
	o.DeviceModel.Set(nil)
}

// UnsetDeviceModel ensures that no value is present for DeviceModel, not even an explicit nil
func (o *Player) UnsetDeviceModel() {
	o.DeviceModel.Unset()
}

// GetDeviceOs returns the DeviceOs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetDeviceOs() string {
	if o == nil || o.DeviceOs.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceOs.Get()
}

// GetDeviceOsOk returns a tuple with the DeviceOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetDeviceOsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeviceOs.Get(), o.DeviceOs.IsSet()
}

// HasDeviceOs returns a boolean if a field has been set.
func (o *Player) HasDeviceOs() bool {
	if o != nil && o.DeviceOs.IsSet() {
		return true
	}

	return false
}

// SetDeviceOs gets a reference to the given NullableString and assigns it to the DeviceOs field.
func (o *Player) SetDeviceOs(v string) {
	o.DeviceOs.Set(&v)
}
// SetDeviceOsNil sets the value for DeviceOs to be an explicit nil
func (o *Player) SetDeviceOsNil() {
	o.DeviceOs.Set(nil)
}

// UnsetDeviceOs ensures that no value is present for DeviceOs, not even an explicit nil
func (o *Player) UnsetDeviceOs() {
	o.DeviceOs.Unset()
}

// GetAdId returns the AdId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetAdId() string {
	if o == nil || o.AdId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AdId.Get()
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdId.Get(), o.AdId.IsSet()
}

// HasAdId returns a boolean if a field has been set.
func (o *Player) HasAdId() bool {
	if o != nil && o.AdId.IsSet() {
		return true
	}

	return false
}

// SetAdId gets a reference to the given NullableString and assigns it to the AdId field.
func (o *Player) SetAdId(v string) {
	o.AdId.Set(&v)
}
// SetAdIdNil sets the value for AdId to be an explicit nil
func (o *Player) SetAdIdNil() {
	o.AdId.Set(nil)
}

// UnsetAdId ensures that no value is present for AdId, not even an explicit nil
func (o *Player) UnsetAdId() {
	o.AdId.Unset()
}

// GetSdk returns the Sdk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetSdk() string {
	if o == nil || o.Sdk.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sdk.Get()
}

// GetSdkOk returns a tuple with the Sdk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetSdkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sdk.Get(), o.Sdk.IsSet()
}

// HasSdk returns a boolean if a field has been set.
func (o *Player) HasSdk() bool {
	if o != nil && o.Sdk.IsSet() {
		return true
	}

	return false
}

// SetSdk gets a reference to the given NullableString and assigns it to the Sdk field.
func (o *Player) SetSdk(v string) {
	o.Sdk.Set(&v)
}
// SetSdkNil sets the value for Sdk to be an explicit nil
func (o *Player) SetSdkNil() {
	o.Sdk.Set(nil)
}

// UnsetSdk ensures that no value is present for Sdk, not even an explicit nil
func (o *Player) UnsetSdk() {
	o.Sdk.Unset()
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *Player) GetSessionCount() int32 {
	if o == nil || o.SessionCount == nil {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetSessionCountOk() (*int32, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *Player) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *Player) SetSessionCount(v int32) {
	o.SessionCount = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetTags() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Player) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *Player) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetAmountSpent returns the AmountSpent field value if set, zero value otherwise.
func (o *Player) GetAmountSpent() float32 {
	if o == nil || o.AmountSpent == nil {
		var ret float32
		return ret
	}
	return *o.AmountSpent
}

// GetAmountSpentOk returns a tuple with the AmountSpent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetAmountSpentOk() (*float32, bool) {
	if o == nil || o.AmountSpent == nil {
		return nil, false
	}
	return o.AmountSpent, true
}

// HasAmountSpent returns a boolean if a field has been set.
func (o *Player) HasAmountSpent() bool {
	if o != nil && o.AmountSpent != nil {
		return true
	}

	return false
}

// SetAmountSpent gets a reference to the given float32 and assigns it to the AmountSpent field.
func (o *Player) SetAmountSpent(v float32) {
	o.AmountSpent = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Player) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Player) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Player) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetPlaytime returns the Playtime field value if set, zero value otherwise.
func (o *Player) GetPlaytime() int64 {
	if o == nil || o.Playtime == nil {
		var ret int64
		return ret
	}
	return *o.Playtime
}

// GetPlaytimeOk returns a tuple with the Playtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetPlaytimeOk() (*int64, bool) {
	if o == nil || o.Playtime == nil {
		return nil, false
	}
	return o.Playtime, true
}

// HasPlaytime returns a boolean if a field has been set.
func (o *Player) HasPlaytime() bool {
	if o != nil && o.Playtime != nil {
		return true
	}

	return false
}

// SetPlaytime gets a reference to the given int64 and assigns it to the Playtime field.
func (o *Player) SetPlaytime(v int64) {
	o.Playtime = &v
}

// GetBadgeCount returns the BadgeCount field value if set, zero value otherwise.
func (o *Player) GetBadgeCount() int32 {
	if o == nil || o.BadgeCount == nil {
		var ret int32
		return ret
	}
	return *o.BadgeCount
}

// GetBadgeCountOk returns a tuple with the BadgeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetBadgeCountOk() (*int32, bool) {
	if o == nil || o.BadgeCount == nil {
		return nil, false
	}
	return o.BadgeCount, true
}

// HasBadgeCount returns a boolean if a field has been set.
func (o *Player) HasBadgeCount() bool {
	if o != nil && o.BadgeCount != nil {
		return true
	}

	return false
}

// SetBadgeCount gets a reference to the given int32 and assigns it to the BadgeCount field.
func (o *Player) SetBadgeCount(v int32) {
	o.BadgeCount = &v
}

// GetLastActive returns the LastActive field value if set, zero value otherwise.
func (o *Player) GetLastActive() int32 {
	if o == nil || o.LastActive == nil {
		var ret int32
		return ret
	}
	return *o.LastActive
}

// GetLastActiveOk returns a tuple with the LastActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetLastActiveOk() (*int32, bool) {
	if o == nil || o.LastActive == nil {
		return nil, false
	}
	return o.LastActive, true
}

// HasLastActive returns a boolean if a field has been set.
func (o *Player) HasLastActive() bool {
	if o != nil && o.LastActive != nil {
		return true
	}

	return false
}

// SetLastActive gets a reference to the given int32 and assigns it to the LastActive field.
func (o *Player) SetLastActive(v int32) {
	o.LastActive = &v
}

// GetNotificationTypes returns the NotificationTypes field value if set, zero value otherwise.
func (o *Player) GetNotificationTypes() int32 {
	if o == nil || o.NotificationTypes == nil {
		var ret int32
		return ret
	}
	return *o.NotificationTypes
}

// GetNotificationTypesOk returns a tuple with the NotificationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetNotificationTypesOk() (*int32, bool) {
	if o == nil || o.NotificationTypes == nil {
		return nil, false
	}
	return o.NotificationTypes, true
}

// HasNotificationTypes returns a boolean if a field has been set.
func (o *Player) HasNotificationTypes() bool {
	if o != nil && o.NotificationTypes != nil {
		return true
	}

	return false
}

// SetNotificationTypes gets a reference to the given int32 and assigns it to the NotificationTypes field.
func (o *Player) SetNotificationTypes(v int32) {
	o.NotificationTypes = &v
}

// GetTestType returns the TestType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Player) GetTestType() int32 {
	if o == nil || o.TestType.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TestType.Get()
}

// GetTestTypeOk returns a tuple with the TestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Player) GetTestTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TestType.Get(), o.TestType.IsSet()
}

// HasTestType returns a boolean if a field has been set.
func (o *Player) HasTestType() bool {
	if o != nil && o.TestType.IsSet() {
		return true
	}

	return false
}

// SetTestType gets a reference to the given NullableInt32 and assigns it to the TestType field.
func (o *Player) SetTestType(v int32) {
	o.TestType.Set(&v)
}
// SetTestTypeNil sets the value for TestType to be an explicit nil
func (o *Player) SetTestTypeNil() {
	o.TestType.Set(nil)
}

// UnsetTestType ensures that no value is present for TestType, not even an explicit nil
func (o *Player) UnsetTestType() {
	o.TestType.Unset()
}

// GetLong returns the Long field value if set, zero value otherwise.
func (o *Player) GetLong() float32 {
	if o == nil || o.Long == nil {
		var ret float32
		return ret
	}
	return *o.Long
}

// GetLongOk returns a tuple with the Long field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetLongOk() (*float32, bool) {
	if o == nil || o.Long == nil {
		return nil, false
	}
	return o.Long, true
}

// HasLong returns a boolean if a field has been set.
func (o *Player) HasLong() bool {
	if o != nil && o.Long != nil {
		return true
	}

	return false
}

// SetLong gets a reference to the given float32 and assigns it to the Long field.
func (o *Player) SetLong(v float32) {
	o.Long = &v
}

// GetLat returns the Lat field value if set, zero value otherwise.
func (o *Player) GetLat() float32 {
	if o == nil || o.Lat == nil {
		var ret float32
		return ret
	}
	return *o.Lat
}

// GetLatOk returns a tuple with the Lat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetLatOk() (*float32, bool) {
	if o == nil || o.Lat == nil {
		return nil, false
	}
	return o.Lat, true
}

// HasLat returns a boolean if a field has been set.
func (o *Player) HasLat() bool {
	if o != nil && o.Lat != nil {
		return true
	}

	return false
}

// SetLat gets a reference to the given float32 and assigns it to the Lat field.
func (o *Player) SetLat(v float32) {
	o.Lat = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Player) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Player) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Player) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Player) SetCountry(v string) {
	o.Country = &v
}

func (o Player) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.InvalidIdentifier != nil {
		toSerialize["invalid_identifier"] = o.InvalidIdentifier
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if true {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.ExternalUserId.IsSet() {
		toSerialize["external_user_id"] = o.ExternalUserId.Get()
	}
	if o.ExternalUserIdAuthHash != nil {
		toSerialize["external_user_id_auth_hash"] = o.ExternalUserIdAuthHash
	}
	if o.EmailAuthHash != nil {
		toSerialize["email_auth_hash"] = o.EmailAuthHash
	}
	if o.Identifier.IsSet() {
		toSerialize["identifier"] = o.Identifier.Get()
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Timezone.IsSet() {
		toSerialize["timezone"] = o.Timezone.Get()
	}
	if o.GameVersion.IsSet() {
		toSerialize["game_version"] = o.GameVersion.Get()
	}
	if o.DeviceModel.IsSet() {
		toSerialize["device_model"] = o.DeviceModel.Get()
	}
	if o.DeviceOs.IsSet() {
		toSerialize["device_os"] = o.DeviceOs.Get()
	}
	if o.AdId.IsSet() {
		toSerialize["ad_id"] = o.AdId.Get()
	}
	if o.Sdk.IsSet() {
		toSerialize["sdk"] = o.Sdk.Get()
	}
	if o.SessionCount != nil {
		toSerialize["session_count"] = o.SessionCount
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.AmountSpent != nil {
		toSerialize["amount_spent"] = o.AmountSpent
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Playtime != nil {
		toSerialize["playtime"] = o.Playtime
	}
	if o.BadgeCount != nil {
		toSerialize["badge_count"] = o.BadgeCount
	}
	if o.LastActive != nil {
		toSerialize["last_active"] = o.LastActive
	}
	if o.NotificationTypes != nil {
		toSerialize["notification_types"] = o.NotificationTypes
	}
	if o.TestType.IsSet() {
		toSerialize["test_type"] = o.TestType.Get()
	}
	if o.Long != nil {
		toSerialize["long"] = o.Long
	}
	if o.Lat != nil {
		toSerialize["lat"] = o.Lat
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Player) UnmarshalJSON(bytes []byte) (err error) {
	varPlayer := _Player{}

	if err = json.Unmarshal(bytes, &varPlayer); err == nil {
		*o = Player(varPlayer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "invalid_identifier")
		delete(additionalProperties, "app_id")
		delete(additionalProperties, "device_type")
		delete(additionalProperties, "external_user_id")
		delete(additionalProperties, "external_user_id_auth_hash")
		delete(additionalProperties, "email_auth_hash")
		delete(additionalProperties, "identifier")
		delete(additionalProperties, "language")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "game_version")
		delete(additionalProperties, "device_model")
		delete(additionalProperties, "device_os")
		delete(additionalProperties, "ad_id")
		delete(additionalProperties, "sdk")
		delete(additionalProperties, "session_count")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "amount_spent")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "playtime")
		delete(additionalProperties, "badge_count")
		delete(additionalProperties, "last_active")
		delete(additionalProperties, "notification_types")
		delete(additionalProperties, "test_type")
		delete(additionalProperties, "long")
		delete(additionalProperties, "lat")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlayer struct {
	value *Player
	isSet bool
}

func (v NullablePlayer) Get() *Player {
	return v.value
}

func (v *NullablePlayer) Set(val *Player) {
	v.value = val
	v.isSet = true
}

func (v NullablePlayer) IsSet() bool {
	return v.isSet
}

func (v *NullablePlayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlayer(val *Player) *NullablePlayer {
	return &NullablePlayer{value: val, isSet: true}
}

func (v NullablePlayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


