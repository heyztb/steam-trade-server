// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: Steam Trade Server
//	Version: 0.1.0
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// TradeProposal is an object. TradeProposal defines the information necessary for STS to validate and create a trade offer between the appropriate Trade Bot and a user.
type TradeProposal struct {
	// Offer: items to send
	Offer []Asset `json:"offer" mapstructure:"offer"`
	// TradeUrl: Steam Trade Offer URL of the user participating in the trade
	TradeUrl string `json:"trade_url,omitempty" mapstructure:"trade_url,omitempty"`
	// Want: items to receive
	Want []Asset `json:"want" mapstructure:"want"`
}

// Validate implements basic validation for this model
func (m TradeProposal) Validate() error {
	return validation.Errors{
		"offer": validation.Validate(
			m.Offer, validation.NotNil,
		),
		"tradeUrl": validation.Validate(
			m.TradeUrl, is.RequestURI,
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

// GetTradeUrl returns the TradeUrl property
func (m TradeProposal) GetTradeUrl() string {
	return m.TradeUrl
}

// SetTradeUrl sets the TradeUrl property
func (m *TradeProposal) SetTradeUrl(val string) {
	m.TradeUrl = val
}

// GetWant returns the Want property
func (m TradeProposal) GetWant() []Asset {
	return m.Want
}

// SetWant sets the Want property
func (m *TradeProposal) SetWant(val []Asset) {
	m.Want = val
}
