package main

import (
	"log"
	"os"

	"fmt"

	"sort"

	"time"

	"github.com/yosuke-furukawa/programming-go-study/ch04/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})
	var oneMonthFound bool
	var oneYearFound bool
	var overYearFound bool
	for _, item := range result.Items {
		oneMonth := time.Now().AddDate(0, -1, 0)
		oneYear := time.Now().AddDate(-1, 0, 0)
		if item.CreatedAt.After(oneMonth) {
			if !oneMonthFound {
				oneMonthFound = true
				fmt.Printf("--------- recent from %.10s ------------\n", oneMonth)
			}
			fmt.Printf("#%-5d %9.9s %.55s %.10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
			continue
		}
		if item.CreatedAt.After(oneYear) {
			if !oneYearFound {
				oneYearFound = true
				fmt.Printf("--------- last 1 year from %.10s ------------\n", oneYear)
			}
			fmt.Printf("#%-5d %9.9s %.55s %.10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
			continue
		}
		if !overYearFound {
			fmt.Println("--------- 1 year later ------------")
			overYearFound = true
		}
		fmt.Printf("#%-5d %9.9s %.55s %.10s\n", item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}
