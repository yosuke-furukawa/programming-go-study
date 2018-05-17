package gh

import (
	"os"
	"strconv"
	"testing"
)

func TestRepository_CreateReadEditClose(t *testing.T) {
	token := os.Getenv("GITHUB_TOKEN")
	repo := &Repository{
		"yosuke-furukawa",
		"test",
		token,
	}

	result, err := repo.Create(Issue{
		"test",
		"test",
		"",
	})

	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if result.Number == "" {
		t.Error("result is empty")
	}
	number, err := strconv.Atoi(string(result.Number))
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	result, err = repo.Read(number)

	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if result.Title != "test" {
		t.Error("result is not test")
	}
	result, err = repo.Edit(number, Issue{
		"test123",
		"test123",
		"",
	})
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if result.Title != "test123" {
		t.Error("result is not test123")
	}

	result, err = repo.Close(number)
	if err != nil {
		t.Errorf("error is thrown %v", err)
	}

	if result.State != "closed" {
		t.Error("result is not closed", result)
	}

}
