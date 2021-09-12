package factorio_search

import (
	"github.com/catehulu/factorio-calculator/internal/models"
)

type ItemOutput struct {
	Name      string
	Tier      int
	PerSecond float32
	PerMinute float32
}

type SearchResult map[string]*ItemOutput

func CalculateProduction(r models.Recipe, sr SearchResult, q float32, tier int) {
	/** Quantity per seconds*/

	i := ItemOutput{
		Name:      r.Name,
		Tier:      tier,
		PerSecond: q,
		PerMinute: q * 60,
	}

	/** Check if already exists, and sum quantity also tier */
	if val, ok := sr[i.Name]; ok {
		if val.Tier < i.Tier {
			val.Tier = i.Tier
		}

		val.PerSecond += i.PerSecond
		val.PerMinute = val.PerSecond * 60
	} else {
		sr[i.Name] = &i
	}

	/** Recursive, call next recipe */
	for _, ingredients := range r.Ingredients {
		nr, err := models.GetRecipe(ingredients.Name)
		if err == nil {
			CalculateProduction(nr, sr, ingredients.Amount*q, tier+1)
		}
	}
}

func FormatValue(sr SearchResult) []ItemOutput {
	arr := []ItemOutput{}
	for _, v := range sr {
		arr = append(arr, *v)
	}
	return arr
}
