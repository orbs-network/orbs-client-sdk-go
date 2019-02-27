package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

type ExecutionResult string

const (
	EXECUTION_RESULT_PARSE_ERROR                 ExecutionResult = "<PARSE_ERROR>"
	EXECUTION_RESULT_SUCCESS                     ExecutionResult = "SUCCESS"
	EXECUTION_RESULT_ERROR_SMART_CONTRACT        ExecutionResult = "ERROR_SMART_CONTRACT"
	EXECUTION_RESULT_ERROR_INPUT                 ExecutionResult = "ERROR_INPUT"
	EXECUTION_RESULT_ERROR_CONTRACT_NOT_DEPLOYED ExecutionResult = "ERROR_CONTRACT_NOT_DEPLOYED"
	EXECUTION_RESULT_ERROR_UNEXPECTED            ExecutionResult = "ERROR_UNEXPECTED"
	EXECUTION_RESULT_NOT_EXECUTED                ExecutionResult = "NOT_EXECUTED"
)

func (x ExecutionResult) String() string {
	return string(x)
}

func executionResultDecode(executionResult protocol.ExecutionResult) (ExecutionResult, error) {
	switch executionResult {
	case protocol.EXECUTION_RESULT_RESERVED:
		return EXECUTION_RESULT_PARSE_ERROR, errors.Errorf("reserved ExecutionResult received")
	case protocol.EXECUTION_RESULT_SUCCESS:
		return EXECUTION_RESULT_SUCCESS, nil
	case protocol.EXECUTION_RESULT_ERROR_SMART_CONTRACT:
		return EXECUTION_RESULT_ERROR_SMART_CONTRACT, nil
	case protocol.EXECUTION_RESULT_ERROR_INPUT:
		return EXECUTION_RESULT_ERROR_INPUT, nil
	case protocol.EXECUTION_RESULT_ERROR_CONTRACT_NOT_DEPLOYED:
		return EXECUTION_RESULT_ERROR_CONTRACT_NOT_DEPLOYED, nil
	case protocol.EXECUTION_RESULT_ERROR_UNEXPECTED:
		return EXECUTION_RESULT_ERROR_UNEXPECTED, nil
	case protocol.EXECUTION_RESULT_NOT_EXECUTED:
		return EXECUTION_RESULT_NOT_EXECUTED, nil
	default:
		return EXECUTION_RESULT_PARSE_ERROR, errors.Errorf("unsupported ExecutionResult received: %d", executionResult)
	}
}
