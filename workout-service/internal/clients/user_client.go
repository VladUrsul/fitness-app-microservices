package clients

import (
	"fmt"
	"net/http"
)

const UserServiceURL = "http://user-service:8081/api/v1/users/"

func VerifyUserExists(userID uint) bool {
	url := fmt.Sprintf("%s%d", UserServiceURL, userID)
	resp, err := HttpClient.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}
