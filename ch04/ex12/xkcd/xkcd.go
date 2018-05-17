package xkcd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const XKCD_URL = "https://xkcd.com/%d/info.0.json"
const MAX_ID = 100

type Xkcd struct {
	Transcript string `json:"transcript"`
	Img        string `json:"img"`
}

func Download(filepath string) error {
	var xkcds []Xkcd
	for i := 1; i < MAX_ID; i++ {
		fmt.Println("start ", i)
		url := fmt.Sprintf(XKCD_URL, i)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var xkcd Xkcd
		if err := json.NewDecoder(resp.Body).Decode(&xkcd); err != nil {
			return err
		}
		log.Println(xkcd)
		xkcds = append(xkcds, xkcd)
		fmt.Println("done ", i)
	}
	data, err := json.Marshal(xkcds)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filepath, data, 0644)

}
