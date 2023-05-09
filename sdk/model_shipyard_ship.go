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

// checks if the ShipyardShip type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipyardShip{}

// ShipyardShip 
type ShipyardShip struct {
	Type *ShipType `json:"type,omitempty"`
	Name string `json:"name"`
	Description string `json:"description"`
	PurchasePrice int32 `json:"purchasePrice"`
	Frame ShipFrame `json:"frame"`
	Reactor ShipReactor `json:"reactor"`
	Engine ShipEngine `json:"engine"`
	Modules []ShipModule `json:"modules"`
	Mounts []ShipMount `json:"mounts"`
}

// NewShipyardShip instantiates a new ShipyardShip object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipyardShip(name string, description string, purchasePrice int32, frame ShipFrame, reactor ShipReactor, engine ShipEngine, modules []ShipModule, mounts []ShipMount) *ShipyardShip {
	this := ShipyardShip{}
	this.Name = name
	this.Description = description
	this.PurchasePrice = purchasePrice
	this.Frame = frame
	this.Reactor = reactor
	this.Engine = engine
	this.Modules = modules
	this.Mounts = mounts
	return &this
}

// NewShipyardShipWithDefaults instantiates a new ShipyardShip object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipyardShipWithDefaults() *ShipyardShip {
	this := ShipyardShip{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ShipyardShip) GetType() ShipType {
	if o == nil || IsNil(o.Type) {
		var ret ShipType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetTypeOk() (*ShipType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ShipyardShip) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ShipType and assigns it to the Type field.
func (o *ShipyardShip) SetType(v ShipType) {
	o.Type = &v
}

// GetName returns the Name field value
func (o *ShipyardShip) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipyardShip) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipyardShip) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipyardShip) SetDescription(v string) {
	o.Description = v
}

// GetPurchasePrice returns the PurchasePrice field value
func (o *ShipyardShip) GetPurchasePrice() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PurchasePrice
}

// GetPurchasePriceOk returns a tuple with the PurchasePrice field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetPurchasePriceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PurchasePrice, true
}

// SetPurchasePrice sets field value
func (o *ShipyardShip) SetPurchasePrice(v int32) {
	o.PurchasePrice = v
}

// GetFrame returns the Frame field value
func (o *ShipyardShip) GetFrame() ShipFrame {
	if o == nil {
		var ret ShipFrame
		return ret
	}

	return o.Frame
}

// GetFrameOk returns a tuple with the Frame field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetFrameOk() (*ShipFrame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frame, true
}

// SetFrame sets field value
func (o *ShipyardShip) SetFrame(v ShipFrame) {
	o.Frame = v
}

// GetReactor returns the Reactor field value
func (o *ShipyardShip) GetReactor() ShipReactor {
	if o == nil {
		var ret ShipReactor
		return ret
	}

	return o.Reactor
}

// GetReactorOk returns a tuple with the Reactor field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetReactorOk() (*ShipReactor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reactor, true
}

// SetReactor sets field value
func (o *ShipyardShip) SetReactor(v ShipReactor) {
	o.Reactor = v
}

// GetEngine returns the Engine field value
func (o *ShipyardShip) GetEngine() ShipEngine {
	if o == nil {
		var ret ShipEngine
		return ret
	}

	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetEngineOk() (*ShipEngine, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value
func (o *ShipyardShip) SetEngine(v ShipEngine) {
	o.Engine = v
}

// GetModules returns the Modules field value
func (o *ShipyardShip) GetModules() []ShipModule {
	if o == nil {
		var ret []ShipModule
		return ret
	}

	return o.Modules
}

// GetModulesOk returns a tuple with the Modules field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetModulesOk() ([]ShipModule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modules, true
}

// SetModules sets field value
func (o *ShipyardShip) SetModules(v []ShipModule) {
	o.Modules = v
}

// GetMounts returns the Mounts field value
func (o *ShipyardShip) GetMounts() []ShipMount {
	if o == nil {
		var ret []ShipMount
		return ret
	}

	return o.Mounts
}

// GetMountsOk returns a tuple with the Mounts field value
// and a boolean to check if the value has been set.
func (o *ShipyardShip) GetMountsOk() ([]ShipMount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mounts, true
}

// SetMounts sets field value
func (o *ShipyardShip) SetMounts(v []ShipMount) {
	o.Mounts = v
}

func (o ShipyardShip) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipyardShip) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["purchasePrice"] = o.PurchasePrice
	toSerialize["frame"] = o.Frame
	toSerialize["reactor"] = o.Reactor
	toSerialize["engine"] = o.Engine
	toSerialize["modules"] = o.Modules
	toSerialize["mounts"] = o.Mounts
	return toSerialize, nil
}

type NullableShipyardShip struct {
	value *ShipyardShip
	isSet bool
}

func (v NullableShipyardShip) Get() *ShipyardShip {
	return v.value
}

func (v *NullableShipyardShip) Set(val *ShipyardShip) {
	v.value = val
	v.isSet = true
}

func (v NullableShipyardShip) IsSet() bool {
	return v.isSet
}

func (v *NullableShipyardShip) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipyardShip(val *ShipyardShip) *NullableShipyardShip {
	return &NullableShipyardShip{value: val, isSet: true}
}

func (v NullableShipyardShip) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipyardShip) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


