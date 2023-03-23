#ifndef ORDER_BOOK_HPP
#define ORDER_BOOK_HPP

#include "common_types.hpp"
#include "block_queue.hpp"
#include "instrument_tpool.hpp"

#include <map>
#include <atomic>
#include <unordered_map>
#include <string>
#include <thread>



class OrderBook {
    public:
        OrderBook();
        void insert(Order o);
        void cancel();

    private:
        // map of instrument to vect<threads>
        std::map<char[9], instrument_tpool> instr_t_p;

};


#endif
