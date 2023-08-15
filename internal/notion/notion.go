package notion

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	databaseCache    DatabaseResult
	databaseCacheSet time.Time
)

const CACHE_TIMEOUT = 1 * time.Minute

func FetchDatabase() (DatabaseResult, error) {
	if time.Since(databaseCacheSet) < CACHE_TIMEOUT {
		return databaseCache, nil
	}

	db_id := os.Getenv("NOTION_DB_ID")
	url := "https://api.notion.com/v1/databases/" + db_id + "/query"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := client.Do(req)
	if err != nil {
		return DatabaseResult{}, err
	}
	defer res.Body.Close()

	var j DatabaseResult
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return DatabaseResult{}, err
	}

	databaseCache = j
	databaseCacheSet = time.Now()

	return j, nil
}

func CreatePage(body string) error {
	url := "https://api.notion.com/v1/pages"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	_, err := client.Do(req)
	if err != nil {
		return err
	}
	return nil
}

func FetchPage(id string) (Page, error) {
	url := "https://api.notion.com/v1/pages/" + id
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := client.Do(req)
	if err != nil {
		return Page{}, err
	}
	defer res.Body.Close()

	var j PageResult
	err = json.NewDecoder(res.Body).Decode(&j)
	if err != nil {
		return Page{}, err
	}

	return ConvertPageResult(j), nil
}

func UpdatePage(id string, body string) error {
	url := "https://api.notion.com/v1/pages/" + id
	client := &http.Client{}
	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Authorization", "Bearer "+os.Getenv("NOTION_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("notion " + fmt.Sprint(res.StatusCode))
	}

	return nil
}
