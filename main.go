package main

import (
	"fmt"
	"log"
	"os"
	"path"

	maxminddb "github.com/oschwald/maxminddb-golang"
	"github.com/schollz/progressbar/v3"
)

var db string
var destDir string

func main() {
	getFlags()
	mmdb, err := maxminddb.Open(db)
	if err != nil {
		log.Fatal(err)
	}
	defer mmdb.Close()

	record := struct {
		Country struct {
			ISOCode string `maxminddb:"iso_code"`
		} `maxminddb:"country"`
	}{}

	if err = os.Mkdir(destDir, os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	isoToFile := make(map[string](*os.File))

	networks := mmdb.Networks()
	i := 0
	bar := progressbar.Default(-1)
	for networks.Next() {
		subnet, err := networks.Network(&record)
		if err != nil {
			log.Fatal(err)
		}
		country := record.Country.ISOCode
		var file = isoToFile[country]
		if file == nil {
			file, err = os.Create(path.Join(destDir, country+".txt"))
			if err != nil {
				log.Fatal(err)
			}
			isoToFile[country] = file
		}
		_, err = file.WriteString(subnet.String() + "\n")
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(1)
		i++
	}
	for _, file := range isoToFile {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("\nWrote %d records from %s\n", i, db)
}
