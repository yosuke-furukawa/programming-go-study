package main

import (
	"log"
	"os"

	"fmt"

	"github.com/yosuke-furukawa/programming-go-study/ch04/ex11/src/gh"
)

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	owner := os.Args[1]
	repository := os.Args[2]
	repo := gh.Repository{
		owner,
		repository,
		token,
	}

	method := os.Args[3]
	title := os.Args[4]
	body := os.Args[5]

	switch method {
	case "create":
		result, err := repo.Create(gh.Issue{title, body, ""})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(result)
		fmt.Printf("created. see: %s\n", result.Url)
	}
}
