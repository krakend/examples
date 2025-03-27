package main

import (
	libpb "github.com/krakend/examples/grpc/images/grpc/genlibs/lib"
)

var airports = []libpb.Location{
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "UK",
			City:        "London",
			AddressLine: "Heathrow Airport",
		},
		Position: &libpb.GeoPosition{
			Latitude:  51.4734,
			Longitude: -0.48899,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "DE",
			City:        "Frankfurt",
			AddressLine: "Frankfurt International Airport",
		},
		Position: &libpb.GeoPosition{
			Latitude:  50.05029,
			Longitude: 8.5663,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "IT",
			City:        "Rome",
			AddressLine: "Fiumicino Airport",
		},
		Position: &libpb.GeoPosition{
			Latitude:  41.8033,
			Longitude: 12.2503,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "ES",
			City:        "Barcelona",
			AddressLine: "Barcelona-El Prat",
		},
		Position: &libpb.GeoPosition{
			Latitude:  41.2975,
			Longitude: 2.083,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "FR",
			City:        "Paris",
			AddressLine: "Charles de Gaulle",
		},
		Position: &libpb.GeoPosition{
			Latitude:  49.0096,
			Longitude: 2.55166,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "DE",
			City:        "Munich",
			AddressLine: "Franz Josef Strauss",
		},
		Position: &libpb.GeoPosition{
			Latitude:  4635087,
			Longitude: 11.77,
		},
	},
}
