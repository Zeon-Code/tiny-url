package test

import (
	"context"
	"time"
)

type FakeMetric struct {
	LastMemoryInvalid     bool
	LastMemoryHitKey      string
	LastMemoryHitLatency  time.Duration
	LastMemoryMissKey     string
	LastMemoryMissLatency time.Duration
	LastMemoryBypass      bool
}

func NewFakeMetric() *FakeMetric {
	return &FakeMetric{}
}

func (m *FakeMetric) MemoryHit(ctx context.Context, key string, duration time.Duration) {
	m.LastMemoryHitKey = key
	m.LastMemoryHitLatency = duration
}

func (m *FakeMetric) MemoryMiss(ctx context.Context, key string, duration time.Duration) {
	m.LastMemoryMissKey = key
	m.LastMemoryMissLatency = duration
}

func (m *FakeMetric) MemoryInvalid(ctx context.Context, key string) {
	m.LastMemoryInvalid = true
}

func (m *FakeMetric) MemoryBypassed(ctx context.Context) {
	m.LastMemoryBypass = true
}
