#ifndef COMMON_TYPES_H
#define COMMON_TYPES_H
#include <string>
#include <chrono>
#include <atomic>

enum SideType{
    buy = 'B',
    sell = 'S'
};

struct Order{
    uint32_t order_id;
    char instrument[9];
    uint32_t price;
    uint32_t count;
    int64_t timestamp;
    SideType side;
};

struct R_order{
    Order order;
    uint32_t e_id;  //std::atomic<uint32_t>
};

struct {
    bool operator() (const R_order l, const R_order r) const { return l.order.price > r.order.price}
} r_order_compare_buy;


struct {
    bool operator() (const R_order l, const R_order r) const { return l.order.price < r.order.price}
} r_order_compare_sell;


inline std::chrono::microseconds::rep getCurrentTimestamp() noexcept
{
	return std::chrono::duration_cast<std::chrono::nanoseconds>(std::chrono::steady_clock::now().time_since_epoch()).count();
}
#endif