package database

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"github.com/Gu1t4rist/NetPulse/internal/speedtest"
)

func LogResult(result *speedtest.Result) error {
	file, err := os.OpenFile("speedtest_results.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{
		time.Now().Format(time.RFC3339),
		strconv.FormatFloat(result.DownloadMbps, 'f', 2, 64),
		strconv.FormatFloat(result.UploadMbps, 'f', 2, 64),
		strconv.FormatFloat(result.PingMs, 'f', 2, 64),
	
	}

	return writer.Write(record)
}