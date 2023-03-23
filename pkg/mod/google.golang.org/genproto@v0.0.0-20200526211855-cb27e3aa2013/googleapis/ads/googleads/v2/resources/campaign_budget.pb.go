// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.2
// source: google/ads/googleads/v2/resources/campaign_budget.proto

package resources

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v2/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A campaign budget.
type CampaignBudget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. The resource name of the campaign budget.
	// Campaign budget resource names have the form:
	//
	// `customers/{customer_id}/campaignBudgets/{budget_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the campaign budget.
	//
	// A campaign budget is created using the CampaignBudgetService create
	// operation and is assigned a budget ID. A budget ID can be shared across
	// different campaigns; the system will then allocate the campaign budget
	// among different campaigns to get optimum results.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the campaign budget.
	//
	// When creating a campaign budget through CampaignBudgetService, every
	// explicitly shared campaign budget must have a non-null, non-empty name.
	// Campaign budgets that are not explicitly shared derive their name from the
	// attached campaign's name.
	//
	// The length of this string must be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit. Monthly spend is capped at 30.4 times this amount.
	AmountMicros *wrappers.Int64Value `protobuf:"bytes,5,opt,name=amount_micros,json=amountMicros,proto3" json:"amount_micros,omitempty"`
	// The lifetime amount of the budget, in the local currency for the account.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	TotalAmountMicros *wrappers.Int64Value `protobuf:"bytes,10,opt,name=total_amount_micros,json=totalAmountMicros,proto3" json:"total_amount_micros,omitempty"`
	// Output only. The status of this campaign budget. This field is read-only.
	Status enums.BudgetStatusEnum_BudgetStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v2.enums.BudgetStatusEnum_BudgetStatus" json:"status,omitempty"`
	// The delivery method that determines the rate at which the campaign budget
	// is spent.
	//
	// Defaults to STANDARD if unspecified in a create operation.
	DeliveryMethod enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod `protobuf:"varint,7,opt,name=delivery_method,json=deliveryMethod,proto3,enum=google.ads.googleads.v2.enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod" json:"delivery_method,omitempty"`
	// Specifies whether the budget is explicitly shared. Defaults to true if
	// unspecified in a create operation.
	//
	// If true, the budget was created with the purpose of sharing
	// across one or more campaigns.
	//
	// If false, the budget was created with the intention of only being used
	// with a single campaign. The budget's name and status will stay in sync
	// with the campaign's name and status. Attempting to share the budget with a
	// second campaign will result in an error.
	//
	// A non-shared budget can become an explicitly shared. The same operation
	// must also assign the budget a name.
	//
	// A shared campaign budget can never become non-shared.
	ExplicitlyShared *wrappers.BoolValue `protobuf:"bytes,8,opt,name=explicitly_shared,json=explicitlyShared,proto3" json:"explicitly_shared,omitempty"`
	// Output only. The number of campaigns actively using the budget.
	//
	// This field is read-only.
	ReferenceCount *wrappers.Int64Value `protobuf:"bytes,9,opt,name=reference_count,json=referenceCount,proto3" json:"reference_count,omitempty"`
	// Output only. Indicates whether there is a recommended budget for this campaign budget.
	//
	// This field is read-only.
	HasRecommendedBudget *wrappers.BoolValue `protobuf:"bytes,11,opt,name=has_recommended_budget,json=hasRecommendedBudget,proto3" json:"has_recommended_budget,omitempty"`
	// Output only. The recommended budget amount. If no recommendation is available, this will
	// be set to the budget amount.
	// Amount is specified in micros, where one million is equivalent to one
	// currency unit.
	//
	// This field is read-only.
	RecommendedBudgetAmountMicros *wrappers.Int64Value `protobuf:"bytes,12,opt,name=recommended_budget_amount_micros,json=recommendedBudgetAmountMicros,proto3" json:"recommended_budget_amount_micros,omitempty"`
	// Immutable. Period over which to spend the budget. Defaults to DAILY if not specified.
	Period enums.BudgetPeriodEnum_BudgetPeriod `protobuf:"varint,13,opt,name=period,proto3,enum=google.ads.googleads.v2.enums.BudgetPeriodEnum_BudgetPeriod" json:"period,omitempty"`
	// Output only. The estimated change in weekly clicks if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyClicks *wrappers.Int64Value `protobuf:"bytes,14,opt,name=recommended_budget_estimated_change_weekly_clicks,json=recommendedBudgetEstimatedChangeWeeklyClicks,proto3" json:"recommended_budget_estimated_change_weekly_clicks,omitempty"`
	// Output only. The estimated change in weekly cost in micros if the recommended budget is
	// applied. One million is equivalent to one currency unit.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyCostMicros *wrappers.Int64Value `protobuf:"bytes,15,opt,name=recommended_budget_estimated_change_weekly_cost_micros,json=recommendedBudgetEstimatedChangeWeeklyCostMicros,proto3" json:"recommended_budget_estimated_change_weekly_cost_micros,omitempty"`
	// Output only. The estimated change in weekly interactions if the recommended budget is
	// applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyInteractions *wrappers.Int64Value `protobuf:"bytes,16,opt,name=recommended_budget_estimated_change_weekly_interactions,json=recommendedBudgetEstimatedChangeWeeklyInteractions,proto3" json:"recommended_budget_estimated_change_weekly_interactions,omitempty"`
	// Output only. The estimated change in weekly views if the recommended budget is applied.
	//
	// This field is read-only.
	RecommendedBudgetEstimatedChangeWeeklyViews *wrappers.Int64Value `protobuf:"bytes,17,opt,name=recommended_budget_estimated_change_weekly_views,json=recommendedBudgetEstimatedChangeWeeklyViews,proto3" json:"recommended_budget_estimated_change_weekly_views,omitempty"`
	// Immutable. The type of the campaign budget.
	Type enums.BudgetTypeEnum_BudgetType `protobuf:"varint,18,opt,name=type,proto3,enum=google.ads.googleads.v2.enums.BudgetTypeEnum_BudgetType" json:"type,omitempty"`
}

