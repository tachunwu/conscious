package main

import (
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Init individual conscious
	i := &Individual{
		ConsciousCh: make(chan struct{}, 1),
	}

	// Set environment
	e := &Environment{
		ObserveCh: make(chan struct{}, 1),
	}
	e.Include(i)

	// Producer function
	go func() {
		for {
			i.Serve(func(i *Individual) {
				i.Lock.Lock()
				defer i.Lock.Unlock()
				// Get
				log.Println(i.State)
				// Set
				i.State = i.State + 1
				// Get
				log.Println(i.State)
			})
			time.Sleep(1 * time.Second)
		}
	}()

	// Consumer function
	go func() {
		for {
			i.Serve(func(i *Individual) {
				i.Lock.Lock()
				defer i.Lock.Unlock()
				// Get
				log.Println(i.State)
				// Set
				i.State = i.State - 1
				// Get
				log.Println(i.State)
			})
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		runtime.Gosched()
	}
}

type Individual struct {
	ConsciousCh chan struct{}
	State       int
	Lock        sync.Mutex
}

func (i *Individual) Exist(observeCh chan struct{}) {
	go func() {
		for {
			<-i.ConsciousCh
			observeCh <- struct{}{}
			log.Println("Individual exist")
			time.Sleep(time.Second)
		}
	}()
}

func (i *Individual) Serve(fn func(i *Individual)) {
	fn(i)
}

type Environment struct {
	ObserveCh   chan struct{}
	Individuals []*Individual
}

func (e *Environment) Include(i *Individual) {
	e.Individuals = append(e.Individuals, i)
}

func (e *Environment) Exist() {
	go func() {
		for {
			for _, i := range e.Individuals {
				i.ConsciousCh <- struct{}{}
			}
			time.Sleep(time.Second)
		}
	}()
}

func (e *Environment) Observe() {
	go func() {
		for {
			<-e.ObserveCh
			log.Println("Environment observe")
			time.Sleep(time.Second)
		}
	}()
}
