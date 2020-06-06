package query

import "github.com/shurcooL/graphql"

/*
EpicStoreFreeGames is a Query for
getting the free games off the Epic Store
*/
var EpicStoreFreeGames struct {
	Catalog struct {
		SearchStore struct {
			Elements []struct {
				Title       graphql.String
				Description graphql.String
				KeyImages   []struct {
					Type graphql.String
					URL  graphql.String
				}
				Promotions struct {
					PromotionalOffers []struct {
						PromotionalOffers []struct {
							EndDate graphql.String
						}
					}
				} `graphql:"promotions(category: \"freegames\")"`
			}
		} `graphql:"searchStore(allowCountries: \"US\", category: \"freegames\", count: 1000, country: \"US\", locale: \"en-US\", sortBy: \"effectiveDate\", sortDir: \"asc\")"`
	} `graphql:"Catalog"`
}