func (x *CampaignBudget) Reset() {
	*x = CampaignBudget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_ads_googleads_v2_resources_campaign_budget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CampaignBudget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignBudget) ProtoMessage() {}

func (x *CampaignBudget) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v2_resources_campaign_budget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignBudget.ProtoReflect.Descriptor instead.
func (*CampaignBudget) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignBudget) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignBudget) GetId() *wrappers.Int64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CampaignBudget) GetName() *wrappers.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *CampaignBudget) GetAmountMicros() *wrappers.Int64Value {
	if x != nil {
		return x.AmountMicros
	}
	return nil
}

func (x *CampaignBudget) GetTotalAmountMicros() *wrappers.Int64Value {
	if x != nil {
		return x.TotalAmountMicros
	}
	return nil
}

func (x *CampaignBudget) GetStatus() enums.BudgetStatusEnum_BudgetStatus {
	if x != nil {
		return x.Status
	}
	return enums.BudgetStatusEnum_UNSPECIFIED
}

func (x *CampaignBudget) GetDeliveryMethod() enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod {
	if x != nil {
		return x.DeliveryMethod
	}
	return enums.BudgetDeliveryMethodEnum_UNSPECIFIED
}

func (x *CampaignBudget) GetExplicitlyShared() *wrappers.BoolValue {
	if x != nil {
		return x.ExplicitlyShared
	}
	return nil
}

func (x *CampaignBudget) GetReferenceCount() *wrappers.Int64Value {
	if x != nil {
		return x.ReferenceCount
	}
	return nil
}

func (x *CampaignBudget) GetHasRecommendedBudget() *wrappers.BoolValue {
	if x != nil {
		return x.HasRecommendedBudget
	}
	return nil
}

func (x *CampaignBudget) GetRecommendedBudgetAmountMicros() *wrappers.Int64Value {
	if x != nil {
		return x.RecommendedBudgetAmountMicros
	}
	return nil
}

func (x *CampaignBudget) GetPeriod() enums.BudgetPeriodEnum_BudgetPeriod {
	if x != nil {
		return x.Period
	}
	return enums.BudgetPeriodEnum_UNSPECIFIED
}

