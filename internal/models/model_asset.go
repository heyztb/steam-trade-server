// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Steam Trade Server
//	Version: 0.1.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Asset is an object. Asset defines the information needed to find a specific item on Steam
type Asset struct {
	// Amount:
	Amount int32 `json:"amount" mapstructure:"amount"`
	// AppID: The game the item belongs to
	AppID int32 `json:"appID" mapstructure:"appID"`
	// AssetID:
	AssetID int32 `json:"assetID" mapstructure:"assetID"`
	// ClassID:
	ClassID int32 `json:"classID" mapstructure:"classID"`
	// ContextID:
	ContextID int32 `json:"contextID" mapstructure:"contextID"`
	// CurrencyID:
	CurrencyID int32 `json:"currencyID" mapstructure:"currencyID"`
	// InstanceID:
	InstanceID int32 `json:"instanceID" mapstructure:"instanceID"`
	// Missing:
	Missing bool `json:"missing" mapstructure:"missing"`
}

// Validate implements basic validation for this model
func (m Asset) Validate() error {
	return validation.Errors{}.Filter()
}

// GetAmount returns the Amount property
func (m Asset) GetAmount() int32 {
	return m.Amount
}

// SetAmount sets the Amount property
func (m *Asset) SetAmount(val int32) {
	m.Amount = val
}

// GetAppID returns the AppID property
func (m Asset) GetAppID() int32 {
	return m.AppID
}

// SetAppID sets the AppID property
func (m *Asset) SetAppID(val int32) {
	m.AppID = val
}

// GetAssetID returns the AssetID property
func (m Asset) GetAssetID() int32 {
	return m.AssetID
}

// SetAssetID sets the AssetID property
func (m *Asset) SetAssetID(val int32) {
	m.AssetID = val
}

// GetClassID returns the ClassID property
func (m Asset) GetClassID() int32 {
	return m.ClassID
}

// SetClassID sets the ClassID property
func (m *Asset) SetClassID(val int32) {
	m.ClassID = val
}

// GetContextID returns the ContextID property
func (m Asset) GetContextID() int32 {
	return m.ContextID
}

// SetContextID sets the ContextID property
func (m *Asset) SetContextID(val int32) {
	m.ContextID = val
}

// GetCurrencyID returns the CurrencyID property
func (m Asset) GetCurrencyID() int32 {
	return m.CurrencyID
}

// SetCurrencyID sets the CurrencyID property
func (m *Asset) SetCurrencyID(val int32) {
	m.CurrencyID = val
}

// GetInstanceID returns the InstanceID property
func (m Asset) GetInstanceID() int32 {
	return m.InstanceID
}

// SetInstanceID sets the InstanceID property
func (m *Asset) SetInstanceID(val int32) {
	m.InstanceID = val
}

// GetMissing returns the Missing property
func (m Asset) GetMissing() bool {
	return m.Missing
}

// SetMissing sets the Missing property
func (m *Asset) SetMissing(val bool) {
	m.Missing = val
}
