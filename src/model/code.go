package model

// Code 邀请码
type Code struct {
	Code string `json:"code,omitempty"`
	Root int    `json:"root,omitempty"`
}