func (x *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyClicks() *wrappers.Int64Value {
	if x != nil {
		return x.RecommendedBudgetEstimatedChangeWeeklyClicks
	}
	return nil
}

func (x *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyCostMicros() *wrappers.Int64Value {
	if x != nil {
		return x.RecommendedBudgetEstimatedChangeWeeklyCostMicros
	}
	return nil
}

func (x *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyInteractions() *wrappers.Int64Value {
	if x != nil {
		return x.RecommendedBudgetEstimatedChangeWeeklyInteractions
	}
	return nil
}

func (x *CampaignBudget) GetRecommendedBudgetEstimatedChangeWeeklyViews() *wrappers.Int64Value {
	if x != nil {
		return x.RecommendedBudgetEstimatedChangeWeeklyViews
	}
	return nil
}

func (x *CampaignBudget) GetType() enums.BudgetTypeEnum_BudgetType {
	if x != nil {
		return x.Type
	}
	return enums.BudgetTypeEnum_UNSPECIFIED
}

var File_google_ads_googleads_v2_resources_campaign_budget_proto protoreflect.FileDescriptor

var file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x1a, 0x3a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76,
	0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70,
	0x65, 0x72, 0x69, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x62, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x0d, 0x0a, 0x0e, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x0d,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x05, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x4b, 0x0a, 0x13, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x59, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61,
	0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x75, 0x0a, 0x0f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x4c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x45, 0x6e,
	0x75, 0x6d, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x47, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x6c, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10,
	0x65, 0x78, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x6c, 0x79, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x12, 0x49, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x55, 0x0a, 0x16, 0x68,
	0x61, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x62,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x68, 0x61,
	0x73, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x69, 0x0a, 0x20, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x1d,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x59, 0x0a,
	0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x64, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42,
	0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x42, 0x03, 0xe0, 0x41, 0x05,
	0x52, 0x06, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x89, 0x01, 0x0a, 0x31, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f,
	0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x2c, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x43, 0x6c,
	0x69, 0x63, 0x6b, 0x73, 0x12, 0x92, 0x01, 0x0a, 0x36, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x73, 0x74, 0x69,
	0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x77, 0x65, 0x65,
	0x6b, 0x6c, 0x79, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x5f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x30, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x43,
	0x6f, 0x73, 0x74, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x12, 0x95, 0x01, 0x0a, 0x37, 0x72, 0x65,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x32, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x57,
	0x65, 0x65, 0x6b, 0x6c, 0x79, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x87, 0x01, 0x0a, 0x30, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65,
	0x64, 0x5f, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79,
	0x5f, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x2b,
	0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x56, 0x69, 0x65, 0x77, 0x73, 0x12, 0x51, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x64,
	0xea, 0x41, 0x61, 0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x61,
	0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x36, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x7d, 0x2f, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65,
	0x74, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x5f, 0x62, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x7d, 0x42, 0x80, 0x02, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x64, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x42, 0x13,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x61, 0x64, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0xa2, 0x02, 0x03, 0x47, 0x41, 0x41, 0xaa, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x41, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x2e, 0x56,
	0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xca, 0x02, 0x21, 0x47, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x41, 0x64, 0x73, 0x5c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41,
	0x64, 0x73, 0x5c, 0x56, 0x32, 0x5c, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0xea,
	0x02, 0x25, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x47,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x41, 0x64, 0x73, 0x3a, 0x3a, 0x56, 0x32, 0x3a, 0x3a, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescData = file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDesc
)

func file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescData)
	})
	return file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDescData
}

