package application

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yakob-abada/go-api/go/app/entity"
)

func TestWithUpcomingAvailableSession(t *testing.T) {
	nowString := "2023-06-26 07:00:00"
	now, _ = time.Parse("2006-01-02 15:04:00", nowString)

	testedDateString := "2023-06-26 09:00:00"
	testedDate, _ := time.Parse("2006-01-02 15:04:00", testedDateString)

	session := &entity.Session{
		Id:       1,
		Time:     testedDate,
		Name:     "Spin",
		Duration: 30,
		IsFull:   false,
	}

	sut := NewActiveSessionSpecification()
	assert.True(t, sut.IsSatisfied(session))
}

func TestWithPassedAvailableSession(t *testing.T) {
	nowString := "2023-06-26 07:00:00"
	now, _ = time.Parse("2006-01-02 15:04:00", nowString)

	testedDateString := "2023-06-26 05:00:00"
	testedDate, _ := time.Parse("2006-01-02 15:04:00", testedDateString)

	session := &entity.Session{
		Id:       1,
		Time:     testedDate,
		Name:     "Spin",
		Duration: 30,
		IsFull:   false,
	}

	sut := NewActiveSessionSpecification()
	assert.False(t, sut.IsSatisfied(session))
}

func TestWithUpcomingUnAvailableSession(t *testing.T) {
	nowString := "2023-06-26 07:00:00"
	now, _ = time.Parse("2006-01-02 15:04:00", nowString)

	testedDateString := "2023-06-26 09:00:00"
	testedDate, _ := time.Parse("2006-01-02 15:04:00", testedDateString)

	session := &entity.Session{
		Id:       1,
		Time:     testedDate,
		Name:     "Spin",
		Duration: 30,
		IsFull:   true,
	}

	sut := NewActiveSessionSpecification()
	assert.False(t, sut.IsSatisfied(session))
}
