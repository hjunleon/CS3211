package main

import "C"
import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Engine struct {
	// done chan struct{}
}

var req_ob_ch = make(chan string)
var res_ob_ch = create_ob_ch(req_ob_ch)

// var reqResCh = make(chan chan OrderBook)
// var _ = create_ob_ch(req_ob_ch)

var done = make(chan struct{})

// func create_ob_ch() <-chan chan<- OrderBook {
// 	res_ob_ch := make(<-chan chan<- OrderBook)
// 	orderbooks := map[string]OrderBook{}
// 	go func() {
// 		defer close(res_ob_ch)
// 		for {
// 			select {
// 			case <-done:
// 				return
// 			case inst := <-req_ob_ch:
// 				val, ok := orderbooks[inst]
// 				if !ok {
// 					val = OrderBook{
// 						make(chan Order),
// 						make(chan Order),
// 						make(chan Order),
// 					}
// 					orderbooks[inst] = val
// 				}
// 				res_ob_ch <- val
// 			}
// 		}
// 	}()
// 	return res_ob_ch
// }

func create_ob_ch(req_ob_ch <-chan string) chan OrderBook {
	res_ob_ch := make(chan OrderBook)
	orderbooks := map[string]OrderBook{}
	go func() {
		defer close(res_ob_ch)
		for {
			select {
			case <-done:
				return
			case inst := <-req_ob_ch:
				val, ok := orderbooks[inst]
				if !ok {
					val = init_orderbook()
					orderbooks[inst] = val
				}
				res_ob_ch <- val
			}
		}
	}()
	return res_ob_ch
}

func (e *Engine) accept(ctx context.Context, conn net.Conn) {
	log_debug(2, DEBUG_EG, "accepting connection ", "test")
	go func() {
		<-ctx.Done()
		conn.Close()
	}()
	go handleConn(conn)
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		in, err := readInput(conn)
		if err != nil {
			if err != io.EOF {
				_, _ = fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			}
			return
		}
		switch in.orderType {
		case inputCancel:
			fmt.Fprintf(os.Stderr, "Got cancel ID: %v\n", in.orderId)
			outputOrderDeleted(in, true, GetCurrentTimestamp())
		default:
			fmt.Fprintf(os.Stderr, "Got order: %c %v x %v @ %v ID: %v\n",
				in.orderType, in.instrument, in.count, in.price, in.orderId)
			req_ob_ch <- in.instrument
			ob := <-res_ob_ch
			order := Order{
				in.orderId,
				in.price,
				in.count,
				GetCurrentTimestamp(),
			}
			ob.order_queue <- in.orderId
			if in.orderType == inputBuy {
				ob.match_buy_ch <- order
			} else {
				ob.match_sell_ch <- order
			}
			outputOrderAdded(in, GetCurrentTimestamp())
		}
		outputOrderExecuted(123, 124, 1, 2000, 10, GetCurrentTimestamp())
	}
}

func GetCurrentTimestamp() int64 {
	return time.Now().UnixNano()
}
