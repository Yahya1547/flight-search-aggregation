package utils

func formatDuration(totalMinutes int) string {
	hours := totalMinutes / 60
	minutes := totalMinutes % 60
	return  fmt.Sprintf("%dh %dm", hours, minutes)
}

func parseDurationToMinutes(durationStr string) int {
	var hours, minutes int
	fmt.Sscanf(durationStr, "%dh %dm", &hours, &minutes)
	return hours*60 + minutes
}