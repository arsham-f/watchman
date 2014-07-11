package main

import (
	. "time"
)

type Stopwatch struct {
	start Time
	end   Time
}

func (s *Stopwatch) Start() {
	s.start = Now()
}

func (s *Stopwatch) Stop(label string) {
	s.end = Now()
	Infof("%s took %s", label, s.end.Sub(s.start))
}
