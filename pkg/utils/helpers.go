package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"strings"
	"time"
)

// GenerateID generates a unique ID string
func GenerateID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

// GenerateBookingID generates a booking ID with format BK-XXXXX
func GenerateBookingID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("BK-%X", b)
}

// GenerateTicketNumber generates a ticket number with format TK-XX-XXXX
func GenerateTicketNumber(provider string) string {
	b := make([]byte, 4)
	rand.Read(b)
	prefix := strings.ToUpper(provider[:2])
	if len(provider) < 2 {
		prefix = strings.ToUpper(provider)
	}
	return fmt.Sprintf("TK-%s-%X", prefix, b)
}

// ParseDate parses a date string in format YYYY-MM-DD
func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// FormatDate formats a time.Time to YYYY-MM-DD
func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

// FormatDateTime formats a time.Time to YYYY-MM-DD HH:MM:SS
func FormatDateTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// ParseDateTime parses a datetime string in format YYYY-MM-DD HH:MM:SS
func ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", dateTimeStr)
}

// CalculateDuration calculates duration in minutes between two times
func CalculateDuration(from, to time.Time) int {
	return int(to.Sub(from).Minutes())
}

// CalculateGap calculates the gap (in minutes) between arrival and departure
func CalculateGap(arrivalTime, departureTime time.Time) int {
	gap := departureTime.Sub(arrivalTime)
	return int(gap.Minutes())
}

// IsValidConnection checks if connection gap is within allowed range (1-24 hours)
func IsValidConnection(arrivalTime, departureTime time.Time) bool {
	gap := departureTime.Sub(arrivalTime)
	minGap := 60 * time.Minute          // 1 hour minimum
	maxGap := 24 * time.Hour            // 24 hours maximum
	return gap >= minGap && gap <= maxGap
}

// RoundToTwoDecimals rounds a float to 2 decimal places
func RoundToTwoDecimals(value float64) float64 {
	return float64(int(value*100)) / 100
}

// CalculateAge calculates age from date of birth
func CalculateAge(dateOfBirth time.Time) int {
	today := time.Now()
	age := today.Year() - dateOfBirth.Year()
	if today.Month() < dateOfBirth.Month() ||
		(today.Month() == dateOfBirth.Month() && today.Day() < dateOfBirth.Day()) {
		age--
	}
	return age
}

// StringInSlice checks if a string is in a slice
func StringInSlice(str string, slice []string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// CalculateDistance calculates distance in km using Haversine formula
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusKm = 6371.0

	// Convert to radians
	lat1Rad := lat1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	deltaLat := (lat2 - lat1) * math.Pi / 180
	deltaLon := (lon2 - lon1) * math.Pi / 180

	// Haversine formula
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)

	c := 2 * math.Asin(math.Sqrt(a))

	return earthRadiusKm * c
}

// Ternary operator simulation
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
