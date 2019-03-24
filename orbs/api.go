// Copyright 2019 the orbs-client-sdk-go authors
// This file is part of the orbs-client-sdk-go library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package orbs

import (
	"bytes"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const CONTENT_TYPE_MEMBUFFERS = "application/membuffers"
const (
	SEND_TRANSACTION_URL              = "/api/v1/send-transaction"
	SEND_TRANSACTION_ASYNC_URL        = "/api/v1/send-transaction-async"
	CALL_METHOD_URL                   = "/api/v1/run-query"
	GET_TRANSACTION_STATUS_URL        = "/api/v1/get-transaction-status"
	GET_TRANSACTION_RECEIPT_PROOF_URL = "/api/v1/get-transaction-receipt-proof"
	GET_BLOCK_URL                     = "/api/v1/get-block"
)

func (c *OrbsClient) SendTransaction(rawTransaction []byte) (response *codec.SendTransactionResponse, err error) {
	res, buf, err := c.sendHttpPost(SEND_TRANSACTION_URL, rawTransaction)
	if err != nil {
		return
	}

	response, err = codec.DecodeSendTransactionResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) SendTransactionAsync(rawTransaction []byte) (response *codec.SendTransactionResponse, err error) {
	res, buf, err := c.sendHttpPost(SEND_TRANSACTION_ASYNC_URL, rawTransaction)
	if err != nil {
		return
	}

	response, err = codec.DecodeSendTransactionResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusAccepted {
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) SendQuery(rawQuery []byte) (response *codec.RunQueryResponse, err error) {
	res, buf, err := c.sendHttpPost(CALL_METHOD_URL, rawQuery)
	if err != nil {
		return
	}

	// TODO: improve handling of errors according to content-type header (if text/plain then don't parse response)

	response, err = codec.DecodeRunQueryResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) GetTransactionStatus(txId string) (response *codec.GetTransactionStatusResponse, err error) {
	payload, err := c.createGetTransactionStatusPayload(txId)
	if err != nil {
		return
	}

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
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) GetTransactionReceiptProof(txId string) (response *codec.GetTransactionReceiptProofResponse, err error) {
	payload, err := c.createGetTransactionReceiptProofPayload(txId)
	if err != nil {
		return
	}

	res, buf, err := c.sendHttpPost(GET_TRANSACTION_RECEIPT_PROOF_URL, payload)
	if err != nil {
		return
	}

	response, err = codec.DecodeGetTransactionReceiptProofResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) GetBlock(blockHeight uint64) (response *codec.GetBlockResponse, err error) {
	payload, err := c.createGetBlockPayload(blockHeight)
	if err != nil {
		return
	}

	res, buf, err := c.sendHttpPost(GET_BLOCK_URL, payload)
	if err != nil {
		return
	}

	response, err = codec.DecodeGetBlockResponse(buf)
	if err != nil {
		err = errors.Wrap(err, "failed decoding response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = errors.Errorf("http status %s", res.Status)
		return
	}

	return
}

func (c *OrbsClient) sendHttpPost(relativeUrl string, payload []byte) (*http.Response, []byte, error) {
	if len(payload) == 0 {
		return nil, nil, errors.New("payload sent by http is empty")
	}

	res, err := http.Post(c.Endpoint+relativeUrl, CONTENT_TYPE_MEMBUFFERS, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed sending http post")
	}

	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, buf, errors.Wrap(err, "failed reading http response")
	}

	// check if we have the content type response we expect
	contentType := res.Header.Get("Content-Type")
	if contentType != CONTENT_TYPE_MEMBUFFERS {

		// handle real 404 (incorrect endpoint) gracefully
		if res.StatusCode == 404 {
			// TODO: streamline these errors
			return res, buf, errors.Wrap(NoConnectionError, "http 404 not found")
		}

		if contentType == "text/plain" || contentType == "application/json" {
			return nil, buf, errors.Errorf("http request failed: %s", string(buf))
		} else {
			return nil, buf, errors.Errorf("http request failed with Content-Type '%s': %x", contentType, buf)
		}
	}

	return res, buf, nil
}

// TODO: streamline these errors
var NoConnectionError = errors.New("cannot connect to server")
