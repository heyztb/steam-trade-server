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

// TradeProposal is an object. TradeProposal defines the information necessary for STS to validate and create a trade offer between the appropriate Trade Bot and a user.
type TradeProposal struct {
	// Offer: items to send
	Offer []Asset `json:"offer" mapstructure:"offer"`
	// Usersteam64: Steam64 ID of the user participating in the trade offer
	Usersteam64 int64 `json:"usersteam64" mapstructure:"usersteam64"`
	// Want: items to receive
	Want []Asset `json:"want" mapstructure:"want"`
}

// Validate implements basic validation for this model
func (m TradeProposal) Validate() error {
	return validation.Errors{
		"offer": validation.Validate(
			m.Offer, validation.NotNil,
		),
		"want": validation.Validate(
			m.Want, validation.NotNil,
		),
	}.Filter()
}

// GetOffer returns the Offer property
func (m TradeProposal) GetOffer() []Asset {
	return m.Offer
}

// SetOffer sets the Offer property
func (m *TradeProposal) SetOffer(val []Asset) {
	m.Offer = val
}

// GetUsersteam64 returns the Usersteam64 property
func (m TradeProposal) GetUsersteam64() int64 {
	return m.Usersteam64
}

// SetUsersteam64 sets the Usersteam64 property
func (m *TradeProposal) SetUsersteam64(val int64) {
	m.Usersteam64 = val
}

// GetWant returns the Want property
func (m TradeProposal) GetWant() []Asset {
	return m.Want
}

// SetWant sets the Want property
func (m *TradeProposal) SetWant(val []Asset) {
	m.Want = val
}
