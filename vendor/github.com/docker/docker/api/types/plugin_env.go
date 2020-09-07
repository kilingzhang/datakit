package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PluginEnv plugin env
// swagger:model PluginEnv
type PluginEnv struct {

	// description
	// Required: true
	Description string `json:"Description"`

	// name
	// Required: true
	Name string `json:"Name"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// value
	// Required: true
	Value *string `json:"Value"`
}
