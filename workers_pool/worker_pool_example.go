package main

import (
	"fmt"
	"sync"
	"time"
)

/*In this example, we define a Job struct that represents a task to be executed by the worker pool.
We also define a Result struct that contains the output of a task. */

func main() {
	startTime := time.Now()
	noOfJobs := 50
	go allocate(noOfJobs)

	done := make(chan bool)
	go result(done)

	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func createWorkerPool(workers int) {
	var weightgroup sync.WaitGroup
	for i := 0; i < workers; i++ {
		weightgroup.Add(1)
		go worker(&weightgroup)
	}
	weightgroup.Wait()
	close(results)
}

func worker(weight_group *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	weight_group.Done()
}

type Result struct {
	job         Job
	sumofdigits int
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

type Job struct {
	id       int
	randomno int
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := 1
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}
