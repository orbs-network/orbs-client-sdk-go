package main

import (
	"bytes"
	"github.com/orbs-network/orbs-client-sdk-go/codec"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

const CONTENT_TYPE = "application/membuffers"

func (c *OrbsClient) SendTransaction(payload []byte) (response *codec.SendTransactionResponse, err error) {
	res, err := http.Post(c.Endpoint, CONTENT_TYPE, bytes.NewReader(payload))
	if err != nil {
		err = errors.Wrap(err, "failed sending http post")
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		err = errors.Wrap(err, "failed reading http response")
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
	res, err := http.Post(c.Endpoint, CONTENT_TYPE, bytes.NewReader(payload))
	if err != nil {
		err = errors.Wrap(err, "failed sending http post")
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		err = errors.Wrap(err, "failed reading http response")
		return
	}

	response, err = codec.DecodeCallMethodResponse(buf)
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

func (c *OrbsClient) GetTransactionStatus(payload []byte) (response *codec.GetTransactionStatusResponse, err error) {
	res, err := http.Post(c.Endpoint, CONTENT_TYPE, bytes.NewReader(payload))
	if err != nil {
		err = errors.Wrap(err, "failed sending http post")
		return
	}

	buf, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		err = errors.Wrap(err, "failed reading http response")
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
