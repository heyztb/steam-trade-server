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

// Error is an object. Error contains information about what went wrong with a request
type Error struct {
	// Code: Error code
	Code int32 `json:"code" mapstructure:"code"`
	// Message: Error message
	Message string `json:"message" mapstructure:"message"`
}

// Validate implements basic validation for this model
func (m Error) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCode returns the Code property
func (m Error) GetCode() int32 {
	return m.Code
}

// SetCode sets the Code property
func (m *Error) SetCode(val int32) {
	m.Code = val
}

// GetMessage returns the Message property
func (m Error) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *Error) SetMessage(val string) {
	m.Message = val
}
