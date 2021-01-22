/*
 * IAM
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPermanentAccessKeyRequest struct {
	AccessKey string `json:"access_key"`
}

func (o ShowPermanentAccessKeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPermanentAccessKeyRequest struct{}"
	}

	return strings.Join([]string{"ShowPermanentAccessKeyRequest", string(data)}, " ")
}
