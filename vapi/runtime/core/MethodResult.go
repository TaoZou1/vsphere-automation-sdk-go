package core

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/vapi/runtime/data"
)

type MethodResult struct {
	output data.DataValue
	error  *data.ErrorValue
}

func NewMethodResult(output data.DataValue, error *data.ErrorValue) MethodResult {
	return MethodResult{output: output, error: error}
}

func (methodResult MethodResult) Output() data.DataValue {
	return methodResult.output
}
func (methodResult MethodResult) Error() *data.ErrorValue {
	return methodResult.error
}
func (methodResult MethodResult) IsSuccess() bool {
	return methodResult.error == (*data.ErrorValue)(nil)
}