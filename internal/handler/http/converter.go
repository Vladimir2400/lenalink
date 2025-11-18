package http

import (
	"fmt"
	"time"

	"github.com/lenalink/backend/internal/domain"
	"github.com/lenalink/backend/internal/handler/http/dto"
)

// ToStopResponse converts domain.Stop to DTO
func ToStopResponse(stop domain.Stop) dto.StopResponse {
	return dto.StopResponse{
		ID:        stop.ID,
		Name:      stop.Name,
		City:      stop.City,
		Latitude:  stop.Latitude,
		Longitude: stop.Longitude,
	}
}

// ToSegmentResponse converts domain.Segment to DTO
func ToSegmentResponse(seg *domain.Segment) dto.SegmentResponse {
	duration := seg.ArrivalTime.Sub(seg.DepartureTime)
	durationStr := formatDuration(duration)

	return dto.SegmentResponse{
		ID:            seg.ID,
		TransportType: string(seg.TransportType),
		Provider:      seg.Provider,
		From:          ToStopResponse(seg.StartStop),
		To:            ToStopResponse(seg.EndStop),
		DepartureTime: seg.DepartureTime,
		ArrivalTime:   seg.ArrivalTime,
		Duration:      durationStr,
		Price:         seg.Price,
		Distance:      seg.Distance,
		SeatCount:     seg.SeatCount,
	}
}

// ToRouteResponse converts domain.Route to DTO
func ToRouteResponse(route *domain.Route, routeType string) dto.RouteResponse {
	segments := make([]dto.SegmentResponse, len(route.Segments))
	totalDistance := 0

	for i, seg := range route.Segments {
		segments[i] = ToSegmentResponse(&seg)
		totalDistance += seg.Distance
	}

	// Calculate total duration - handle empty segments safely
	var totalDuration time.Duration
	if len(route.Segments) > 0 {
		totalDuration = route.Segments[len(route.Segments)-1].ArrivalTime.Sub(
			route.Segments[0].DepartureTime)
	}

	return dto.RouteResponse{
		ID:            route.ID,
		Type:          routeType,
		Segments:      segments,
		TotalPrice:    route.TotalPrice,
		TotalDistance: totalDistance,
		TotalDuration: formatDuration(totalDuration),
		ReliabilityScore: route.ReliabilityScore,
	}
}

// ToBookedSegmentResponse converts domain.BookedSegment to DTO
func ToBookedSegmentResponse(booked *domain.BookedSegment) dto.BookedSegmentResponse {
	return dto.BookedSegmentResponse{
		ID:                 booked.ID,
		SegmentID:          booked.SegmentID,
		Provider:           booked.Provider,
		TransportType:      string(booked.TransportType),
		From:               ToStopResponse(booked.From),
		To:                 ToStopResponse(booked.To),
		DepartureTime:      booked.DepartureTime,
		ArrivalTime:        booked.ArrivalTime,
		TicketNumber:       booked.TicketNumber,
		Price:              booked.Price,
		Commission:         booked.Commission,
		TotalPrice:         booked.TotalPrice,
		BookingStatus:      string(booked.BookingStatus),
		ProviderBookingRef: booked.ProviderBookingRef,
	}
}

// ToPaymentResponse converts domain.Payment to DTO
func ToPaymentResponse(payment *domain.Payment) *dto.PaymentResponse {
	if payment == nil {
		return nil
	}

	resp := &dto.PaymentResponse{
		ID:                payment.ID,
		OrderID:           payment.OrderID,
		Amount:            payment.Amount,
		Currency:          payment.Currency,
		Method:            string(payment.Method),
		Status:            string(payment.Status),
		ProviderPaymentID: payment.ProviderPaymentID,
		ConfirmationURL:   payment.ConfirmationURL,
		CreatedAt:         payment.CreatedAt,
		CompletedAt:       payment.CompletedAt,
		FailureReason:     payment.FailureReason,
	}

	return resp
}

// ToBookingResponse converts domain.Booking to DTO
func ToBookingResponse(booking *domain.Booking) dto.BookingResponse {
	segments := make([]dto.BookedSegmentResponse, len(booking.Segments))
	for i, seg := range booking.Segments {
		segments[i] = ToBookedSegmentResponse(&seg)
	}

	return dto.BookingResponse{
		ID:               booking.ID,
		RouteID:          booking.RouteID,
		Status:           string(booking.Status),
		Passenger:        ToPassengerResponse(&booking.Passenger),
		Segments:         segments,
		TotalPrice:       booking.TotalPrice,
		TotalCommission:  booking.TotalCommission,
		InsurancePremium: booking.InsurancePremium,
		GrandTotal:       booking.GrandTotal,
		IncludeInsurance: booking.IncludeInsurance,
		Payment:          ToPaymentResponse(booking.Payment),
		CreatedAt:        booking.CreatedAt,
		ConfirmedAt:      booking.ConfirmedAt,
		CancelledAt:      booking.CancelledAt,
		CancellationReason: booking.CancellationReason,
	}
}

// ToPassengerResponse converts domain.Passenger to DTO
func ToPassengerResponse(p *domain.Passenger) dto.PassengerResponse {
	return dto.PassengerResponse{
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		MiddleName: p.MiddleName,
		Email:      p.Email,
		Phone:      p.Phone,
	}
}

// ToDomainPassenger converts DTO to domain.Passenger
func ToDomainPassenger(req *dto.PassengerRequest) (domain.Passenger, error) {
	dob, err := time.Parse("2006-01-02", req.DateOfBirth)
	if err != nil {
		return domain.Passenger{}, err
	}

	return domain.Passenger{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		MiddleName:     req.MiddleName,
		DateOfBirth:    dob,
		PassportNumber: req.PassportNumber,
		Email:          req.Email,
		Phone:          req.Phone,
	}, nil
}

// formatDuration formats duration as "Xh Ym" format
func formatDuration(d time.Duration) string {
	hours := int(d.Hours())
	minutes := int(d.Minutes()) % 60

	if hours == 0 {
		return fmt.Sprintf("%dm", minutes)
	}
	if minutes == 0 {
		return fmt.Sprintf("%dh", hours)
	}

	return fmt.Sprintf("%dh %dm", hours, minutes)
}
