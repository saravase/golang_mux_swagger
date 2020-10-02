package handlers

// Request completed successfully
// swagger:response successContent
type successContentWrapper struct{}

// No content return from server
// swagger:response noContent
type noContentWrapper struct{}

// Generic error return from server
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Generic error return from server
	// in: body
	Body GenericError
}

// Validation error return from plant validator
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Validation error return from plant validator
	// in: body
	Body ValidationError
}

// swagger:parameters deletePlant singlePlant updatePlant
type plantIDParameterWrapper struct {
	// Fetch the plant id from request URL path
	// in: path
	// required: true
	ID int `json:"id"`
}
