package orbsclient

import (
	"bytes"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const CONTENT_TYPE = "application/membuffers"
const (
	SEND_TRANSACTION_URL       = "/api/v1/send-transaction"
	CALL_METHOD_URL            = "/api/v1/call-method"
	GET_TRANSACTION_STATUS_URL = "/api/v1/get-transaction-status"
)

func (c *OrbsClient) SendTransaction(payload []byte) (response *codec.SendTransactionResponse, err error) {
	res, buf, err := c.sendHttpPost(SEND_TRANSACTION_URL, payload)
	if err != nil {
		return
	}

	response, err = codec.DecodeSendTransactionResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status code %d: %s", res.StatusCode, res.Status)
		return
	}

	return
}

func (c *OrbsClient) CallMethod(payload []byte) (response *codec.CallMethodResponse, err error) {
	res, buf, err := c.sendHttpPost(CALL_METHOD_URL, payload)
	if err != nil {
		return
	}

	// TODO: improve handling of errors according to content-type header (if text/plain then don't parse response)

	response, err = codec.DecodeCallMethodResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status code %d (%s)", res.StatusCode, res.Status)
		return
	}

	return
}

func (c *OrbsClient) GetTransactionStatus(payload []byte) (response *codec.GetTransactionStatusResponse, err error) {
	res, buf, err := c.sendHttpPost(GET_TRANSACTION_STATUS_URL, payload)
	if err != nil {
		return
	}

	response, err = codec.DecodeGetTransactionStatusResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status code %d: %s", res.StatusCode, res.Status)
		return
	}

	return
}

func (c *OrbsClient) sendHttpPost(relativeUrl string, payload []byte) (*http.Response, []byte, error) {
	if len(payload) == 0 {
		return nil, nil, errors.New("payload sent by http is empty")
	}

	res, err := http.Post(c.Endpoint+relativeUrl, CONTENT_TYPE, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed sending http post")
	}

	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, buf, errors.Wrap(err, "failed reading http response")
	}

	return res, buf, nil
}
