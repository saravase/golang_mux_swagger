package data

// A list of plants returns in the response
// swagger:response plantsResponse
type plantsResponseWrapper struct {
	// All plants in the system
	// in: body
	Body []Plant
}

// Single plant data return in the response
// swagger:response plantResponse
type plantResponseWrapper struct {
	// Single plant in the system
	// in: body
	Body Plant
}

// swagger:parameters addPlant updatePlant
type plantParameterWrapper struct {
	// Plant data properties from rrequest body
	// Note: id field is ignored both add and update operation
	// in: body
	// required: true
	Body Plant
}
