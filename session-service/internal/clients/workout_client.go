package clients

import (
	"fmt"
	"net/http"
)

const WorkoutServiceURL = "http://workout-service:8082/api/v1/workouts/"

func VerifyWorkoutExists(workoutID uint) bool {
	url := fmt.Sprintf("%s%d", WorkoutServiceURL, workoutID)
	resp, err := HttpClient.Get(url)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}
