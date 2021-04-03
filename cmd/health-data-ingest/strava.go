package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/itsmurugappan/kubecon/pkg/ce"
	"github.com/itsmurugappan/kubecon/pkg/types"
)

func triggerStravaEvents() {
	c := ce.CEClient()
	i := 1
	if err := filepath.Walk(os.Getenv("KO_DATA_PATH")+"/strava", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		e := ce.GetCE(types.Strava_Type, "https://www.strava.com")
		e.SetData("application/json", data)
		fmt.Printf("data %s\n", string(data))
		fmt.Printf("e %v\n", e)
		c.SendMessage(e)
		i++
		return nil
	}); err != nil {
		log.Printf("err in fb ingestion : %s", err.Error())
	}
}
