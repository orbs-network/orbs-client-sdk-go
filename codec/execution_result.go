package codec

import (
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/pkg/errors"
)

type ExecutionResult string

const (
	EXECUTION_RESULT_SUCCESS               ExecutionResult = "SUCCESS"
	EXECUTION_RESULT_ERROR_SMART_CONTRACT  ExecutionResult = "ERROR_SMART_CONTRACT"
	EXECUTION_RESULT_ERROR_INPUT           ExecutionResult = "ERROR_INPUT"
	EXECUTION_RESULT_ERROR_UNEXPECTED      ExecutionResult = "ERROR_UNEXPECTED"
	EXECUTION_RESULT_STATE_WRITE_IN_A_CALL ExecutionResult = "STATE_WRITE_IN_A_CALL"
)

func executionResultDecode(executionResult protocol.ExecutionResult) (ExecutionResult, error) {
	switch executionResult {
	case protocol.EXECUTION_RESULT_RESERVED:
		return "", errors.Errorf("reserved ExecutionResult received")
	case protocol.EXECUTION_RESULT_SUCCESS:
		return EXECUTION_RESULT_SUCCESS, nil
	case protocol.EXECUTION_RESULT_ERROR_SMART_CONTRACT:
		return EXECUTION_RESULT_ERROR_SMART_CONTRACT, nil
	case protocol.EXECUTION_RESULT_ERROR_INPUT:
		return EXECUTION_RESULT_ERROR_INPUT, nil
	case protocol.EXECUTION_RESULT_ERROR_UNEXPECTED:
		return EXECUTION_RESULT_ERROR_UNEXPECTED, nil
	case protocol.EXECUTION_RESULT_STATE_WRITE_IN_A_CALL:
		return EXECUTION_RESULT_STATE_WRITE_IN_A_CALL, nil
	default:
		return "", errors.Errorf("unsupported ExecutionResult received: %d", executionResult)
	}
}
