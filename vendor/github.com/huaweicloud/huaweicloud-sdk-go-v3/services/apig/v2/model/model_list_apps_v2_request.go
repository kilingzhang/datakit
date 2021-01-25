/*
 * APIG
 *
 * API网关（API Gateway）是为开发者、合作伙伴提供的高性能、高可用、高安全的API托管服务，帮助用户轻松构建、管理和发布任意规模的API。
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAppsV2Request struct {
	ProjectId     string  `json:"project_id"`
	InstanceId    string  `json:"instance_id"`
	Id            *string `json:"id,omitempty"`
	Name          *string `json:"name,omitempty"`
	Status        *int32  `json:"status,omitempty"`
	AppKey        *string `json:"app_key,omitempty"`
	Creator       *string `json:"creator,omitempty"`
	Offset        *int64  `json:"offset,omitempty"`
	Limit         *int32  `json:"limit,omitempty"`
	PreciseSearch *string `json:"precise_search,omitempty"`
}

func (o ListAppsV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAppsV2Request struct{}"
	}

	return strings.Join([]string{"ListAppsV2Request", string(data)}, " ")
}
