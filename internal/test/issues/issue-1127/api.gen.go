// Package issue1127 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/Mecenate-Inc/oapi-codegen/v2 version v2.0.0-00010101000000-000000000000 DO NOT EDIT.
package issue1127

// Whatever defines model for Whatever.
type Whatever struct {
	SomeProperty *string `json:"someProperty,omitempty"`
}

// CreateWhateverApplicationWildcardPlusJSONRequestBody defines body for CreateWhatever for application/*+json ContentType.
type CreateWhateverApplicationWildcardPlusJSONRequestBody = Whatever

// CreateWhateverJSONRequestBody defines body for CreateWhatever for application/json ContentType.
type CreateWhateverJSONRequestBody = Whatever

// CreateWhateverApplicationJSONPatchPlusJSONRequestBody defines body for CreateWhatever for application/json-patch+json ContentType.
type CreateWhateverApplicationJSONPatchPlusJSONRequestBody = Whatever
