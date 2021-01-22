/*
 * EVS
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CinderListVolumeTypesResponse struct {
	VolumeTypes    *[]VolumeType `json:"volume_types,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CinderListVolumeTypesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CinderListVolumeTypesResponse struct{}"
	}

	return strings.Join([]string{"CinderListVolumeTypesResponse", string(data)}, " ")
}
