package handlers

import (
	"applicationDesignTest/models"
	"applicationDesignTest/utils"
	"encoding/json"
	"net/http"
	"time"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder models.Order
	err := json.NewDecoder(r.Body).Decode(&newOrder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		utils.LogErrorf("Failed to decode request body: %s", err)
		return
	}

	daysToBook := utils.DaysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range models.Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			models.Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		utils.LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	models.Orders = append(models.Orders, newOrder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	utils.LogInfo("Order successfully created: %v", newOrder)
}
