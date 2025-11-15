package utils

import (
	"crypto/rand"
	"fmt"
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

// CalculateDistance calculates approximate distance in km using Haversine formula
func CalculateDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadiusKm = 6371.0

	lat1Rad := degreesToRadians(lat1)
	lat2Rad := degreesToRadians(lat2)
	deltaLat := degreesToRadians(lat2 - lat1)
	deltaLon := degreesToRadians(lon2 - lon1)

	a := haversine(deltaLat) + cosine(lat1Rad)*cosine(lat2Rad)*haversine(deltaLon)
	c := 2 * asin(squareRoot(a))

	return earthRadiusKm * c
}

func degreesToRadians(degrees float64) float64 {
	return degrees * 3.14159265359 / 180
}

func cosine(rad float64) float64 {
	// Approximate cosine using Taylor series
	x := rad
	result := 1.0
	term := 1.0
	for i := 1; i < 10; i++ {
		term *= -x * x / (float64(2*i*(2*i-1)))
		result += term
	}
	return result
}

func haversine(delta float64) float64 {
	return sinePower(delta / 2)
}

func sinePower(rad float64) float64 {
	return sine(rad) * sine(rad)
}

func sine(rad float64) float64 {
	// Approximate sine using Taylor series
	x := rad
	result := x
	term := x
	for i := 1; i < 10; i++ {
		term *= -x * x / (float64((2*i + 1) * (2 * i)))
		result += term
	}
	return result
}

func squareRoot(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := x
	for i := 0; i < 10; i++ {
		z = (z + x/z) / 2
	}
	return z
}

func asin(x float64) float64 {
	// Clamp x to [-1, 1]
	if x > 1 {
		x = 1
	} else if x < -1 {
		x = -1
	}
	return atan(x / squareRoot(1-x*x))
}

func atan(x float64) float64 {
	// Approximate arctangent using Taylor series
	result := 0.0
	power := x
	sign := 1.0

	for i := 1; i < 20; i += 2 {
		result += sign * power / float64(i)
		power *= x * x
		sign *= -1
	}

	return result
}

// Ternary operator simulation
func Ternary(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
