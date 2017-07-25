// Code generated by cdpgen. DO NOT EDIT.

package serviceworker

// UnregisterArgs represents the arguments for Unregister in the ServiceWorker domain.
type UnregisterArgs struct {
	ScopeURL string `json:"scopeURL"` //
}

// NewUnregisterArgs initializes UnregisterArgs with the required arguments.
func NewUnregisterArgs(scopeURL string) *UnregisterArgs {
	args := new(UnregisterArgs)
	args.ScopeURL = scopeURL
	return args
}

// UpdateRegistrationArgs represents the arguments for UpdateRegistration in the ServiceWorker domain.
type UpdateRegistrationArgs struct {
	ScopeURL string `json:"scopeURL"` //
}

// NewUpdateRegistrationArgs initializes UpdateRegistrationArgs with the required arguments.
func NewUpdateRegistrationArgs(scopeURL string) *UpdateRegistrationArgs {
	args := new(UpdateRegistrationArgs)
	args.ScopeURL = scopeURL
	return args
}

// StartWorkerArgs represents the arguments for StartWorker in the ServiceWorker domain.
type StartWorkerArgs struct {
	ScopeURL string `json:"scopeURL"` //
}

// NewStartWorkerArgs initializes StartWorkerArgs with the required arguments.
func NewStartWorkerArgs(scopeURL string) *StartWorkerArgs {
	args := new(StartWorkerArgs)
	args.ScopeURL = scopeURL
	return args
}

// SkipWaitingArgs represents the arguments for SkipWaiting in the ServiceWorker domain.
type SkipWaitingArgs struct {
	ScopeURL string `json:"scopeURL"` //
}

// NewSkipWaitingArgs initializes SkipWaitingArgs with the required arguments.
func NewSkipWaitingArgs(scopeURL string) *SkipWaitingArgs {
	args := new(SkipWaitingArgs)
	args.ScopeURL = scopeURL
	return args
}

// StopWorkerArgs represents the arguments for StopWorker in the ServiceWorker domain.
type StopWorkerArgs struct {
	VersionID string `json:"versionId"` //
}

// NewStopWorkerArgs initializes StopWorkerArgs with the required arguments.
func NewStopWorkerArgs(versionID string) *StopWorkerArgs {
	args := new(StopWorkerArgs)
	args.VersionID = versionID
	return args
}

// InspectWorkerArgs represents the arguments for InspectWorker in the ServiceWorker domain.
type InspectWorkerArgs struct {
	VersionID string `json:"versionId"` //
}

// NewInspectWorkerArgs initializes InspectWorkerArgs with the required arguments.
func NewInspectWorkerArgs(versionID string) *InspectWorkerArgs {
	args := new(InspectWorkerArgs)
	args.VersionID = versionID
	return args
}

// SetForceUpdateOnPageLoadArgs represents the arguments for SetForceUpdateOnPageLoad in the ServiceWorker domain.
type SetForceUpdateOnPageLoadArgs struct {
	ForceUpdateOnPageLoad bool `json:"forceUpdateOnPageLoad"` //
}

// NewSetForceUpdateOnPageLoadArgs initializes SetForceUpdateOnPageLoadArgs with the required arguments.
func NewSetForceUpdateOnPageLoadArgs(forceUpdateOnPageLoad bool) *SetForceUpdateOnPageLoadArgs {
	args := new(SetForceUpdateOnPageLoadArgs)
	args.ForceUpdateOnPageLoad = forceUpdateOnPageLoad
	return args
}

// DeliverPushMessageArgs represents the arguments for DeliverPushMessage in the ServiceWorker domain.
type DeliverPushMessageArgs struct {
	Origin         string `json:"origin"`         //
	RegistrationID string `json:"registrationId"` //
	Data           string `json:"data"`           //
}

// NewDeliverPushMessageArgs initializes DeliverPushMessageArgs with the required arguments.
func NewDeliverPushMessageArgs(origin string, registrationID string, data string) *DeliverPushMessageArgs {
	args := new(DeliverPushMessageArgs)
	args.Origin = origin
	args.RegistrationID = registrationID
	args.Data = data
	return args
}

// DispatchSyncEventArgs represents the arguments for DispatchSyncEvent in the ServiceWorker domain.
type DispatchSyncEventArgs struct {
	Origin         string `json:"origin"`         //
	RegistrationID string `json:"registrationId"` //
	Tag            string `json:"tag"`            //
	LastChance     bool   `json:"lastChance"`     //
}

// NewDispatchSyncEventArgs initializes DispatchSyncEventArgs with the required arguments.
func NewDispatchSyncEventArgs(origin string, registrationID string, tag string, lastChance bool) *DispatchSyncEventArgs {
	args := new(DispatchSyncEventArgs)
	args.Origin = origin
	args.RegistrationID = registrationID
	args.Tag = tag
	args.LastChance = lastChance
	return args
}