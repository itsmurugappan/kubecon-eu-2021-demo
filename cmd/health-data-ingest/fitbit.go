package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/itsmurugappan/kubecon/pkg/ce"
	"github.com/itsmurugappan/kubecon/pkg/types"
)

func triggerFBEvents() {
	c := ce.CEClient()
	i := 1
	if err := filepath.Walk(os.Getenv("KO_DATA_PATH")+"/fitbit", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		e := ce.GetCE(types.Fitbit_Type, "https://api.fitbit.com")
		e.SetData("application/json", data)
		fmt.Printf("data %s\n", string(data))
		ext, _ := json.Marshal(types.Meta{os.Getenv("ACT_DATE"), fmt.Sprintf("fbuser-%d", i)})
		e.SetExtension(types.Meta_Extension, string(ext))
		fmt.Printf("e %v\n", e)
		c.SendMessage(e)
		i++
		return nil
	}); err != nil {
		log.Printf("err in fb ingestion : %s", err.Error())
	}
}
