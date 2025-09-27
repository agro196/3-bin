package main

import (
	"time"
)

type bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type binList []bin

func newBin(id string, private bool, name string) bin {
	return bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func main() {

}
