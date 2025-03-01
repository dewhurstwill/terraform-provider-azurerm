package services

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadAuthFailureMode string

const (
	AadAuthFailureModeHTTPFourZeroOneWithBearerChallenge AadAuthFailureMode = "http401WithBearerChallenge"
	AadAuthFailureModeHTTPFourZeroThree                  AadAuthFailureMode = "http403"
)

func PossibleValuesForAadAuthFailureMode() []string {
	return []string{
		string(AadAuthFailureModeHTTPFourZeroOneWithBearerChallenge),
		string(AadAuthFailureModeHTTPFourZeroThree),
	}
}

func parseAadAuthFailureMode(input string) (*AadAuthFailureMode, error) {
	vals := map[string]AadAuthFailureMode{
		"http401withbearerchallenge": AadAuthFailureModeHTTPFourZeroOneWithBearerChallenge,
		"http403":                    AadAuthFailureModeHTTPFourZeroThree,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AadAuthFailureMode(input)
	return &out, nil
}

type HostingMode string

const (
	HostingModeDefault     HostingMode = "default"
	HostingModeHighDensity HostingMode = "highDensity"
)

func PossibleValuesForHostingMode() []string {
	return []string{
		string(HostingModeDefault),
		string(HostingModeHighDensity),
	}
}

func parseHostingMode(input string) (*HostingMode, error) {
	vals := map[string]HostingMode{
		"default":     HostingModeDefault,
		"highdensity": HostingModeHighDensity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostingMode(input)
	return &out, nil
}

type PrivateLinkServiceConnectionProvisioningState string

const (
	PrivateLinkServiceConnectionProvisioningStateCanceled   PrivateLinkServiceConnectionProvisioningState = "Canceled"
	PrivateLinkServiceConnectionProvisioningStateDeleting   PrivateLinkServiceConnectionProvisioningState = "Deleting"
	PrivateLinkServiceConnectionProvisioningStateFailed     PrivateLinkServiceConnectionProvisioningState = "Failed"
	PrivateLinkServiceConnectionProvisioningStateIncomplete PrivateLinkServiceConnectionProvisioningState = "Incomplete"
	PrivateLinkServiceConnectionProvisioningStateSucceeded  PrivateLinkServiceConnectionProvisioningState = "Succeeded"
	PrivateLinkServiceConnectionProvisioningStateUpdating   PrivateLinkServiceConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateLinkServiceConnectionProvisioningState() []string {
	return []string{
		string(PrivateLinkServiceConnectionProvisioningStateCanceled),
		string(PrivateLinkServiceConnectionProvisioningStateDeleting),
		string(PrivateLinkServiceConnectionProvisioningStateFailed),
		string(PrivateLinkServiceConnectionProvisioningStateIncomplete),
		string(PrivateLinkServiceConnectionProvisioningStateSucceeded),
		string(PrivateLinkServiceConnectionProvisioningStateUpdating),
	}
}

func parsePrivateLinkServiceConnectionProvisioningState(input string) (*PrivateLinkServiceConnectionProvisioningState, error) {
	vals := map[string]PrivateLinkServiceConnectionProvisioningState{
		"canceled":   PrivateLinkServiceConnectionProvisioningStateCanceled,
		"deleting":   PrivateLinkServiceConnectionProvisioningStateDeleting,
		"failed":     PrivateLinkServiceConnectionProvisioningStateFailed,
		"incomplete": PrivateLinkServiceConnectionProvisioningStateIncomplete,
		"succeeded":  PrivateLinkServiceConnectionProvisioningStateSucceeded,
		"updating":   PrivateLinkServiceConnectionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionProvisioningState(input)
	return &out, nil
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}

func parsePrivateLinkServiceConnectionStatus(input string) (*PrivateLinkServiceConnectionStatus, error) {
	vals := map[string]PrivateLinkServiceConnectionStatus{
		"approved":     PrivateLinkServiceConnectionStatusApproved,
		"disconnected": PrivateLinkServiceConnectionStatusDisconnected,
		"pending":      PrivateLinkServiceConnectionStatusPending,
		"rejected":     PrivateLinkServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateFailed       ProvisioningState = "failed"
	ProvisioningStateProvisioning ProvisioningState = "provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

func parsePublicNetworkAccess(input string) (*PublicNetworkAccess, error) {
	vals := map[string]PublicNetworkAccess{
		"disabled": PublicNetworkAccessDisabled,
		"enabled":  PublicNetworkAccessEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PublicNetworkAccess(input)
	return &out, nil
}

type ResourceType string

const (
	ResourceTypeSearchServices ResourceType = "searchServices"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeSearchServices),
	}
}

func parseResourceType(input string) (*ResourceType, error) {
	vals := map[string]ResourceType{
		"searchservices": ResourceTypeSearchServices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceType(input)
	return &out, nil
}

type SearchEncryptionComplianceStatus string

const (
	SearchEncryptionComplianceStatusCompliant    SearchEncryptionComplianceStatus = "Compliant"
	SearchEncryptionComplianceStatusNonCompliant SearchEncryptionComplianceStatus = "NonCompliant"
)

func PossibleValuesForSearchEncryptionComplianceStatus() []string {
	return []string{
		string(SearchEncryptionComplianceStatusCompliant),
		string(SearchEncryptionComplianceStatusNonCompliant),
	}
}

func parseSearchEncryptionComplianceStatus(input string) (*SearchEncryptionComplianceStatus, error) {
	vals := map[string]SearchEncryptionComplianceStatus{
		"compliant":    SearchEncryptionComplianceStatusCompliant,
		"noncompliant": SearchEncryptionComplianceStatusNonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchEncryptionComplianceStatus(input)
	return &out, nil
}

type SearchEncryptionWithCmk string

const (
	SearchEncryptionWithCmkDisabled    SearchEncryptionWithCmk = "Disabled"
	SearchEncryptionWithCmkEnabled     SearchEncryptionWithCmk = "Enabled"
	SearchEncryptionWithCmkUnspecified SearchEncryptionWithCmk = "Unspecified"
)

func PossibleValuesForSearchEncryptionWithCmk() []string {
	return []string{
		string(SearchEncryptionWithCmkDisabled),
		string(SearchEncryptionWithCmkEnabled),
		string(SearchEncryptionWithCmkUnspecified),
	}
}

func parseSearchEncryptionWithCmk(input string) (*SearchEncryptionWithCmk, error) {
	vals := map[string]SearchEncryptionWithCmk{
		"disabled":    SearchEncryptionWithCmkDisabled,
		"enabled":     SearchEncryptionWithCmkEnabled,
		"unspecified": SearchEncryptionWithCmkUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchEncryptionWithCmk(input)
	return &out, nil
}

type SearchServiceStatus string

const (
	SearchServiceStatusDegraded     SearchServiceStatus = "degraded"
	SearchServiceStatusDeleting     SearchServiceStatus = "deleting"
	SearchServiceStatusDisabled     SearchServiceStatus = "disabled"
	SearchServiceStatusError        SearchServiceStatus = "error"
	SearchServiceStatusProvisioning SearchServiceStatus = "provisioning"
	SearchServiceStatusRunning      SearchServiceStatus = "running"
	SearchServiceStatusStopped      SearchServiceStatus = "stopped"
)

func PossibleValuesForSearchServiceStatus() []string {
	return []string{
		string(SearchServiceStatusDegraded),
		string(SearchServiceStatusDeleting),
		string(SearchServiceStatusDisabled),
		string(SearchServiceStatusError),
		string(SearchServiceStatusProvisioning),
		string(SearchServiceStatusRunning),
		string(SearchServiceStatusStopped),
	}
}

func parseSearchServiceStatus(input string) (*SearchServiceStatus, error) {
	vals := map[string]SearchServiceStatus{
		"degraded":     SearchServiceStatusDegraded,
		"deleting":     SearchServiceStatusDeleting,
		"disabled":     SearchServiceStatusDisabled,
		"error":        SearchServiceStatusError,
		"provisioning": SearchServiceStatusProvisioning,
		"running":      SearchServiceStatusRunning,
		"stopped":      SearchServiceStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchServiceStatus(input)
	return &out, nil
}

type SharedPrivateLinkResourceProvisioningState string

const (
	SharedPrivateLinkResourceProvisioningStateDeleting   SharedPrivateLinkResourceProvisioningState = "Deleting"
	SharedPrivateLinkResourceProvisioningStateFailed     SharedPrivateLinkResourceProvisioningState = "Failed"
	SharedPrivateLinkResourceProvisioningStateIncomplete SharedPrivateLinkResourceProvisioningState = "Incomplete"
	SharedPrivateLinkResourceProvisioningStateSucceeded  SharedPrivateLinkResourceProvisioningState = "Succeeded"
	SharedPrivateLinkResourceProvisioningStateUpdating   SharedPrivateLinkResourceProvisioningState = "Updating"
)

func PossibleValuesForSharedPrivateLinkResourceProvisioningState() []string {
	return []string{
		string(SharedPrivateLinkResourceProvisioningStateDeleting),
		string(SharedPrivateLinkResourceProvisioningStateFailed),
		string(SharedPrivateLinkResourceProvisioningStateIncomplete),
		string(SharedPrivateLinkResourceProvisioningStateSucceeded),
		string(SharedPrivateLinkResourceProvisioningStateUpdating),
	}
}

func parseSharedPrivateLinkResourceProvisioningState(input string) (*SharedPrivateLinkResourceProvisioningState, error) {
	vals := map[string]SharedPrivateLinkResourceProvisioningState{
		"deleting":   SharedPrivateLinkResourceProvisioningStateDeleting,
		"failed":     SharedPrivateLinkResourceProvisioningStateFailed,
		"incomplete": SharedPrivateLinkResourceProvisioningStateIncomplete,
		"succeeded":  SharedPrivateLinkResourceProvisioningStateSucceeded,
		"updating":   SharedPrivateLinkResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPrivateLinkResourceProvisioningState(input)
	return &out, nil
}

type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
)

func PossibleValuesForSharedPrivateLinkResourceStatus() []string {
	return []string{
		string(SharedPrivateLinkResourceStatusApproved),
		string(SharedPrivateLinkResourceStatusDisconnected),
		string(SharedPrivateLinkResourceStatusPending),
		string(SharedPrivateLinkResourceStatusRejected),
	}
}

func parseSharedPrivateLinkResourceStatus(input string) (*SharedPrivateLinkResourceStatus, error) {
	vals := map[string]SharedPrivateLinkResourceStatus{
		"approved":     SharedPrivateLinkResourceStatusApproved,
		"disconnected": SharedPrivateLinkResourceStatusDisconnected,
		"pending":      SharedPrivateLinkResourceStatusPending,
		"rejected":     SharedPrivateLinkResourceStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPrivateLinkResourceStatus(input)
	return &out, nil
}

type SkuName string

const (
	SkuNameBasic                SkuName = "basic"
	SkuNameFree                 SkuName = "free"
	SkuNameStandard             SkuName = "standard"
	SkuNameStandardThree        SkuName = "standard3"
	SkuNameStandardTwo          SkuName = "standard2"
	SkuNameStorageOptimizedLOne SkuName = "storage_optimized_l1"
	SkuNameStorageOptimizedLTwo SkuName = "storage_optimized_l2"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameFree),
		string(SkuNameStandard),
		string(SkuNameStandardThree),
		string(SkuNameStandardTwo),
		string(SkuNameStorageOptimizedLOne),
		string(SkuNameStorageOptimizedLTwo),
	}
}

func parseSkuName(input string) (*SkuName, error) {
	vals := map[string]SkuName{
		"basic":                SkuNameBasic,
		"free":                 SkuNameFree,
		"standard":             SkuNameStandard,
		"standard3":            SkuNameStandardThree,
		"standard2":            SkuNameStandardTwo,
		"storage_optimized_l1": SkuNameStorageOptimizedLOne,
		"storage_optimized_l2": SkuNameStorageOptimizedLTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuName(input)
	return &out, nil
}

type UnavailableNameReason string

const (
	UnavailableNameReasonAlreadyExists UnavailableNameReason = "AlreadyExists"
	UnavailableNameReasonInvalid       UnavailableNameReason = "Invalid"
)

func PossibleValuesForUnavailableNameReason() []string {
	return []string{
		string(UnavailableNameReasonAlreadyExists),
		string(UnavailableNameReasonInvalid),
	}
}

func parseUnavailableNameReason(input string) (*UnavailableNameReason, error) {
	vals := map[string]UnavailableNameReason{
		"alreadyexists": UnavailableNameReasonAlreadyExists,
		"invalid":       UnavailableNameReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnavailableNameReason(input)
	return &out, nil
}
