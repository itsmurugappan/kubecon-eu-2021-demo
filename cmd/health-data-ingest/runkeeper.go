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

func triggerRKEvents() {
	c := ce.CEClient()
	i := 1
	if err := filepath.Walk(os.Getenv("KO_DATA_PATH")+"/runkeeper", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		e := ce.GetCE(types.Runkeeper_Type, "https://www.runkeeper.com")
		e.SetData("application/json", types.Runkeeper{string(data)})
		e.SetExtension(types.Meta_Extension, fmt.Sprintf("runkeeper-%d", i))
		fmt.Printf("e %v\n", e)
		c.SendMessage(e)
		i++
		return nil
	}); err != nil {
		log.Printf("err in fb ingestion : %s", err.Error())
	}
}
