package Requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(url string) error {
	fmt.Println("Fetching data from:", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error: %s", res.Status)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &JsonMapData)
	if err != nil {
		return err
	}
	return nil

}
