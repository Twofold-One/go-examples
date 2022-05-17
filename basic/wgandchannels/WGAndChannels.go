package wgandchannels

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"

	"github.com/Twofold-One/go-examples/basic/wgandchannels/visit"
)

type Task struct {
	Date   string
	Visits []visit.Visit
}

// output struct
type DailyStat struct {
	Date   string         `json:"date"`
	ByPage map[string]int `json:"byPage"`
}

func WGAndChannelsExample() {
	data, err := ioutil.ReadFile("/home/twofold_one/GitProjects/go/go-examples/basic/wgandchannels/data.json")
	if err != nil {
		log.Fatal(err)
	}
	dayStats := make(map[string][]visit.Visit)
	err = json.Unmarshal(data, &dayStats)
	if err != nil {
		log.Fatal(err)
	}

	// wait group / channel creation
	var w8 sync.WaitGroup
	w8.Add(len(dayStats))

	inputCh := make(chan Task, 10)
	outputCh := make(chan DailyStat, len(dayStats))

	// create the workers
	numberOfWorkers := 10
	for k := 0; k < numberOfWorkers; k++ {
		go worker(inputCh, k, outputCh, &w8)
	}
	// send the tasks
	for date, visits := range dayStats {
		inputCh <- Task{
			Date:   date,
			Visits: visits,
		}
	}
	// we say that we will not send any new data on the input channel
	close(inputCh)
	// wait for all tasks to be complited
	w8.Wait()
	// when all treatment is finished we close the output channel
	close(outputCh)
	// collect the result
	done := make([]DailyStat, 0, len(dayStats))
	for out := range outputCh {
		done = append(done, out)
	}

	res, err := json.Marshal(done)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("/home/twofold_one/GitProjects/go/go-examples/basic/wgandchannels/result.json", res, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")
}

func worker(in chan Task, workerId int, out chan DailyStat, w8 *sync.WaitGroup) {
	for received := range in {
		m := make(map[string]int)
		for _, v := range received.Visits {
			m[v.Page]++
		}
		out <- DailyStat{
			Date:   received.Date,
			ByPage: m,
		}
		log.Printf("[worker %d] finished task\n", workerId)
	}
	// when the channel is closed the for loop is exited
	log.Println("worker quit")
	w8.Done()
}
