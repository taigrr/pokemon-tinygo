package bt

import (
	"strings"
	"sync"
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

var (
	advertising   bool
	btInit        sync.Once
	adv           *bluetooth.Advertisement
	lastBLEChange time.Time
)

func Advertise() {
	if time.Since(lastBLEChange) < 5*time.Second {
		// Don't allow rapid toggling of BLE advertising, which can cause a panic
		return
	}
	lastBLEChange = time.Now()

	switch advertising {
	case false:
		// Enable BLE interface.
		btInit.Do(func() {
			must("enable BLE stack", adapter.Enable())
			adv = adapter.DefaultAdvertisement()
			must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
				LocalName: "x.com/warptux",
			}))
		})
		// Define the peripheral device info.
		must("start adv", adv.Start())
		println("advertising...")
		advertising = true
	case true:
		// Stop advertising
		must("stop adv", adv.Stop())
		println("no longer advertising...")
		advertising = false
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}

func Scan() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())
	ch := make(chan bluetooth.ScanResult, 1)
	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, result bluetooth.ScanResult) {
		println("found device:", result.Address.String(), result.RSSI, result.LocalName())
		if strings.Contains(result.LocalName(), "") {
			if result.AdvertisementPayload.LocalName() == "pokemontinygo" {
				adapter.StopScan()
				ch <- result
			}
		}
	})
	must("start scan", err)
}
