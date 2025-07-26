package api

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"
)

func StartServer() {
	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("speedtest_results.csv")
		if err != nil {
			http.Error(w, "Could not open results file", 500)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Could not read results file", 500)
			return
		}

		type Result struct {
			Timestamp string  `json:"timestamp"`
			Download  float64 `json:"download"`
			Upload    float64 `json:"upload"`
			Ping      float64 `json:"ping"`
		}

		n := 50
		if nStr := r.URL.Query().Get("n"); nStr != "" {
			if nParsed, err := strconv.Atoi(nStr); err == nil && nParsed > 0 {
				n = nParsed
			}
		}

		results := []Result{}
		start := 0
		if len(records) > n {
			start = len(records) - n
		}
		for _, record := range records[start:] {
			if len(record) < 4 {
				continue // skip malformed lines
			}
			download, _ := strconv.ParseFloat(record[1], 64)
			upload, _ := strconv.ParseFloat(record[2], 64)
			ping, _ := strconv.ParseFloat(record[3], 64)
			results = append(results, Result{
				Timestamp: record[0],
				Download:  download,
				Upload:    upload,
				Ping:      ping,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	http.HandleFunc("/stats/weekly", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("speedtest_results.csv")
		if err != nil {
			http.Error(w, "Could not open results file", 500)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Could not open results file", 500)
			return
		}

		type Stat struct {
			AvgDownload float64 `json:"avg_download"`
			MinDownload float64 `json:"min_download"`
			MaxDownload float64 `json:"max_download"`
			Count       int     `json:"count"`
		}

		var downloads []float64
		now := time.Now()
		oneWeekAgo := now.AddDate(0, 0, -7)

		for _, record := range records {
			if len(record) < 4 {
				continue
			}
			t, err := time.Parse(time.RFC3339, record[0])
			if err != nil {
				continue
			}
			if t.Before(oneWeekAgo) {
				continue
			}
			download, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				continue
			}
			downloads = append(downloads, download)
		}

		var avg, min, max float64
		count := len(downloads)
		if count > 0 {
			min = downloads[0]
			max = downloads[0]
			sum := 0.0
			for _, d := range downloads {
				sum += d
				if d < min {
					min = d
				}
				if d > max {
					max = d
				}
			}
			avg = sum / float64(count)
		}

		stat := Stat{
			AvgDownload: avg,
			MinDownload: min,
			MaxDownload: max,
			Count:       count,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(stat)
	})

	http.ListenAndServe(":8000", nil)
}
