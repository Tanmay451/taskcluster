// This source code file is AUTO-GENERATED by github.com/taskcluster/jsonschema2go

package tchooksevents

import (
	"encoding/json"
	"errors"
)

type (
	// Message reporting that a hook has changed
	//
	// Defined properties:
	//
	//  struct {
	//
	//  	// `hookGroupId` of the hook that was changed
	//  	//
	//  	// See https://taskcluster-staging.net/schemas/hooks/v1/pulse-hook-changed-message.json#/properties/hookGroupId
	//  	HookGroupID string `json:"hookGroupId"`
	//
	//  	// `hookId` of the hook that was changed
	//  	//
	//  	// See https://taskcluster-staging.net/schemas/hooks/v1/pulse-hook-changed-message.json#/properties/hookId
	//  	HookID string `json:"hookId"`
	//  }
	//
	// Additional properties allowed
	//
	// See https://taskcluster-staging.net/schemas/hooks/v1/pulse-hook-changed-message.json#
	HookChangedMessage json.RawMessage
)

// MarshalJSON calls json.RawMessage method of the same name. Required since
// HookChangedMessage is of type json.RawMessage...
func (this *HookChangedMessage) MarshalJSON() ([]byte, error) {
	x := json.RawMessage(*this)
	return (&x).MarshalJSON()
}

// UnmarshalJSON is a copy of the json.RawMessage implementation.
func (this *HookChangedMessage) UnmarshalJSON(data []byte) error {
	if this == nil {
		return errors.New("HookChangedMessage: UnmarshalJSON on nil pointer")
	}
	*this = append((*this)[0:0], data...)
	return nil
}
