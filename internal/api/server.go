package api

import (
	"encoding/csv"
	"net/http"
	"os"
)

func StartServer() {
	http.HandleFunc("/results", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("speedtest_log.csv")
		if err != nil {
			http.Error(w, "Could not open log file", 500)
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Could not read log file", 500)
			return
		}

		for _, record := range records {
			w.Write([]byte(record[0] + " - Download: " + record[1] + " Mbps, Upload: " + record[2] + " Mbps, Ping: " + record[3] + " ms\n"))
		}
	})

	http.ListenAndServe(":8000", nil)
}