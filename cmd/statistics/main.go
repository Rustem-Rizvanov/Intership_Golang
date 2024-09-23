package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type StatsHandler struct {
    db *sql.DB
}

func (h *StatsHandler) GetStats(w http.ResponseWriter, r *http.Request) {

    
	rows, err := h.db.Query("SELECT COUNT(*), AVG(processing_time) FROM request_statistics WHERE created_at > NOW() - INTERVAL '1 DAY'")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var count int
    var avgTime float64
    if rows.Next() {
        err := rows.Scan(&count, &avgTime)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }

    // Форматирование ответа
    response := map[string]interface{}{
        "count":      count,
        "average_time": avgTime,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    db, err := sql.Open("postgres", os.Getenv("STATISTICS_POSTGRES_URL"))
    if err != nil {
        log.Fatalf("Error connecting to the database: %v", err)
    }
    defer db.Close()

    router := mux.NewRouter()
    statsHandler := &StatsHandler{db: db}
    
    router.HandleFunc("/stats", statsHandler.GetStats).Methods("GET")

    log.Println("Starting statistics service on :5462")
    if err := http.ListenAndServe(":5462", router); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
