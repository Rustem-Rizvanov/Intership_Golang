package statistics

import (
    "database/sql"
    "time"
)

type StatisticsRepository struct {
    db *sql.DB
}

func NewStatisticsRepository(db *sql.DB) *StatisticsRepository {
    return &StatisticsRepository{db: db}
}

func (r *StatisticsRepository) Save(processingTime time.Duration, outcome string) error {
    _, err := r.db.Exec(`
        INSERT INTO request_statistics (processing_time, outcome, created_at) 
        VALUES ($1, $2, $3)`,
        processingTime.Seconds(), outcome, time.Now())
    return err
}

func (r *StatisticsRepository) GetTotalRequests() (int, error) {
    var count int
    err := r.db.QueryRow("SELECT COUNT(*) FROM request_statistics").Scan(&count)
    return count, err
}

func (r *StatisticsRepository) GetRequestsInLast24Hours() (int, error) {
    var count int
    err := r.db.QueryRow(`
        SELECT COUNT(*) FROM request_statistics 
        WHERE created_at > NOW() - INTERVAL '24 HOURS'`).Scan(&count)
    return count, err
}

func (r *StatisticsRepository) GetAvgProcessingTime() (float64, error) {
    var avg float64
    err := r.db.QueryRow("SELECT AVG(processing_time) FROM request_statistics").Scan(&avg)
    return avg, err
}
