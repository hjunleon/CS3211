#ifndef B_Q_HPP
#define B_Q_HPP

#include <mutex>
#include <optional>
#include <queue>
#include <condition_variable>
#include <thread>


template <typename T>


// Unbounded queue with enqueue, non-blocking and blocking dequeue
class JobQueue {
  std::queue<T> jobs;
  std::mutex mut;
  std::condition_variable cond;

 public:
  JobQueue() : jobs{}, mut{}, cond{} {}

  void enqueue(T job) {
    {
      std::unique_lock lock{mut};
      jobs.push(job);
    }

    cond.notify_one();
  }
    
  // std::optional<T> try_dequeue() {
  //   std::unique_lock lock{mut};

  //   if (jobs.empty()) {
  //     return std::nullopt;
  //   } else {
  //     T job = jobs.front();
  //     jobs.pop();

  //     return job;
  //   }
  // }

  T dequeue() {
    std::unique_lock lock{mut};
    
    while (jobs.empty()) {
        cond.wait(lock);
    }
    T job = jobs.front();
    jobs.pop();
    
    return job;
  }

  bool empty() {
    std::unique_lock lk{mut};
    return jobs.empty();
  }
};
#endif