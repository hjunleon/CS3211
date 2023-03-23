// This file contains declarations for the main Engine class. You will
// need to add declarations to this file as you develop your Engine.

#ifndef ENGINE_HPP
#define ENGINE_HPP

#include <chrono>

#include "io.hpp"
#include "order_book.hpp"

struct Engine
{
	// void Engine();
public:
	void accept(ClientConnection conn);

private:
	void connection_thread(ClientConnection conn);
	OrderBook book;
};



#endif
