#include "order_book.hpp"
#include "io.hpp"
#include "engine.hpp"

OrderBook::OrderBook(){
    // for(uint i = 0; i < MAX_INSTR_CNT; i++){
    //     spawn_instrument_thread();
    // }
};


void OrderBook::insert(Order o)
{
    // if(o.side == SideType::buy){
    //     // if (buy_orders.find(o.instrument) == buy_orders.end()){
    //     //     buy_orders[o.instrument] = new JobQueue<Order>();
    //     // }
    //     buy_orders[o.instrument].enqueue(o);
    // } else {
    //     // if (sell_orders.find(o.instrument) == sell_orders.end()){
    //     //     sell_orders[o.instrument] = new JobQueue<Order>();
    //     // }
    //     sell_orders[o.instrument].enqueue(o);
    // }
    instr_t_p[o.instrument].insertOrder(o);
    // orders.enqueue(o);
    // order_cnt.fetch_add(1);
    // order_sync_cond.notify_one();
    return;
};



void OrderBook::cancel()
{
    return;
};
