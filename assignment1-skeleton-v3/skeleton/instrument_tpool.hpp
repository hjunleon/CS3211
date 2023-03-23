#ifndef INSTR_TPOOL_HPP
#define INSTR_TPOOL_HPP

#define MAX_INSTR_CNT 1
#include "common_types.hpp"
#include "block_queue.hpp"
#include "block_p_q.hpp"
#include <atomic>
#include <mutex>
#include <thread>
#include <condition_variable>
#include <vector>

class instrument_tpool
{
private:
    /* data */

        std::mutex order_sync_mtx;
        std::condition_variable order_sync_cond;
        std::atomic<unsigned long> order_cnt{0};
        std::atomic<unsigned long> next_order{0};
        std::vector<std::thread> t_list;
        JobQueue<Order> unprocessed_orders;
        JobPQ<R_order> r_sell_orders;
        JobPQ<R_order> r_buy_orders;

        void spawn_instrument_thread();
        void instr_matching_thread();
public:
    instrument_tpool(char instrument[9]);
    void insertOrder(Order o);
    ~instrument_tpool();


    std::string instr;
};


#endif