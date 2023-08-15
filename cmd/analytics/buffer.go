package main

import (
	"fmt"
	"log"
	"time"

	analytics "github.com/ivynya/analytics/pkg"
)

type Buffer map[string]BufferData
type BufferData struct {
	Visits       int
	Interactions int
}

var buffer = make(Buffer)

// Queues update to campaign with nID and visits/interactions
func bufferData(nID string, visits int, interactions int) {
	if visits <= 0 && interactions <= 0 {
		return
	}

	if _, ok := buffer[nID]; ok {
		buffer[nID] = BufferData{
			Visits:       buffer[nID].Visits + visits,
			Interactions: buffer[nID].Interactions + interactions,
		}
	} else {
		buffer[nID] = BufferData{
			Visits:       visits,
			Interactions: interactions,
		}
	}
}

// Goroutine that pushes updates to Notion every 10 seconds
func bufferFlushLoop() {
	for {
		time.Sleep(10 * time.Second)
		if len(buffer) > 0 {
			log.Println("[UPD] " + fmt.Sprint(len(buffer)) + " campaigns")
		}

		for nID, data := range buffer {
			err := analytics.UpdateVisits(nID, data.Visits, false)
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			err = analytics.UpdateInteractions(nID, data.Interactions)
			if err != nil {
				log.Println("[ERR] " + err.Error())
			}
			delete(buffer, nID)
		}
	}
}
