package handler

import "telegram-bot/internal/models"

func (h *Handler) IdentifyParams(filter *models.Filter) (map[string]string, error) {
	//todo create logic for identifying params by mood and taste
	params := make(map[string]string, 3)
	for _, item := range filter.Attr {
		switch item {
		case "fine":
			params["abv_gt"] = "3.0"
		case "sad":
			params["abv_gt"] = "6.0"
		case "party":
			params["abv_gt"] = "2.0"
		case "bitter":
			params["food"] = filter.Attr[1]
		case "sweet":
			params["food"] = filter.Attr[1]
		case "neutral":
			params["food"] = filter.Attr[1]
		case "spicy":
			params["food"] = filter.Attr[1]
		}
	}
	return params, nil
}
