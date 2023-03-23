package main

import "container/list"

type AddToBook struct {
	order  Order
	res_ch chan struct{}
}

type OrderBook struct {
	match_buy_ch    chan Order
	match_sell_ch   chan Order
	delete_order_ch chan Order
	event_cnt_ch    chan chan int
	order_queue     chan uint32
	best_ask_ch     chan chan O_Wrapper
	add_to_sell_ch  chan AddToBook
	best_bid_ch     chan chan O_Wrapper
	add_to_buy_ch   chan AddToBook
}

func init_orderbook() OrderBook {
	ob := OrderBook{
		make(chan Order, 32),
		make(chan Order, 32),
		make(chan Order, 32),
		ob_event_counter(),
		create_order_queue(),
		make(chan chan O_Wrapper),
		make(chan AddToBook),
		make(chan chan O_Wrapper),
		make(chan AddToBook),
	}
	// go ob.match_order()
	go ob.match_buy_order()
	go ob.match_sell_order()

	// for i := 0; i < 4; i++ {
	// 	go ob.match_buy_order()
	// 	go ob.match_sell_order()
	// }
	return ob
}

func ob_event_counter() chan chan int {
	event_cnt_channel := make(chan chan int)
	go func() {
		event_cnt := 1
		defer close(event_cnt_channel)
		for {
			select {
			case <-done:
				return
			case send_ch := <-event_cnt_channel:
				event_cnt++
				send_ch <- event_cnt
			}
		}
	}()
	return event_cnt_channel
}

// creates an order queue to serialise which order to handle first
// Order queue is a buffered channel that can accept reading and sending or order ids to
func create_order_queue() chan uint32 {
	order_queue := make(chan uint32, 32)
	go func() {
		defer close(order_queue)
		for {
			select {
			case <-done:
				return
			}
		}
	}()
	return order_queue
}

// func da_sell_book() {

// 	best_ask_ch := make(chan chan O_Wrapper)
// 	add_to_sell_ch := make(chan struct{
// 		o Order
// 		tes chan struct {}
// 	})
// 	go func(){
// 		sell_book := list.New()
// 		for {

// 		}
// 	}()
// 	return {
// 		best_ask_ch
// 	}

// }

func (ob *OrderBook) da_sell_book() {
	// sell_book := list.New()
	// push dummy head
	// sell_book.PushBack(init_o_wrap(DUMMY_ORDER))
	// sell_head := sell_book.Front()
	sell_head := init_o_wrap(DUMMY_ORDER)
	for {
		select {
		case <-done:
			return
		case ba_ret_ch := <-ob.best_ask_ch:
			curNode := sell_head
			// curNode.o_lock <- struct{}{}
			for {
				curNode.o_lock <- struct{}{}
				if curNode.o.count == 0 {
					break
				}

				curNode = *curNode.o_next
			}
			ba_ret_ch <- curNode

		case req := <-ob.add_to_sell_ch:
			curNode := sell_head
			isAdded := false
			// checks if this order is the best price ask so far
			curNode.o_lock <- struct{}{}
			if curNode.o.price > req.order.price || (curNode.o.price == req.order.price && curNode.o.timestamp < req.order.timestamp) {
				newHead := init_o_wrap(req.order)
				newHead.o_next = &sell_head
				sell_head = newHead
				isAdded = true
			}
			<-curNode.o_lock
			if isAdded {
				req.res_ch <- struct{}{}
				continue
			}

			// checks where in the middle should place order
			// locks cur and next nodes
			for {
				curNode.o_lock <- struct{}{}
				nextNode := curNode.o_next
				nextNode.o_lock <- struct{}{}
				if curNode.o.price < req.order.price && req.order.price < nextNode.o.price {
					isAdded = true
				} else if curNode.o.price == req.order.price {

					if req.order.price == nextNode.o.price && req.order.timestamp < nextNode.o.timestamp {
						// within the same-price region
						isAdded = true
					} else if req.order.price < nextNode.o.price {
						// end of same-price region
						isAdded = true
					}
				}

				<-curNode.o_lock
				curNode = *nextNode
				nextNode = nextNode.o_next
				<-curNode.o_lock
			}
		}
	}
}

func (ob *OrderBook) da_buy_book() {
	buy_book := list.New()
	for {

	}
}

// create a routine that matches buy orders from a specific instrument
func (ob *OrderBook) match_buy_order() {

	// sell_book := list.New()

	go func() {
		for {
			select {
			case <-done:
				return
			case order := <-ob.match_buy_ch:
				ob.order_queue <- order.orderId

			}
		}
	}()
}

// create a routine that matches buy orders from a specific instrument
func (ob *OrderBook) match_sell_order() {

}
