package etcd

import (
	"net/http"
	"time"
)

const (
	rawResponse = iota
	normalResponse
)

type responseType int

type RawResponse struct {
	Body   []byte
	Header http.Header
}

func toRawResp(i interface{}, err error) (*RawResponse, error) {
	if err != nil {
		return nil, err
	}

	resp := i.(*RawResponse)

	return resp, nil
}

// The response object from the server.
type Response struct {
	Action    string  `json:"action"`
	Key       string  `json:"key"`
	Dir       bool    `json:"dir,omitempty"`
	PrevValue string  `json:"prevValue,omitempty"`
	Value     string  `json:"value,omitempty"`
	Kvs       kvPairs `json:"kvs,omitempty"`

	// If the key did not exist before the action,
	// this field should be set to true
	NewKey bool `json:"newKey,omitempty"`

	Expiration *time.Time `json:"expiration,omitempty"`

	// Time to live in second
	TTL int64 `json:"ttl,omitempty"`

	// The command index of the raft machine when the command is executed
	ModifiedIndex uint64 `json:"modifiedIndex"`
}

func toResp(i interface{}, err error) (*Response, error) {

	if err != nil {
		return nil, err
	}

	resp := i.(*Response)

	return resp, nil
}

// When user list a directory, we add all the node into key-value pair slice
type KeyValuePair struct {
	Key     string  `json:"key, omitempty"`
	Value   string  `json:"value,omitempty"`
	Dir     bool    `json:"dir,omitempty"`
	KVPairs kvPairs `json:"kvs,omitempty"`
	TTL     int64   `json:"ttl,omitempty"`
}

type kvPairs []KeyValuePair

// interfaces for sorting
func (kvs kvPairs) Len() int {
	return len(kvs)
}

func (kvs kvPairs) Less(i, j int) bool {
	return kvs[i].Key < kvs[j].Key
}

func (kvs kvPairs) Swap(i, j int) {
	kvs[i], kvs[j] = kvs[j], kvs[i]
}
