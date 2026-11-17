package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Rojsn/PS-vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "PS-vreme",
		Usage: "Vremenske meritve",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name:  "minTemp",
				Value: -50.0,
				Usage: "Najnižja dovoljena temperatura meritve",
			},
			&cli.FloatFlag{
				Name:  "maxTemp",
				Value: 50.0,
				Usage: "Najvišja dovoljena temperatura meritve",
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
				Usage: "Najmanjše število meritev, potrebnih za izračun statistike",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			minT := c.Float("minTemp")
			maxT := c.Float("maxTemp")
			minM := int(c.Int("stMeritev"))

			postaje := make(map[string]vreme.Postaja)

			vreme.DodajPostajo(postaje, "KP", "Koper Luka", "Obalna")
			vreme.DodajPostajo(postaje, "KR", "Kredarica", "Gorenjska")
			vreme.DodajPostajo(postaje, "LJ", "Ljubljana Bežigrad", "Osrednja")
			vreme.DodajPostajo(postaje, "MB", "Maribor Tabor", "Štajerska")

			vreme.DodajMeritev(postaje, "KP", 24.5, minT, maxT)
			vreme.DodajMeritev(postaje, "KP", 27.0, minT, maxT)
			vreme.DodajMeritev(postaje, "KP", 28.4, minT, maxT)

			vreme.DodajMeritev(postaje, "KR", -1.5, minT, maxT)
			vreme.DodajMeritev(postaje, "KR", 0.5, minT, maxT)
			vreme.DodajMeritev(postaje, "KR", 3.5, minT, maxT)

			vreme.DodajMeritev(postaje, "LJ", 16.5, minT, maxT)
			vreme.DodajMeritev(postaje, "LJ", 21.5, minT, maxT)

			vreme.DodajMeritev(postaje, "MB", 12.0, minT, maxT)
			vreme.DodajMeritev(postaje, "MB", 15.0, minT, maxT)
			vreme.DodajMeritev(postaje, "MB", 13.5, minT, maxT)

			fmt.Println("--- DEMONSTRACIJA DODAJANJA MERITEV ---")
			vreme.DodajMeritev(postaje, "LJ", 19.0, minT, maxT)
			vreme.DodajMeritev(postaje, "KP", 55.0, minT, maxT)
			vreme.DodajMeritev(postaje, "NG", 20.0, minT, maxT)

			fmt.Println()
			vreme.IzpisMeritev(postaje)
			vreme.IzpisPorocila(postaje, minM)

			return nil
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