var file_google_ads_googleads_v2_resources_campaign_budget_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v2_resources_campaign_budget_proto_goTypes = []interface{}{
	(*CampaignBudget)(nil),                                   // 0: google.ads.googleads.v2.resources.CampaignBudget
	(*wrappers.Int64Value)(nil),                              // 1: google.protobuf.Int64Value
	(*wrappers.StringValue)(nil),                             // 2: google.protobuf.StringValue
	(enums.BudgetStatusEnum_BudgetStatus)(0),                 // 3: google.ads.googleads.v2.enums.BudgetStatusEnum.BudgetStatus
	(enums.BudgetDeliveryMethodEnum_BudgetDeliveryMethod)(0), // 4: google.ads.googleads.v2.enums.BudgetDeliveryMethodEnum.BudgetDeliveryMethod
	(*wrappers.BoolValue)(nil),                               // 5: google.protobuf.BoolValue
	(enums.BudgetPeriodEnum_BudgetPeriod)(0),                 // 6: google.ads.googleads.v2.enums.BudgetPeriodEnum.BudgetPeriod
	(enums.BudgetTypeEnum_BudgetType)(0),                     // 7: google.ads.googleads.v2.enums.BudgetTypeEnum.BudgetType
}
var file_google_ads_googleads_v2_resources_campaign_budget_proto_depIdxs = []int32{
	1,  // 0: google.ads.googleads.v2.resources.CampaignBudget.id:type_name -> google.protobuf.Int64Value
	2,  // 1: google.ads.googleads.v2.resources.CampaignBudget.name:type_name -> google.protobuf.StringValue
	1,  // 2: google.ads.googleads.v2.resources.CampaignBudget.amount_micros:type_name -> google.protobuf.Int64Value
	1,  // 3: google.ads.googleads.v2.resources.CampaignBudget.total_amount_micros:type_name -> google.protobuf.Int64Value
	3,  // 4: google.ads.googleads.v2.resources.CampaignBudget.status:type_name -> google.ads.googleads.v2.enums.BudgetStatusEnum.BudgetStatus
	4,  // 5: google.ads.googleads.v2.resources.CampaignBudget.delivery_method:type_name -> google.ads.googleads.v2.enums.BudgetDeliveryMethodEnum.BudgetDeliveryMethod
	5,  // 6: google.ads.googleads.v2.resources.CampaignBudget.explicitly_shared:type_name -> google.protobuf.BoolValue
	1,  // 7: google.ads.googleads.v2.resources.CampaignBudget.reference_count:type_name -> google.protobuf.Int64Value
	5,  // 8: google.ads.googleads.v2.resources.CampaignBudget.has_recommended_budget:type_name -> google.protobuf.BoolValue
	1,  // 9: google.ads.googleads.v2.resources.CampaignBudget.recommended_budget_amount_micros:type_name -> google.protobuf.Int64Value
	6,  // 10: google.ads.googleads.v2.resources.CampaignBudget.period:type_name -> google.ads.googleads.v2.enums.BudgetPeriodEnum.BudgetPeriod
	1,  // 11: google.ads.googleads.v2.resources.CampaignBudget.recommended_budget_estimated_change_weekly_clicks:type_name -> google.protobuf.Int64Value
	1,  // 12: google.ads.googleads.v2.resources.CampaignBudget.recommended_budget_estimated_change_weekly_cost_micros:type_name -> google.protobuf.Int64Value
	1,  // 13: google.ads.googleads.v2.resources.CampaignBudget.recommended_budget_estimated_change_weekly_interactions:type_name -> google.protobuf.Int64Value
	1,  // 14: google.ads.googleads.v2.resources.CampaignBudget.recommended_budget_estimated_change_weekly_views:type_name -> google.protobuf.Int64Value
	7,  // 15: google.ads.googleads.v2.resources.CampaignBudget.type:type_name -> google.ads.googleads.v2.enums.BudgetTypeEnum.BudgetType
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v2_resources_campaign_budget_proto_init() }
func file_google_ads_googleads_v2_resources_campaign_budget_proto_init() {
	if File_google_ads_googleads_v2_resources_campaign_budget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_ads_googleads_v2_resources_campaign_budget_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CampaignBudget); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v2_resources_campaign_budget_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v2_resources_campaign_budget_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v2_resources_campaign_budget_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v2_resources_campaign_budget_proto = out.File
	file_google_ads_googleads_v2_resources_campaign_budget_proto_rawDesc = nil
	file_google_ads_googleads_v2_resources_campaign_budget_proto_goTypes = nil
	file_google_ads_googleads_v2_resources_campaign_budget_proto_depIdxs = nil
}
