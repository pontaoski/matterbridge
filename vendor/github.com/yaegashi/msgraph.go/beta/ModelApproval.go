// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ApprovalWorkflowProvider undocumented
type ApprovalWorkflowProvider struct {
	// Entity is the base model of ApprovalWorkflowProvider
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// BusinessFlows undocumented
	BusinessFlows []BusinessFlow `json:"businessFlows,omitempty"`
	// PolicyTemplates undocumented
	PolicyTemplates []GovernancePolicyTemplate `json:"policyTemplates,omitempty"`
	// Requests undocumented
	Requests []RequestObject `json:"requests,omitempty"`
	// RequestsAwaitingMyDecision undocumented
	RequestsAwaitingMyDecision []RequestObject `json:"requestsAwaitingMyDecision,omitempty"`
	// BusinessFlowsWithRequestsAwaitingMyDecision undocumented
	BusinessFlowsWithRequestsAwaitingMyDecision []BusinessFlow `json:"businessFlowsWithRequestsAwaitingMyDecision,omitempty"`
}
