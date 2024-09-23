package statistics

import (
    "encoding/json"
    "net/http"
)

type StatisticsController struct {
    service *StatisticsService
}

func NewStatisticsController(service *StatisticsService) *StatisticsController {
    return &StatisticsController{service: service}
}

func (c *StatisticsController) GetStatisticsHandler(w http.ResponseWriter, r *http.Request) {
    stats, err := c.service.GetStatistics()
    if err != nil {
        http.Error(w, "Unable to fetch statistics", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}
