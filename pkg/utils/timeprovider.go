package utils

import "time"

type TimeProvider interface {
	Now() time.Time
	UTCNow() time.Time
}

type DefaultTimeProvider struct{}

func NewDefaultTimeProvider() *DefaultTimeProvider {
	return &DefaultTimeProvider{}
}

func (d *DefaultTimeProvider) Now() time.Time {
	return time.Now()
}

func (d *DefaultTimeProvider) UTCNow() time.Time {
	return time.Now().UTC()
}
