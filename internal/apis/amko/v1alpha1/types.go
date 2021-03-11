/*
 * Copyright 2019-2020 VMware, Inc.
 * All Rights Reserved.
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*   http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// GSLBConfig is the top-level type
type GSLBConfig struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// spec for GSLB Config
	Spec GSLBConfigSpec `json:"spec,omitempty"`
	// +optional
	Status GSLBConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GSLBConfigList is a list of GSLBConfig resources
type GSLBConfigList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GSLBConfig `json:"items"`
}

// GSLBConfigSpec is the GSLB configuration
type GSLBConfigSpec struct {
	GSLBLeader      GSLBLeader      `json:"gslbLeader,omitempty"`
	MemberClusters  []MemberCluster `json:"memberClusters,omitempty"`
	RefreshInterval int             `json:"refreshInterval,omitempty"`
	LogLevel        string          `json:"logLevel,omitempty"`
}

// GSLBLeader is the leader node in the GSLB cluster
type GSLBLeader struct {
	Credentials       string `json:"credentials,omitempty"`
	ControllerVersion string `json:"controllerVersion,omitempty"`
	ControllerIP      string `json:"controllerIP,omitempty"`
}

// MemberCluster defines a GSLB member cluster details
type MemberCluster struct {
	ClusterContext string `json:"clusterContext,omitempty"`
}

// GSLBConfigStatus represents the state and status message of the GSLB cluster
type GSLBConfigStatus struct {
	State string `json:"state,omitempty"`
}

// how the Global services are going to be named
const (
	GSNameType = "HOSTNAME"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type GSLBConfigSpecList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GSLBConfigSpec `json:"items"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// GlobalDeploymentPolicy is the top-level type: Global Deployment Policy
// encloses all the rules, actions and configuration required for deploying
// applications.
type GlobalDeploymentPolicy struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// spec for GSLB Config
	Spec GDPSpec `json:"spec,omitempty"`
	// +optional
	Status GDPStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GlobalDeploymentPolicyList is a list of GDP resources
type GlobalDeploymentPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalDeploymentPolicy `json:"items"`
}

// GDPSpec encloses all the properties of a GDP object.
type GDPSpec struct {
	MatchRules    MatchRules         `json:"matchRules,omitempty"`
	MatchClusters []string           `json:"matchClusters,omitempty"`
	TrafficSplit  []TrafficSplitElem `json:"trafficSplit,omitempty"`
}

// MatchRules is the match criteria needed to select the kubernetes/openshift objects.
type MatchRules struct {
	AppSelector       `json:"appSelector,omitempty"`
	NamespaceSelector `json:"namespaceSelector,omitempty"`
}

// AppSelector selects the applications based on their labels
type AppSelector struct {
	Label map[string]string `json:"label,omitempty"`
}

// NamespaceSelector selects the applications based on their labels
type NamespaceSelector struct {
	Label map[string]string `json:"label,omitempty"`
}

// Objects on which rules will be applied
const (
	// RouteObj only applies to openshift Routes
	RouteObj = "ROUTE"
	// IngressObj applies to K8S Ingresses
	IngressObj = "INGRESS"
	// LBSvc applies to service type LoadBalancer
	LBSvcObj = "LBSVC"
	// NSObj applies to namespaces
	NSObj = "Namespace"
)

// TrafficSplitElem determines how much traffic to be routed to a cluster.
type TrafficSplitElem struct {
	// Cluster is the cluster context
	Cluster string `json:"cluster,omitempty"`
	Weight  uint32 `json:"weight,omitempty"`
}

// GDPStatus gives the current status of the policy object.
type GDPStatus struct {
	ErrorStatus string `json:"errorStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// GSLBHostRule is the top-level type which allows a user to override certain
// fields of a GSLB Service.
type GSLBHostRule struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// spec for GSLB Config
	Spec GSLBHostRuleSpec `json:"spec,omitempty"`
	// +optional
	Status GSLBHostRuleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GSLBHostRuleList is a list of GSLBHostRule resources
type GSLBHostRuleList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GSLBHostRule `json:"items"`
}

// GSLBHostRuleSpec defines all the properties of a GSLB Service that can be overriden
// by a user.
type GSLBHostRuleSpec struct {
	// Fqdn is the fqdn of the GSLB Service for which the below properties can be
	// changed.
	Fqdn string `json:"fqdn,omitempty"`
	// TTL is Time To Live in seconds. This tells a DNS resolver how long to hold this DNS
	// record.
	TTL int `json:"ttl,omitempty"`
	// SitePersistenceEnabled if set to true, enables stickiness to the same site where
	// the connection from the client was initiated to.
	SitePersistence SitePersistence `json:"sitePersistence,omitempty"`
	// ThirdPartyMembers is a list of third party members site
	ThirdPartyMembers []ThirdPartyMember `json:"thirdPartyMembers,omitempty"`
	// HealthMonitoreRefs is a list of custom health monitors which will monitor the
	// GSLB Service's pool members.
	HealthMonitorRefs []string `json:"healthMonitorRefs,omitempty"`
	// TrafficSplit defines the weightage of traffic that can be routed to each cluster.
	TrafficSplit []TrafficSplitElem `json:"trafficSplit,omitempty"`
}

// GSLBHostRuleStatus contains the current state of the GSLBHostRule resource. If the
// current state is rejected, then an error message is also shown in the Error field.
type GSLBHostRuleStatus struct {
	Error  string `json:"error,omitempty"`
	Status string `json:"status,omitempty"`
}

// SitePersistence
type SitePersistence struct {
	Enabled    bool   `json:"enabled,omitempty"`
	ProfileRef string `json:"profileRef,omitempty"`
}

type ThirdPartyMember struct {
	VIP  string `json:"vip,omitempty"`
	Site string `json:"site,omitempty"`
}
