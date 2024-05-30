/*
Libre Graph API

Libre Graph is a free API for cloud collaboration inspired by the MS Graph API.

API version: v1.0.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package libregraph

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TagAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagAssignment{}

// TagAssignment struct for TagAssignment
type TagAssignment struct {
	ResourceId string   `json:"resourceId"`
	Tags       []string `json:"tags"`
}

type _TagAssignment TagAssignment

// NewTagAssignment instantiates a new TagAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagAssignment(resourceId string, tags []string) *TagAssignment {
	this := TagAssignment{}
	this.ResourceId = resourceId
	this.Tags = tags
	return &this
}

// NewTagAssignmentWithDefaults instantiates a new TagAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagAssignmentWithDefaults() *TagAssignment {
	this := TagAssignment{}
	return &this
}

// GetResourceId returns the ResourceId field value
func (o *TagAssignment) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *TagAssignment) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value
func (o *TagAssignment) SetResourceId(v string) {
	o.ResourceId = v
}

// GetTags returns the Tags field value
func (o *TagAssignment) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *TagAssignment) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *TagAssignment) SetTags(v []string) {
	o.Tags = v
}

func (o TagAssignment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceId"] = o.ResourceId
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *TagAssignment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"resourceId",
		"tags",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTagAssignment := _TagAssignment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagAssignment)

	if err != nil {
		return err
	}

	*o = TagAssignment(varTagAssignment)

	return err
}

type NullableTagAssignment struct {
	value *TagAssignment
	isSet bool
}

func (v NullableTagAssignment) Get() *TagAssignment {
	return v.value
}

func (v *NullableTagAssignment) Set(val *TagAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableTagAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableTagAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagAssignment(val *TagAssignment) *NullableTagAssignment {
	return &NullableTagAssignment{value: val, isSet: true}
}

func (v NullableTagAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
