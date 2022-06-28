package jsonrpc

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type CallRequest interface {
	GetId() uint64
	Handle([]byte)
	//TODO: make like listener with return
	Send(chan []byte) bool
	Cancel()
}

type Listener interface {
	GetMethod() string
	Handle([]byte) ([]byte, error)
}

type Transport interface {
	Call(CallRequest)
	Cancel(id uint64)
	NextId() uint64
	Listen(listener Listener)
}

type Method[Request interface{}, Response interface{}] struct {
	Name string
}

type Notify[Request interface{}] struct {
	Name string
}

type responseOrError[Response interface{}] struct {
	response  *Response
	errorInfo error
}

type requestImpl[Request interface{}, Response interface{}] struct {
	request  Request
	method   string
	response chan responseOrError[Response]
	id       uint64
}

func (ri *requestImpl[Request, Response]) GetId() uint64 {
	return ri.id
}

type jsonRpcResponse[Response interface{}] struct {
	Jsonrpc string    `json:"jsonrpc"`
	Result  *Response `json:"result"`
	Id      uint64    `json:"id"`
}

func (ri *requestImpl[Request, Response]) Handle(data []byte) {
	resp := jsonRpcResponse[Response]{}
	err := json.Unmarshal(data, &resp)
	if err != nil {
		ri.response <- responseOrError[Response]{
			errorInfo: err,
		}
	} else {
		ri.response <- responseOrError[Response]{
			response: resp.Result,
		}
	}
	close(ri.response)
}

type jsonRpcRequest[Request interface{}] struct {
	Jsonrpc string  `json:"jsonrpc"`
	Method  string  `json:"method"`
	Params  Request `json:"params"`
	Id      *uint64 `json:"id,omitempty"`
}

func (ri *requestImpl[Request, Response]) Send(tx chan []byte) bool {
	req := jsonRpcRequest[Request]{
		Jsonrpc: "2.0",
		Method:  ri.method,
		Params:  ri.request,
	}

	if ri.GetId() != 0 {
		id := ri.GetId()
		req.Id = &id
	}

	data, err := json.Marshal(req)
	if err != nil {
		ri.response <- responseOrError[Response]{
			errorInfo: err,
		}
		return false
	}

	tx <- data

	return true

}

func (ri *requestImpl[Request, Response]) Cancel() {
	close(ri.response)
}

func (m *Method[Request, Response]) Call(request Request, timeout time.Duration, transport Transport) (*Response, error) {
	deadlineTicker := time.NewTicker(timeout)
	req := &requestImpl[Request, Response]{
		request:  request,
		method:   m.Name,
		response: make(chan responseOrError[Response], 1),
		id:       transport.NextId(),
	}
	transport.Call(req)

	defer deadlineTicker.Stop()

	select {
	case response, ok := <-req.response:
		if !ok {
			return nil, fmt.Errorf("Request cancelled")
		}
		return response.response, response.errorInfo
	case <-deadlineTicker.C:
		transport.Cancel(req.id)
		return nil, fmt.Errorf("Request timed out")
	}
}

func (m *Notify[Request]) Send(request Request, transport Transport) {
	req := &requestImpl[Request, struct{}]{
		request: request,
		method:  m.Name,
		id:      0,
	}
	transport.Call(req)
}

type listenerImpl[Request interface{}, Response interface{}] struct {
	method   string
	callback func(*Request) Response
}

func (l *listenerImpl[Request, Response]) GetMethod() string {
	return l.method
}

func (l *listenerImpl[Request, Response]) Handle(data []byte) ([]byte, error) {
	req := jsonRpcRequest[Request]{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, err
	}

	resp := l.callback(&req.Params)

	if req.Id == nil {
		return nil, nil
	}

	response := jsonRpcResponse[Response]{
		Jsonrpc: "2.0",
		Id:      *req.Id,
		Result:  &resp,
	}

	if reflect.TypeOf(resp) == reflect.TypeOf(struct{}{}) {
		response.Result = nil
	}

	respData, err := json.Marshal(response)

	return respData, err
}

func (m *Method[Request, Response]) Listen(callback func(*Request) Response, transport Transport) {
	transport.Listen(&listenerImpl[Request, Response]{
		method:   m.Name,
		callback: callback,
	})
}

func (m *Notify[Request]) Listen(callback func(*Request), transport Transport) {
	transport.Listen(&listenerImpl[Request, struct{}]{
		method: m.Name,
		callback: func(req *Request) struct{} {
			callback(req)
			return struct{}{}
		},
	})
}
