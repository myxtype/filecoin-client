package filecoin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync/atomic"
)

type clientRequest struct {
	Id      int64         `json:"id"`
	Version string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

func (r *clientRequest) Bytes() []byte {
	b, _ := json.Marshal(r)
	return b
}

type clientResponse struct {
	Id      uint64           `json:"id"`
	Version string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result"`
	Error   interface{}      `json:"error,omitempty"`
}

func (c *clientResponse) ReadFromResult(x interface{}) error {
	if x == nil {
		return nil
	}
	return json.Unmarshal(*c.Result, x)
}

type Client struct {
	addr  string
	token string
	id    int64
}

func NewClient(addr string, token string) *Client {
	return &Client{
		addr:  addr,
		token: token,
	}
}

func (c *Client) FilecoinMethod(method string) string {
	return fmt.Sprintf("Filecoin.%s", method)
}

func (c *Client) Request(ctx context.Context, method string, result interface{}, params ...interface{}) error {
	request := &clientRequest{
		Id:      atomic.AddInt64(&c.id, 1),
		Version: "2.0",
		Method:  method,
		Params:  params,
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.addr, bytes.NewReader(request.Bytes()))
	if err != nil {
		return err
	}
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	response := &clientResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return err
	}
	if response.Error != nil {
		return fmt.Errorf("jsonrpc call: %v", response.Error)
	}

	return response.ReadFromResult(result)
}
