/*
 * Merge Ticketing API
 *
 * The unified API for building rich integrations with multiple Ticketing platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_ticketing_client

import (
	"encoding/json"
	"time"
)

// SyncStatus # The SyncStatus Object ### Description The `SyncStatus` object is used to represent the syncing state of an account  ### Usage Example View the `SyncStatus` for an account to see how recently its models were synced.
type SyncStatus struct {
	ModelName string `json:"model_name"`
	ModelId string `json:"model_id"`
	LastSyncStart *time.Time `json:"last_sync_start,omitempty"`
	NextSyncStart *time.Time `json:"next_sync_start,omitempty"`
	Status SyncStatusStatusEnum `json:"status"`
	IsInitialSync bool `json:"is_initial_sync"`
    // raw json response by property name
    ResponseRaw map[string]json.RawMessage `json:"-"`
}

// NewSyncStatus instantiates a new SyncStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyncStatus(modelName string, modelId string, status SyncStatusStatusEnum, isInitialSync bool) *SyncStatus {
	this := SyncStatus{}
	this.ModelName = modelName
	this.ModelId = modelId
	this.Status = status
	this.IsInitialSync = isInitialSync
	return &this
}

// NewSyncStatusWithDefaults instantiates a new SyncStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyncStatusWithDefaults() *SyncStatus {
	this := SyncStatus{}
	return &this
}

// GetModelName returns the ModelName field value
func (o *SyncStatus) GetModelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelName
}

// GetModelNameOk returns a tuple with the ModelName field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetModelNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelName, true
}

// SetModelName sets field value
func (o *SyncStatus) SetModelName(v string) {
	o.ModelName = v
}

// GetModelId returns the ModelId field value
func (o *SyncStatus) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetModelIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *SyncStatus) SetModelId(v string) {
	o.ModelId = v
}

// GetLastSyncStart returns the LastSyncStart field value if set, zero value otherwise.
func (o *SyncStatus) GetLastSyncStart() time.Time {
	if o == nil || o.LastSyncStart == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncStart
}

// GetLastSyncStartOk returns a tuple with the LastSyncStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetLastSyncStartOk() (*time.Time, bool) {
	if o == nil || o.LastSyncStart == nil {
		return nil, false
	}
	return o.LastSyncStart, true
}

// HasLastSyncStart returns a boolean if a field has been set.
func (o *SyncStatus) HasLastSyncStart() bool {
	if o != nil && o.LastSyncStart != nil {
		return true
	}

	return false
}

// SetLastSyncStart gets a reference to the given time.Time and assigns it to the LastSyncStart field.
func (o *SyncStatus) SetLastSyncStart(v time.Time) {
	o.LastSyncStart = &v
}

// GetNextSyncStart returns the NextSyncStart field value if set, zero value otherwise.
func (o *SyncStatus) GetNextSyncStart() time.Time {
	if o == nil || o.NextSyncStart == nil {
		var ret time.Time
		return ret
	}
	return *o.NextSyncStart
}

// GetNextSyncStartOk returns a tuple with the NextSyncStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetNextSyncStartOk() (*time.Time, bool) {
	if o == nil || o.NextSyncStart == nil {
		return nil, false
	}
	return o.NextSyncStart, true
}

// HasNextSyncStart returns a boolean if a field has been set.
func (o *SyncStatus) HasNextSyncStart() bool {
	if o != nil && o.NextSyncStart != nil {
		return true
	}

	return false
}

// SetNextSyncStart gets a reference to the given time.Time and assigns it to the NextSyncStart field.
func (o *SyncStatus) SetNextSyncStart(v time.Time) {
	o.NextSyncStart = &v
}

// GetStatus returns the Status field value
func (o *SyncStatus) GetStatus() SyncStatusStatusEnum {
	if o == nil {
		var ret SyncStatusStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetStatusOk() (*SyncStatusStatusEnum, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SyncStatus) SetStatus(v SyncStatusStatusEnum) {
	o.Status = v
}

// GetIsInitialSync returns the IsInitialSync field value
func (o *SyncStatus) GetIsInitialSync() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsInitialSync
}

// GetIsInitialSyncOk returns a tuple with the IsInitialSync field value
// and a boolean to check if the value has been set.
func (o *SyncStatus) GetIsInitialSyncOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsInitialSync, true
}

// SetIsInitialSync sets field value
func (o *SyncStatus) SetIsInitialSync(v bool) {
	o.IsInitialSync = v
}

func (o SyncStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["model_name"] = o.ModelName
	}
	if true {
		toSerialize["model_id"] = o.ModelId
	}
	if o.LastSyncStart != nil {
		toSerialize["last_sync_start"] = o.LastSyncStart
	}
	if o.NextSyncStart != nil {
		toSerialize["next_sync_start"] = o.NextSyncStart
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["is_initial_sync"] = o.IsInitialSync
	}
	return json.Marshal(toSerialize)
}

func (v *SyncStatus) UnmarshalJSON(src []byte) error {
    type SyncStatusUnmarshalTarget SyncStatus

	var intermediateResult SyncStatusUnmarshalTarget
	var err1 = json.Unmarshal(src, &intermediateResult)
    if err1 != nil {
        return err1
    }
    var err2 = json.Unmarshal(src, &intermediateResult.ResponseRaw)
	if err2 != nil {
		return err2
	}

	*v = SyncStatus(intermediateResult)
	return nil
}
type NullableSyncStatus struct {
	value *SyncStatus
	isSet bool
}

func (v NullableSyncStatus) Get() *SyncStatus {
	return v.value
}

func (v *NullableSyncStatus) Set(val *SyncStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSyncStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSyncStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyncStatus(val *SyncStatus) *NullableSyncStatus {
	return &NullableSyncStatus{value: val, isSet: true}
}

func (v NullableSyncStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyncStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	var err1 = json.Unmarshal(src, &v.value)
    if err1 != nil {
        return err1
    }
    return json.Unmarshal(src, &v.value.ResponseRaw)
}


