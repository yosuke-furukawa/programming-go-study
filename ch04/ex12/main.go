package main

import (
	"encoding/json"
	"io/ioutil"

	"os"

	"fmt"
	"strings"

	"github.com/yosuke-furukawa/programming-go-study/ch04/ex12/xkcd"
)

var xkcds []xkcd.Xkcd

func init() {
	data, _ := ioutil.ReadFile("./files/all.json")
	json.Unmarshal(data, &xkcds)
}

func main() {
	//err := xkcd.Download("./files/all.json")
	//if err != nil {
	//	fmt.Errorf("download error %v", err)
	//}

	query := os.Args[1]

	for _, xkcd := range xkcds {
		if strings.Contains(xkcd.Transcript, query) {
			fmt.Println(xkcd.Img)
		}
	}

}
