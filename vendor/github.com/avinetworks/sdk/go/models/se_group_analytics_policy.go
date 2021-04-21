package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// SeGroupAnalyticsPolicy se group analytics policy
// swagger:model SeGroupAnalyticsPolicy
type SeGroupAnalyticsPolicy struct {

	// Thresholds for various events generated by metrics system. Field introduced in 20.1.3.
	MetricsEventThresholds []*MetricsEventThreshold `json:"metrics_event_thresholds,omitempty"`
}