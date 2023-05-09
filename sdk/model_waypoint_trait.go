/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the WaypointTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WaypointTrait{}

// WaypointTrait struct for WaypointTrait
type WaypointTrait struct {
	// The unique identifier of the trait.
	Symbol string `json:"symbol"`
	// The name of the trait.
	Name string `json:"name"`
	// A description of the trait.
	Description string `json:"description"`
}

// NewWaypointTrait instantiates a new WaypointTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWaypointTrait(symbol string, name string, description string) *WaypointTrait {
	this := WaypointTrait{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	return &this
}

// NewWaypointTraitWithDefaults instantiates a new WaypointTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWaypointTraitWithDefaults() *WaypointTrait {
	this := WaypointTrait{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *WaypointTrait) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *WaypointTrait) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *WaypointTrait) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *WaypointTrait) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WaypointTrait) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WaypointTrait) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *WaypointTrait) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *WaypointTrait) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *WaypointTrait) SetDescription(v string) {
	o.Description = v
}

func (o WaypointTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WaypointTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullableWaypointTrait struct {
	value *WaypointTrait
	isSet bool
}

func (v NullableWaypointTrait) Get() *WaypointTrait {
	return v.value
}

func (v *NullableWaypointTrait) Set(val *WaypointTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableWaypointTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableWaypointTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWaypointTrait(val *WaypointTrait) *NullableWaypointTrait {
	return &NullableWaypointTrait{value: val, isSet: true}
}

func (v NullableWaypointTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWaypointTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

