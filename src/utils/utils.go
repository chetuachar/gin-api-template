package utils

import (
	"bufio"
	"encoding/json"
	"gin-api-template/src/logger"
	"gin-api-template/src/models"
	"math"
	"os"
	"strings"
)

// readLastNLogs reads the last N logs, excluding request logs.
func ReadLastNLogs(filePath string, n int) ([]models.LogEntry, error) {
	file, err := os.Open(filePath)
	if err != nil {
		logger.ErrorLn("Error while getting log file details :" + err.Error())
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read all logs into memory
	var allLogs []models.LogEntry
	for scanner.Scan() {
		var logEntry models.LogEntry
		line := scanner.Text()

		// Parse JSON log entry
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil {
			logger.ErrorLn("Error while extracting json values:" + err.Error())
			continue // Skip if not a valid JSON log
		}

		// Check if `msg` contains request log fields
		if isRequestLog(logEntry.Msg) {
			continue // Skip request logs
		}

		allLogs = append(allLogs, logEntry)
	}

	// Get the starting index using max()
	start := int(math.Max(float64(len(allLogs)-n), 0))
	latestLogs := allLogs[start:]

	// Reverse the slice to show latest logs first
	for i, j := 0, len(latestLogs)-1; i < j; i, j = i+1, j-1 {
		latestLogs[i], latestLogs[j] = latestLogs[j], latestLogs[i]
	}

	return latestLogs, nil
}

// isRequestLog checks if a log message contains request-related fields.
func isRequestLog(msg string) bool {
	requestFields := []string{
		"request_client_ip", "request_method", "request_post_data",
		"request_proto", "request_referer", "request_time",
		"request_ua", "request_uri", "version", "x-request-id",
	}

	for _, field := range requestFields {
		if strings.Contains(msg, field) {
			return true // Found request-related field, skip this log
		}
	}
	return false
}
