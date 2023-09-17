package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

  // Here’s the worker goroutine. It repeatedly receives from jobs
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

  // This sends jobs to the worker over the jobs channel,
	for j := 1; j <= 100; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

  // We close the jobs channel to indicate that’s all the work we have.
  close(jobs)
  fmt.Println("sent all jobs")
}
