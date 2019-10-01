package main

import (
	"sort"
	"strconv"
	"strings"
	"sync"
)

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	for i := range in {
		wg.Add(1)
		go WorkerSingleHash(i.(int), out, mutex, wg)
	}
	wg.Wait()
}

func WorkerSingleHash(i int, out chan interface{}, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	mutex.Lock()
	md5 := DataSignerMd5(strconv.Itoa(i))
	mutex.Unlock()

	crc32 := make(chan string)
	go Calculate(crc32, strconv.Itoa(i))

	md5Chan := make(chan string)
	go Calculate(md5Chan, md5)

	out <- <-crc32 + "~" + <-md5Chan
}

func Calculate(ch chan string, s string) {
	ch <- DataSignerCrc32(s)
}

func MultiHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	for i := range in {
		wg.Add(1)
		go WorkerMultiHash(i.(string), out, wg)
	}
	wg.Wait()
}

func WorkerMultiHash(str string, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	waitGroup := &sync.WaitGroup{}

	arr := make([]string, 6)
	for index := 0; index < 6; index++ {
		waitGroup.Add(1)
		go func(str string, index int, arr []string, waitGroup *sync.WaitGroup) {
			defer waitGroup.Done()
			arr[index] = DataSignerCrc32(str)
		}(strconv.Itoa(index)+str, index, arr, waitGroup)
	}

	waitGroup.Wait()
	out <- strings.Join(arr, "")
}

func CombineResults(in, out chan interface{}) {
	var data []string
	for i := range in {
		data = append(data, i.(string))
	}
	sort.Strings(data)
	out <- strings.Join(data, "_")
}

func WorkerPipeline(in, out chan interface{}, job job, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	job(in, out)
}

func ExecutePipeline(jobs ...job) {
	in := make(chan interface{})
	wg := &sync.WaitGroup{}
	for i := 0; i < len(jobs); i++ {
		wg.Add(1)
		out := make(chan interface{})
		go WorkerPipeline(in, out, jobs[i], wg)
		in = out
	}
	wg.Wait()
}
