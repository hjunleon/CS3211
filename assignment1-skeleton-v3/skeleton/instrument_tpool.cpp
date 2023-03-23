#include "instrument_tpool.hpp"
#include "io.hpp"

instrument_tpool::instrument_tpool(char instrument[9])
{
    instr = instrument;
    for (int i = 0; i < MAX_INSTR_CNT; i++)
    {
        spawn_instrument_thread();
    }
}
void instrument_tpool::insertOrder(Order o)
{
    unprocessed_orders.enqueue(o);
}

void instrument_tpool::instr_matching_thread()
{
    while (true)
    {
        std::unique_lock<std::mutex> lk(order_sync_mtx);
        // order_sync_cond.wait(lk, [this](){return next_order < order_cnt;});
        order_sync_cond.wait(lk);
        Order ord = unprocessed_orders.dequeue();
        auto &r_q = (ord.side == SideType::buy) ? r_sell_orders : r_buy_orders;
        uint32_t cnt = ord.count;
        if (!r_q.empty())
        {
            auto r_order = r_q.dequeue();
            cnt = cnt - r_order.order.count;
            if (cnt >= 0)
            {
                // the resting order is fully matched
                // SyncCerr {} << "Got cancel: ID: " << input.order_id << std::endl;
                Output::OrderExecuted(
                    r_order.order.order_id,
                    ord.order_id,
                    r_order.e_id,
                    r_order.order.price,
                    r_order.order.count,
                    getCurrentTimestamp());
            }
            else
            {
                // the resting order is partially matched
                r_order.e_id += 1;
                Output::OrderExecuted(
                    r_order.order.order_id,
                    ord.order_id,
                    r_order.e_id,
                    r_order.order.price,
                    r_order.order.count,
                    getCurrentTimestamp()
                );

                // R_order tmp = R_order{
                //     r_order.order,
                //     r_order.e_id,
                // };
                
                r_q.enqueue(std::move(r_order));
            }
        }
        if (cnt > 0)
        {
            Output::OrderAdded(
                ord.order_id,
                ord.instrument,
                ord.price,
                ord.count,
                (ord.side == SideType::sell),
                getCurrentTimestamp());
            R_order new_r_order = {
                std::move(ord),
                1
            };
            r_q.enqueue(new_r_order);
        }
    }
    // return void;
};

void instrument_tpool::spawn_instrument_thread()
{
    if (t_list.size() == MAX_INSTR_CNT)
        return;
    auto thread = std::thread(&instrument_tpool::instr_matching_thread, instr);
    thread.detach();
    t_list.push_back(thread);
};

instrument_tpool::~instrument_tpool()
{
}
