package handlers

import (
	"context"
	"net/http"

	"github.com/saravase/golang_mux_swagger/plant-api/data"
)

func (plant *Plant) PlantValidationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		plantData := &data.Plant{}
		err := data.FromJSON(plantData, req.Body)
		if err != nil {
			plant.logger.Println("[ERROR] Deserializing plant", err)
			res.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, res)
			return
		}

		// validate the plant
		errs := plant.validation.Validate(plantData)
		if len(errs) != 0 {
			plant.logger.Println("[ERROR] Validating post", errs)
			res.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, res)
			return
		}

		// add the plant data into the context
		reqContext := context.WithValue(req.Context(), KeyPlant{}, plantData)
		req = req.WithContext(reqContext)

		// call the next handler, which can be another middlware or final handler
		next.ServeHTTP(res, req)

	})
}
