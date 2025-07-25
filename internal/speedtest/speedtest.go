package speedtest

import (
	"github.com/showwin/speedtest-go/speedtest"
	"fmt"
)

type Result struct {
	DownloadMbps float64
	UploadMbps float64
	PingMs float64
}

func RunSpeedtest() (*Result, error) {
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targetServer, _ := serverList.FindServer([]int{})

	err := targetServer[0].DownloadTest(false)
	if err != nil {
		return nil, err
	}

	err = targetServer[0].UploadTest(false)
	if err != nil {
		return nil, err
	}

	result := &Result{
		DownloadMbps: targetServer[0].DLSpeed,
		UploadMbps: targetServer[0].ULSpeed,
		PingMs: targetServer[0].Latency.Seconds() * 1000,
	}

	fmt.Printf("Download: %.2f Mbps, Upload: %.2f Mbps, Ping: %.2f ms\n", result.DownloadMbps, result.UploadMbps, result.PingMs)
	return result, nil
}