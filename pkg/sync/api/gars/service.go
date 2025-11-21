package gars

import "context"

// Service wraps Client with typed helpers for common entities used by the application.
type Service struct {
	client *Client
}

// NewService creates service around provided client.
func NewService(client *Client) *Service {
	return &Service{client: client}
}

// Routes returns list of routes with applied options.
func (s *Service) Routes(ctx context.Context, opts ...Option) ([]Route, *ListMetadata, error) {
	var routes []Route
	meta, err := s.client.List(ctx, "Catalog_Маршруты", &routes, opts...)
	if err != nil {
		return nil, nil, err
	}
	return routes, meta, nil
}

// Stops returns list of stops.
func (s *Service) Stops(ctx context.Context, opts ...Option) ([]Stop, *ListMetadata, error) {
	var stops []Stop
	meta, err := s.client.List(ctx, "Catalog_Остановки", &stops, opts...)
	if err != nil {
		return nil, nil, err
	}
	return stops, meta, nil
}

// Trips returns catalog of trips.
func (s *Service) Trips(ctx context.Context, opts ...Option) ([]Trip, *ListMetadata, error) {
	var trips []Trip
	meta, err := s.client.List(ctx, "Catalog_Рейсы", &trips, opts...)
	if err != nil {
		return nil, nil, err
	}
	return trips, meta, nil
}

// TripSchedules returns schedules for trips.
func (s *Service) TripSchedules(ctx context.Context, opts ...Option) ([]TripSchedule, *ListMetadata, error) {
	var schedules []TripSchedule
	meta, err := s.client.List(ctx, "Catalog_РейсыРасписания", &schedules, opts...)
	if err != nil {
		return nil, nil, err
	}
	return schedules, meta, nil
}

// TripScheduleStops returns stops for trip schedule.
func (s *Service) TripScheduleStops(ctx context.Context, opts ...Option) ([]TripScheduleStop, *ListMetadata, error) {
	var stops []TripScheduleStop
	meta, err := s.client.List(ctx, "Catalog_РейсыРасписания_Остановки", &stops, opts...)
	if err != nil {
		return nil, nil, err
	}
	return stops, meta, nil
}

// RouteStops returns list of route stop entries.
func (s *Service) RouteStops(ctx context.Context, opts ...Option) ([]RouteStopEntry, *ListMetadata, error) {
	var stops []RouteStopEntry
	meta, err := s.client.List(ctx, "Catalog_Маршруты_Остановки", &stops, opts...)
	if err != nil {
		return nil, nil, err
	}
	return stops, meta, nil
}

// TripScheduleRegularityDates returns regularity date entries for trip schedules.
func (s *Service) TripScheduleRegularityDates(ctx context.Context, opts ...Option) ([]TripScheduleRegularity, *ListMetadata, error) {
	var items []TripScheduleRegularity
	meta, err := s.client.List(ctx, "Catalog_РейсыРасписания_РегулярностьСписокДат", &items, opts...)
	if err != nil {
		return nil, nil, err
	}
	return items, meta, nil
}

// TripScheduleSeatQuotas returns seat quota entries for trip schedules.
func (s *Service) TripScheduleSeatQuotas(ctx context.Context, opts ...Option) ([]TripScheduleSeatQuota, *ListMetadata, error) {
	var items []TripScheduleSeatQuota
	meta, err := s.client.List(ctx, "Catalog_РейсыРасписания_КвотыМест", &items, opts...)
	if err != nil {
		return nil, nil, err
	}
	return items, meta, nil
}

// Fares returns fares for trips.
func (s *Service) Fares(ctx context.Context, opts ...Option) ([]Fare, *ListMetadata, error) {
	var fares []Fare
	meta, err := s.client.List(ctx, "InformationRegister_ТарифыРейсов", &fares, opts...)
	if err != nil {
		return nil, nil, err
	}
	return fares, meta, nil
}

// SeatAvailability returns seat availability information.
func (s *Service) SeatAvailability(ctx context.Context, opts ...Option) ([]SeatAvailability, *ListMetadata, error) {
	var info []SeatAvailability
	meta, err := s.client.List(ctx, "InformationRegister_ЗанятостьМест", &info, opts...)
	if err != nil {
		return nil, nil, err
	}
	return info, meta, nil
}

// ActualTrips returns entries from InformationRegister_АктуальныеРейсы.
func (s *Service) ActualTrips(ctx context.Context, opts ...Option) ([]ActualTrip, *ListMetadata, error) {
	var records []ActualTrip
	meta, err := s.client.List(ctx, "InformationRegister_АктуальныеРейсы", &records, opts...)
	if err != nil {
		return nil, nil, err
	}
	return records, meta, nil
}

// TripSaleStatuses returns entries from InformationRegister_СостоянияПродажиРейсов.
func (s *Service) TripSaleStatuses(ctx context.Context, opts ...Option) ([]TripSaleStatus, *ListMetadata, error) {
	var records []TripSaleStatus
	meta, err := s.client.List(ctx, "InformationRegister_СостоянияПродажиРейсов", &records, opts...)
	if err != nil {
		return nil, nil, err
	}
	return records, meta, nil
}

// ActiveFares returns entries from InformationRegister_ДействующиеТарифы.
func (s *Service) ActiveFares(ctx context.Context, opts ...Option) ([]ActiveFare, *ListMetadata, error) {
	var records []ActiveFare
	meta, err := s.client.List(ctx, "InformationRegister_ДействующиеТарифы", &records, opts...)
	if err != nil {
		return nil, nil, err
	}
	return records, meta, nil
}

// ServicePrices returns entries from InformationRegister_ЦеныНаУслуги.
func (s *Service) ServicePrices(ctx context.Context, opts ...Option) ([]ServicePrice, *ListMetadata, error) {
	var records []ServicePrice
	meta, err := s.client.List(ctx, "InformationRegister_ЦеныНаУслуги", &records, opts...)
	if err != nil {
		return nil, nil, err
	}
	return records, meta, nil
}

// Fees returns entries from Catalog_Сборы.
func (s *Service) Fees(ctx context.Context, opts ...Option) ([]Fee, *ListMetadata, error) {
	var records []Fee
	meta, err := s.client.List(ctx, "Catalog_Сборы", &records, opts...)
	if err != nil {
		return nil, nil, err
	}
	return records, meta, nil
}
