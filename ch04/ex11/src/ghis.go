package main

import (
	"log"
	"math/rand"
	"os"

	"fmt"

	"os/exec"

	"io/ioutil"

	"flag"

	"time"

	"github.com/yosuke-furukawa/programming-go-study/ch04/ex11/src/gh"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func runEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	path := os.TempDir() + "/tmp/github" + RandStringRunes(10)

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error %v", err)
		os.Exit(1)
	}
	return string(b)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Usage of %s:
  Note: GITHUB_TOKEN environment variable is required.
  Example: GITHUB_TOKEN=<TOKEN> %s --owner yosuke-furukawa --repository test --id 1
  %s [OPTIONS] ARGS...
Options\n
`, os.Args[0], os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
	token := os.Getenv("GITHUB_TOKEN")

	if token == "" {
		log.Fatal("no token!!!")
		flag.Usage()
		os.Exit(1)
	}
	var (
		owner      = flag.String("owner", "", "repository owner/organization [required]")
		repository = flag.String("repository", "", "repository name [required]")
		op         = flag.String("op", "read", "issue operation")
		title      = flag.String("title", "", "issue title [required / create | edit]")
		body       = flag.String("body", "", "issue body [optional]")
		id         = flag.Int("id", 0, "issue id [required / edit | read | close]")
	)
	flag.Parse()

	if *owner == "" || *repository == "" {
		log.Fatal("no owner or repository!!!")
		flag.Usage()
		os.Exit(1)
	}
	repo := gh.Repository{
		Owner: *owner,
		Repo:  *repository,
		Token: token,
	}
	log.Println(repo)

	switch *op {
	case "create":
		if *body == "" {
			*body = runEditor()
		}
		result, err := repo.Create(gh.Issue{Title: *title, Body: *body, State: ""})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(result)
		fmt.Printf("created. see: %s\n", result.Url)
	case "read":
		result, err := repo.Read(*id)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(result)
		fmt.Printf("read. see: %s\n", result.Url)
		fmt.Printf("read. title: %s\n", result.Title)
		fmt.Printf("read. body: %s\n", result.Body)
	case "edit":
		if *body == "" {
			*body = runEditor()
		}
		log.Println(*body)
		result, err := repo.Edit(*id, gh.Issue{Title: *title, Body: *body, State: ""})
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(result)
		fmt.Printf("edited. see: %s\n", result.Url)
		fmt.Printf("edited. title: %s\n", result.Title)
		fmt.Printf("edited. body: %s\n", result.Body)
	case "close":
		result, err := repo.Close(*id)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		log.Println(result)
		fmt.Printf("closed. see: %s\n", result.Url)
		fmt.Printf("closed. title: %s\n", result.Title)
		fmt.Printf("closed. body: %s\n", result.Body)
		fmt.Printf("closed. state: %s\n", result.State)
	}
}
