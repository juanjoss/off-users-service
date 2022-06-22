package docs

import "github.com/juanjoss/off-users-service/model"

// swagger:route GET /users/ssds/random get-requests endpoint
// Returns a random user's device.
// responses:
//   200: randomSsdResponse
//	 default: randomSsdErrorResponse

// Random device request success.
// swagger:response randomSsdResponse
type randomSsdResponseWrapper struct {
	// in:body
	SSD model.SSD
}

// Random device request error.
// swagger:response randomSsdErrorResponse
type randomSsdErrorResponseWrapper struct {
	// in:body
	Msg string
}
