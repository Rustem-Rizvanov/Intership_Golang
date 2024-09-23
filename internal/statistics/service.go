package statistics

import (
    "time"
)

type RequestStatistics struct {
    TotalRequests     int     `json:"total_requests"`
    Last24h           int     `json:"last_24h"`
    AvgProcessingTime float64 `json:"avg_processing_time"`
}

type StatisticsService struct {
    repo *StatisticsRepository
}

func NewStatisticsService(repo *StatisticsRepository) *StatisticsService {
    return &StatisticsService{repo: repo}
}

func (s *StatisticsService) SaveRequest(processingTime time.Duration, outcome string) error {
    return s.repo.Save(processingTime, outcome)
}

func (s *StatisticsService) GetStatistics() (RequestStatistics, error) {
    totalRequests, err := s.repo.GetTotalRequests()
    if err != nil {
        return RequestStatistics{}, err
    }

    last24h, err := s.repo.GetRequestsInLast24Hours()
    if err != nil {
        return RequestStatistics{}, err
    }

    avgProcessingTime, err := s.repo.GetAvgProcessingTime()
    if err != nil {
        return RequestStatistics{}, err
    }

    return RequestStatistics{
        TotalRequests:     totalRequests,
        Last24h:           last24h,
        AvgProcessingTime: avgProcessingTime,
    }, nil
}
