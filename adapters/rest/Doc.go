// Package Doc Craft challenge API
//
// Documentation for Craft challenge
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//  License: MIT http://opensource.org/licenses/MIT
//  Contact: Julien Rollin<linrol.news@gmail.com> https://www.julienrollin.com
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package rest


// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}


// Generic error message returned as a string
// swagger:response genericErrorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response validationErrorResponse
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}



// No content is returned by this API endpoint
// swagger:response notFoundResponse
type notFoundResponseWrapper struct {
}


// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}




// An swagger route parameter model.
//
// This is used for operations that want the code of game in the path
// swagger:parameters gameId playerListId
type GameCode struct {
	// The code of the game
	//
	// in: path
	// required: true
	Code string `json:"code"`
}

