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

type UpdateVolumeOption struct {
	// 新的云硬盘的描述，name和description不能同时为null。最大支持255个字节。
	Description *string `json:"description,omitempty"`
	// 新的云硬盘的名字，name和description不能同时为null。最大支持255个字节。
	Name *string `json:"name,omitempty"`
}

func (o UpdateVolumeOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateVolumeOption struct{}"
	}

	return strings.Join([]string{"UpdateVolumeOption", string(data)}, " ")
}
