/*
 * ECS
 *
 * ECS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type NovaCreateKeypairRequest struct {
	OpenStackAPIVersion *string                       `json:"OpenStack-API-Version,omitempty"`
	Body                *NovaCreateKeypairRequestBody `json:"body,omitempty"`
}

func (o NovaCreateKeypairRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NovaCreateKeypairRequest struct{}"
	}

	return strings.Join([]string{"NovaCreateKeypairRequest", string(data)}, " ")
}
