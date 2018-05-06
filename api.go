package tracker

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetProfile is used to request a profile from fortnite tracker
func GetProfile(platform, name, APIToken string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", ProfileRoute+platform+"/"+name, nil)
	if err != nil {
		log.Fatal("failed to create profile request", err)
	}

	req.Header.Add(headerKeyName, APIToken)

	fmt.Println(req.URL)
	fmt.Println(req.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("failed to get profile from server", err)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body //TODO process json
}
