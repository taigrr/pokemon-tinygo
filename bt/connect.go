package bt

import (
	"strings"
	"sync"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

var (
	advertising bool
	btInit      sync.Once
)

func Advertise() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Define the peripheral device info.
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "PokemonTinyGo",
	}))

	// Start advertising
	must("start adv", adv.Start())

	println("advertising...")
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
