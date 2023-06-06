package main

import (
	"flag"
	"fmt"
	aspace "github.com/nyudlts/go-aspace"
)

var (
	config      string
	environment string
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
}

func main() {
	flag.Parse()

	client, err := aspace.NewClient(config, environment, 20)
	if err != nil {
		panic(err)
	}

	tIds, err := client.GetResourceIDs(2)
	if err != nil {
		panic(err)
	}

	fIds, err := client.GetResourceIDs(3)
	if err != nil {
		panic(err)
	}

	uIds, err := client.GetResourceIDs(6)
	if err != nil {
		panic(err)
	}

	taIds, err := client.GetAccessionIDs(2)
	if err != nil {
		panic(err)
	}

	faIds, err := client.GetAccessionIDs(3)
	if err != nil {
		panic(err)
	}

	uaIds, err := client.GetAccessionIDs(6)
	if err != nil {
		panic(err)
	}

	tdIds, err := client.GetDigitalObjectIDs(2)
	if err != nil {
		panic(err)
	}

	fdIds, err := client.GetDigitalObjectIDs(3)
	if err != nil {
		panic(err)
	}

	udIds, err := client.GetDigitalObjectIDs(6)
	if err != nil {
		panic(err)
	}

	fmt.Println("Archivesspace Object Counts\nResources:")
	fmt.Printf("  TamWag:\t%d\n  Fales:\t%d\n  Archives:\t%d\n", len(tIds), len(fIds), len(uIds))

	fmt.Println("\nAccessions:")
	fmt.Printf("  TamWag:\t%d\n  Fales:\t%d\n  Archives:\t%d\n", len(taIds), len(faIds), len(uaIds))

	fmt.Println("\nDigital Objects:")
	fmt.Printf("  TamWag:\t%d\n  Fales:\t%d\n  Archives:\t%d\n", len(tdIds), len(fdIds), len(udIds))
}
