// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// OperationResult undocumented
type OperationResult string

const (
	// OperationResultVSuccess undocumented
	OperationResultVSuccess OperationResult = "success"
	// OperationResultVFailure undocumented
	OperationResultVFailure OperationResult = "failure"
	// OperationResultVTimeout undocumented
	OperationResultVTimeout OperationResult = "timeout"
	// OperationResultVUnknownFutureValue undocumented
	OperationResultVUnknownFutureValue OperationResult = "unknownFutureValue"
)

var (
	// OperationResultPSuccess is a pointer to OperationResultVSuccess
	OperationResultPSuccess = &_OperationResultPSuccess
	// OperationResultPFailure is a pointer to OperationResultVFailure
	OperationResultPFailure = &_OperationResultPFailure
	// OperationResultPTimeout is a pointer to OperationResultVTimeout
	OperationResultPTimeout = &_OperationResultPTimeout
	// OperationResultPUnknownFutureValue is a pointer to OperationResultVUnknownFutureValue
	OperationResultPUnknownFutureValue = &_OperationResultPUnknownFutureValue
)

var (
	_OperationResultPSuccess            = OperationResultVSuccess
	_OperationResultPFailure            = OperationResultVFailure
	_OperationResultPTimeout            = OperationResultVTimeout
	_OperationResultPUnknownFutureValue = OperationResultVUnknownFutureValue
)

// OperationStatus undocumented
type OperationStatus string

const (
	// OperationStatusVNotStarted undocumented
	OperationStatusVNotStarted OperationStatus = "NotStarted"
	// OperationStatusVRunning undocumented
	OperationStatusVRunning OperationStatus = "Running"
	// OperationStatusVCompleted undocumented
	OperationStatusVCompleted OperationStatus = "Completed"
	// OperationStatusVFailed undocumented
	OperationStatusVFailed OperationStatus = "Failed"
)

var (
	// OperationStatusPNotStarted is a pointer to OperationStatusVNotStarted
	OperationStatusPNotStarted = &_OperationStatusPNotStarted
	// OperationStatusPRunning is a pointer to OperationStatusVRunning
	OperationStatusPRunning = &_OperationStatusPRunning
	// OperationStatusPCompleted is a pointer to OperationStatusVCompleted
	OperationStatusPCompleted = &_OperationStatusPCompleted
	// OperationStatusPFailed is a pointer to OperationStatusVFailed
	OperationStatusPFailed = &_OperationStatusPFailed
)

var (
	_OperationStatusPNotStarted = OperationStatusVNotStarted
	_OperationStatusPRunning    = OperationStatusVRunning
	_OperationStatusPCompleted  = OperationStatusVCompleted
	_OperationStatusPFailed     = OperationStatusVFailed
)
