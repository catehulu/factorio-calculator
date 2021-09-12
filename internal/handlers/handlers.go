package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/catehulu/factorio-calculator/internal/factorio_search"
	"github.com/catehulu/factorio-calculator/internal/models"
	"github.com/catehulu/factorio-calculator/internal/render"
)

func Index(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "index.html", &models.TemplateData{
		Recipe: models.GetAllRecipe(),
	})
}

func Calculate(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		payload := struct {
			Name string  `json:"name"`
			Qty  float32 `json:"qty"`
		}{}

		if err := decoder.Decode(&payload); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if payload.Qty < 1 {
			http.Error(rw, "Zero quantity", http.StatusBadRequest)
			return
		}

		recipe, err := models.GetRecipe(payload.Name)
		if err != nil {
			http.Error(rw, "Item not found!", http.StatusBadRequest)
			return
		}

		sr := make(factorio_search.SearchResult)
		factorio_search.CalculateProduction(recipe, sr, payload.Qty, 1)

		jsonInBytes, err := json.Marshal(sr)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		rw.Write(jsonInBytes)
		return
	}

	http.Error(rw, "Only accept POST request", http.StatusBadRequest)
}
