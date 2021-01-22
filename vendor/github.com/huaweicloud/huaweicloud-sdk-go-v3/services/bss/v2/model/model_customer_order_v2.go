/*
 * BSS
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type CustomerOrderV2 struct {
	// |参数名称：订单ID。| |参数约束及描述：订单ID。|
	OrderId *string `json:"order_id,omitempty"`
	// |参数名称：客户ID。| |参数约束及描述：客户ID。|
	CustomerId *string `json:"customer_id,omitempty"`
	// |参数名称：云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。具体请参见云服务类型云服务类型云服务类型云服务类型。| |参数约束及描述：云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。具体请参见云服务类型云服务类型云服务类型云服务类型。|
	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// |参数名称：客户订单订单来源类型：1：客户2：代理3：合同4：分销商| |参数的约束及描述：客户订单订单来源类型：1：客户2：代理3：合同4：分销商|
	SourceType *int32 `json:"source_type,omitempty"`
	// |参数名称：订单状态：1：待审核3：处理中4：已取消5：已完成6：待支付9：待确认| |参数的约束及描述：订单状态：1：待审核3：处理中4：已取消5：已完成6：待支付9：待确认|
	Status *int32 `json:"status,omitempty"`
	// |参数名称：订单类型：1：开通2：续订3：变更4：退订10：包周期转按需11：按需转包周期12：赠送13：试用14：转商用15：费用调整| |参数的约束及描述：订单类型：1：开通2：续订3：变更4：退订10：包周期转按需11：按需转包周期12：赠送13：试用14：转商用15：费用调整|
	OrderType *int32 `json:"order_type,omitempty"`
	// |参数名称：订单优惠后金额（不含券不含卡的实付价格）。| |参数的约束及描述：订单优惠后金额（不含券不含卡的实付价格）。|
	AmountAfterDiscount *float64 `json:"amount_after_discount,omitempty"`
	// |参数名称：订单金额（官网价）。退订订单中，该金额等于amount。| |参数的约束及描述：订单金额（官网价）。退订订单中，该金额等于amount。|
	OfficialAmount *float64 `json:"official_amount,omitempty"`
	// |参数名称：订单金额度量单位：1：元2：角3：分| |参数的约束及描述：订单金额度量单位：1：元2：角3：分|
	MeasureId *int32 `json:"measure_id,omitempty"`
	// |参数名称：创建时间 。UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。| |参数约束及描述：创建时间 。UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。|
	CreateTime *string `json:"create_time,omitempty"`
	// |参数名称：支付时间。UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。| |参数约束及描述：支付时间。UTC时间，格式：yyyy-MM-ddTHH:mm:ssZ，如“2019-05-06T08:05:01Z”。其中，HH范围是0～23，mm和ss范围是0～59。|
	PaymentTime *string `json:"payment_time,omitempty"`
	// |参数名称：货币编码。| |参数约束及描述：货币编码。如CNY|
	Currency *string `json:"currency,omitempty"`
	// |参数名称：合同ID。| |参数约束及描述：合同ID。|
	ContractId *string             `json:"contract_id,omitempty"`
	AmountInfo *AmountInfomationV2 `json:"amount_info,omitempty"`
}

func (o CustomerOrderV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CustomerOrderV2 struct{}"
	}

	return strings.Join([]string{"CustomerOrderV2", string(data)}, " ")
}
