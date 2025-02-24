package main

import (
	libpb "github.com/krakendio/playground-enterprise/images/grpc/genlib/lib"
)

var trainStations = []libpb.Location{
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "FR",
			City:        "Paris",
			AddressLine: "Gare du Lion",
		},
		Position: &libpb.GeoPosition{
			Latitude:  48.84488,
			Longitude: -2.37404,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "ES",
			City:        "Madrid",
			AddressLine: "Atocha",
		},
		Position: &libpb.GeoPosition{
			Latitude:  40.4066,
			Longitude: -3.6986,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "BE",
			City:        "Liege",
			AddressLine: "Liege-Guillemins",
		},
		Position: &libpb.GeoPosition{
			Latitude:  50.62445,
			Longitude: -5.56651,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "PT",
			City:        "Porto",
			AddressLine: "San Bento",
		},
		Position: &libpb.GeoPosition{
			Latitude:  41.14555,
			Longitude: -8.61034,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "FR",
			City:        "Strasbourg",
			AddressLine: "Gare de Strasbourg",
		},
		Position: &libpb.GeoPosition{
			Latitude:  48.57491,
			Longitude: 7.79502,
		},
	},
	libpb.Location{
		Address: &libpb.Address{
			CountryCode: "DE",
			City:        "Berlin",
			AddressLine: "BERLIN HAUPTBAHNHOF",
		},
		Position: &libpb.GeoPosition{
			Latitude:  52.52497,
			Longitude: 13.3694,
		},
	},
}
