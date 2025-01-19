package main

import (
	"fmt"
	"time"
)

func main() {
	h := &hoge{
		updatedAt: time.Date(2021, 6, 24, 0, 0, 0, 0, time.UTC),
	}
	fmt.Println(getLatestUpdatedAt(h, nil))
}

type hoge struct {
	updatedAt time.Time
}

func (h *hoge) GetUpdatedAt() time.Time {
	return h.updatedAt
}

type updatedAt interface {
	GetUpdatedAt() time.Time
}

func getLatestUpdatedAt(updatedAt ...updatedAt) time.Time {
	if len(updatedAt) == 0 {
		return time.Time{}
	}
	var latest time.Time
	for _, u := range updatedAt {
		if u == nil {
			continue
		}
		if u.GetUpdatedAt().After(latest) {
			latest = u.GetUpdatedAt()
		}
	}
	return latest
}
