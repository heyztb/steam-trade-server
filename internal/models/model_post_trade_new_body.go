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

// PostTradeNewBody is an object.
type PostTradeNewBody struct {
	// TradeProposal: TradeProposal defines the information necessary for STS to validate and create a trade offer between the appropriate Trade Bot and a user.
	TradeProposal TradeProposal `json:"tradeProposal" mapstructure:"tradeProposal"`
}

// Validate implements basic validation for this model
func (m PostTradeNewBody) Validate() error {
	return validation.Errors{
		"tradeProposal": validation.Validate(
			m.TradeProposal, validation.NotNil,
		),
	}.Filter()
}

// GetTradeProposal returns the TradeProposal property
func (m PostTradeNewBody) GetTradeProposal() TradeProposal {
	return m.TradeProposal
}

// SetTradeProposal sets the TradeProposal property
func (m *PostTradeNewBody) SetTradeProposal(val TradeProposal) {
	m.TradeProposal = val
}
