package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Device struct {
	Path   string
	Id     int
	Status string
}

type DeviceSet struct {
	Id      int64
	Name    string
	Devices []*Device
}

func main() {
	input := strings.NewReader(
		`0050E0 device1 /dev/sdd 000001 Available
               /dev/sde 000002 Available
			   /dev/sdc 000003 Occupied
0040E1 device2 /dev/sdf 000004 Available`)

	scanner := bufio.NewScanner(input)

	result := make(map[int64]*DeviceSet, 2)
	var lastKey int64
	for scanner.Scan() {
		ds := &DeviceSet{Devices: make([]*Device, 0)}
		d := &Device{}
		var hex string
		if _, err := fmt.Sscanf(scanner.Text(), "%s %s %s %d %s", &hex, &ds.Name, &d.Path, &d.Id, &d.Status); err != nil {
			if _, err := fmt.Sscanf(scanner.Text(), "      %s %d %s", &d.Path, &d.Id, &d.Status); err != nil {
				log.Fatal(err)
			} else {
				result[lastKey].Devices = append(result[lastKey].Devices, d)
			}
		} else {
			ds.Id, err = strconv.ParseInt(hex, 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			lastKey = ds.Id
			ds.Devices = append(ds.Devices, d)
			result[ds.Id] = ds
		}
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 2, 1, ' ', 0)

	for _, value := range result {
		fmt.Fprintf(w, "%d\t%s\t%s\t%d\t%s\n", value.Id, value.Name, value.Devices[0].Path, value.Devices[0].Id, value.Devices[0].Status)
		for _, device := range value.Devices[1:] {
			fmt.Fprintf(w, "\t\t%s\t%d\t%s\n", device.Path, device.Id, device.Status)
		}
	}

	w.Flush()
}
