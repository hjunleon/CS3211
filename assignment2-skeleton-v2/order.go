package main

type Order struct {
	orderId   uint32
	price     uint32
	count     uint32
	timestamp int64
}

type O_Wrapper struct {
	o      Order
	o_next *O_Wrapper
	o_lock chan struct{}
}

func init_o_wrap(o Order) O_Wrapper {
	return O_Wrapper{
		o,
		nil,
		make(chan struct{}, 1),
	}
}

var DUMMY_ORDER = Order{
	69,
	69,
	0,
	0,
}

// func (ob *OrderBook) match_order() {
// 	isBuy := false
// 	var order Order
// 	sell_book := list.New()
// 	buy_book := list.New()
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		case order = <-ob.match_buy_ch:
// 			isBuy = true
// 		case order = <-ob.match_sell_ch:
// 		}
// 		if isBuy {
// 			if sell_book.Len() == 0 {
// 				sell_book.PushBack(order)
// 				sell_book.
// 			}
// 		} else {
// 			if buy_book.Len() == 0 {
// 				buy_book.PushBack(order)
// 			}
// 		}

// 	}
// }
