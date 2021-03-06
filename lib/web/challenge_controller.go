package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/superboum/moolinet/lib/judge"
)

// ChallengeController is a controller used to display the public list of challenges.
type ChallengeController struct {
	judge *judge.Judge
}

// NewChallengeController returns a ChallengeController with provided Judger.
func NewChallengeController(j *judge.Judge) *ChallengeController {
	c := new(ChallengeController)
	c.judge = j
	return c
}

// ServeHTTP writes the public list of challenges to the result.
func (c *ChallengeController) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(res)
	err := encoder.Encode(c.judge.PublicChallenges)
	if err != nil {
		log.Println("There was an error while encoding the response: " + err.Error())
	}
}
