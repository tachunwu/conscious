package main

import (
	"log"
	"runtime"
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

	// Test indiviual exist
	e.Exist()
	e.Observe()
	i.Exist(e.ObserveCh)
	for {
		runtime.Gosched()
	}
}

type Individual struct {
	ConsciousCh chan struct{}
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
