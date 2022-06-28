package jsonrpc

import (
	"encoding/json"
	"log"
	"sync/atomic"
)

const (
	bufferSize = 64
)

type Client struct {
	rx chan []byte
	tx chan []byte

	shutdown chan struct{}

	pendindRequets map[uint64]CallRequest
	cancelPending  chan uint64
	sendPending    chan CallRequest
	idCounter      uint64

	listeners map[string]Listener
}

func (c *Client) Call(req CallRequest) {
	c.sendPending <- req
}

func (c *Client) Cancel(id uint64) {
	c.cancelPending <- id
}

func (c *Client) NextId() uint64 {
	return atomic.AddUint64(&c.idCounter, 1)
}

func (c *Client) cancelAll() {
	// FIXME: Non served requests in input channel could be lost
	for _, req := range c.pendindRequets {
		req.Cancel()
	}
}

type jsonRpcHeader struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      uint64 `json:"id"`
	Method  string `json:"method"`
}

func (c *Client) routine() {
	defer close(c.cancelPending)
	defer close(c.sendPending)
	defer c.cancelAll()

	for {
		select {
		case <-c.shutdown:
			return
		case id := <-c.cancelPending:
			req, ok := c.pendindRequets[id]
			if ok {
				req.Cancel()
				delete(c.pendindRequets, id)
			} else {
				// FIXME: it should not happen, but better handle this correctly
				log.Println("Cancel non-scheduled method")
			}
		case req := <-c.sendPending:
			if req.Send(c.tx) {
				if req.GetId() != 0 {
					c.pendindRequets[req.GetId()] = req
				}
			} else {
				if req.GetId() != 0 {
					req.Cancel()
				}
			}
		case data, ok := <-c.rx:
			if !ok {
				panic("Rx channel closed")
			}

			//TODO: think about smard dynamic id-based unmarshal instead of two unmarshal calls

			header := jsonRpcHeader{}
			err := json.Unmarshal(data, &header)
			if err != nil {
				//TODO: make this log debug
				log.Println("Failed to parse jsonrpc header: ", err)
				continue
			}

			if header.Method == "" {
				req, ok := c.pendindRequets[header.Id]
				if !ok {
					//TODO: make this log debug
					log.Println("Response to non-existent request")
					continue
				}
				go req.Handle(data)
				delete(c.pendindRequets, req.GetId())
			} else {
				listener, ok := c.listeners[header.Method]
				if !ok {
					//TODO: reply with error
					log.Println("Call unsupported method ", header.Method)
					continue
				}
				go func(data []byte) {
					response, err := listener.Handle(data)
					if err != nil {
						//TODO: reply with error
						log.Println("Call failed ", err)
						return
					}
					if response != nil {
						c.tx <- response
					}
				}(data)
			}
		}
	}
}

func (c *Client) Close() {
	close(c.shutdown)
}

func (c *Client) Listen(listener Listener) {
	c.listeners[listener.GetMethod()] = listener
}

func NewClient(rx chan []byte, tx chan []byte) *Client {
	client := &Client{
		rx: rx,
		tx: tx,

		shutdown: make(chan struct{}, 1),

		pendindRequets: map[uint64]CallRequest{},

		cancelPending: make(chan uint64, bufferSize),
		sendPending:   make(chan CallRequest, bufferSize),
		idCounter:     0,

		listeners: map[string]Listener{},
	}

	go client.routine()
	return client
}
