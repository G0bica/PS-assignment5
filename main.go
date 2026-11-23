package main

import (
	"context"
	"log"
	"os"

	"github.com/G0bica/PS-assignment5/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "Vremenske postaje",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Usage: "Najmanjša veljavna temperatura",
				Value: -50,
			},
			&cli.IntFlag{
				Name:  "maxTemp",
				Usage: "Največja veljavna temperatura",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Usage: "Najmanjše število meritev, potrebno za izračun statistike",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			var slovarPostaj map[string]vreme.Postaja

			slovarPostaj = make(map[string]vreme.Postaja)

			vreme.DodajPostajo("Dravlje", "Ljubljana", []float64{12.4, 14.5}, "LJ01", slovarPostaj)
			vreme.DodajPostajo("Bezigrad", "Ljubljana", []float64{11.3, 12.7, 14.3}, "LJ02", slovarPostaj)
			vreme.DodajPostajo("Ljudski vrt", "Maribor", []float64{13.3, 15.4, 13.5}, "MB01", slovarPostaj)
			vreme.DodajPostajo("Avtobusna postaja", "Kranj", []float64{12.9, 12.1, 12.5}, "KR01", slovarPostaj)
			vreme.DodajPostajo("Spik", "Julijske Alpe", []float64{1.3, 0.7, -3.0, -2.1}, "JA01", slovarPostaj)
			vreme.DodajPostajo("Triglav", "Julijske Alpe", []float64{1.3, 0.7}, "JA02", slovarPostaj)

			vreme.IzpisMeritev(slovarPostaj)
			vreme.IzpisPorocila(slovarPostaj, stMeritev, minTemp, maxTemp)

			vreme.DodajMeritev(slovarPostaj, "JA01", -5, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "JA02", -5, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "JA01", -500, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "JA01", 32100, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "ABCD", 32100, minTemp, maxTemp)
			vreme.DodajMeritev(slovarPostaj, "ABCD", 13, minTemp, maxTemp)

			vreme.IzpisMeritev(slovarPostaj)
			vreme.IzpisPorocila(slovarPostaj, stMeritev, minTemp, maxTemp)
			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
