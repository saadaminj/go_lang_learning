package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {

	defer func(start time.Time) {
		fmt.Printf("Fact = %v ,GetCatFact took %s to complete\n", fact, time.Since(start))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
