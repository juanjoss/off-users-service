package docs

import (
	"github.com/juanjoss/off-users-service/ports"
)

// swagger:route POST /addProductToSSD post-requests endpoints
// Adds a product to a user's device.
// responses:
//   200: addProductToSsdResponse
//	 default: addProductToSsdErrorResponse

// Product added to device successfully.
// swagger:response addProductToSsdResponse
type addProductToSsdResponseWrapper struct {
	// in:body
	Msg string
}

// Error adding product to device.
// swagger:response addProductToSsdErrorResponse
type addProductToSsdErrorResponseWrapper struct {
	// in:body
	Msg string
}

// swagger:parameters endpoints
type addProductToSsdRequestWrapper struct {
	// Send the user's device id, product barcode and quantity.
	// in:body
	Body ports.AddProductToSsdRequest
}
