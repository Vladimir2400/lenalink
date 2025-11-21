package rzd

import (
	"context"
	"time"
)

// MockClient provides mock data for RZD (Russian Railways) integration.
// This is a temporary solution until real RZD API integration is implemented.
type MockClient struct{}

// NewMockClient creates a new mock RZD client.
func NewMockClient() *MockClient {
	return &MockClient{}
}

// GetStations returns a list of mock railway stations.
func (c *MockClient) GetStations(ctx context.Context) ([]Station, error) {
	return []Station{
		// {
		// 	Code:      "2000000",
		// 	Name:      "Москва Ярославская",
		// 	City:      "Москва",
		// 	Region:    "Москва",
		// 	Country:   "RU",
		// 	Latitude:  55.7773,
		// 	Longitude: 37.6556,
		// 	TimeZone:  "Europe/Moscow",
		// },
		// {
		// 	Code:      "2004000",
		// 	Name:      "Санкт-Петербург Московский",
		// 	City:      "Санкт-Петербург",
		// 	Region:    "Ленинградская область",
		// 	Country:   "RU",
		// 	Latitude:  59.9086,
		// 	Longitude: 30.3059,
		// 	TimeZone:  "Europe/Moscow",
		// },
		// {
		// 	Code:      "2200000",
		// 	Name:      "Казань",
		// 	City:      "Казань",
		// 	Region:    "Республика Татарстан",
		// 	Country:   "RU",
		// 	Latitude:  55.7926,
		// 	Longitude: 49.1097,
		// 	TimeZone:  "Europe/Moscow",
		// },
		// {
		// 	Code:      "2030000",
		// 	Name:      "Екатеринбург Пассажирский",
		// 	City:      "Екатеринбург",
		// 	Region:    "Свердловская область",
		// 	Country:   "RU",
		// 	Latitude:  56.8389,
		// 	Longitude: 60.5974,
		// 	TimeZone:  "Asia/Yekaterinburg",
		// },
		// {
		// 	Code:      "2040000",
		// 	Name:      "Новосибирск Главный",
		// 	City:      "Новосибирск",
		// 	Region:    "Новосибирская область",
		// 	Country:   "RU",
		// 	Latitude:  55.0302,
		// 	Longitude: 82.9200,
		// 	TimeZone:  "Asia/Novosibirsk",
		// },
		// {
		// 	Code:      "2060000",
		// 	Name:      "Иркутск Пассажирский",
		// 	City:      "Иркутск",
		// 	Region:    "Иркутская область",
		// 	Country:   "RU",
		// 	Latitude:  52.2697,
		// 	Longitude: 104.3050,
		// 	TimeZone:  "Asia/Irkutsk",
		// },
		// {
		// 	Code:      "2080000",
		// 	Name:      "Владивосток",
		// 	City:      "Владивосток",
		// 	Region:    "Приморский край",
		// 	Country:   "RU",
		// 	Latitude:  43.1332,
		// 	Longitude: 131.9113,
		// 	TimeZone:  "Asia/Vladivostok",
		// },
		// {
		// 	Code:      "2100000",
		// 	Name:      "Якутск",
		// 	City:      "Якутск",
		// 	Region:    "Республика Саха (Якутия)",
		// 	Country:   "RU",
		// 	Latitude:  62.0357,
		// 	Longitude: 129.6758,
		// 	TimeZone:  "Asia/Yakutsk",
		// },
	}, nil
}

// GetTrains returns mock train schedules for a specific route.
func (c *MockClient) GetTrains(ctx context.Context, origin, destination string, date time.Time) ([]Train, error) {
	// Mock trains for popular routes
	// baseDate := date.Truncate(24 * time.Hour)

	trains := []Train{
		// {
		// 	TrainNumber:   "002А",
		// 	TrainName:     "Красная стрела",
		// 	OriginStation: "2000000", // Moscow
		// 	DestStation:   "2004000", // Saint Petersburg
		// 	DepartureTime: baseDate.Add(23*time.Hour + 55*time.Minute),
		// 	ArrivalTime:   baseDate.Add(31*time.Hour + 50*time.Minute),
		// 	Duration:      475, // ~8 hours
		// 	Carrier:       "РЖД",
		// 	TrainType:     "Фирменный",
		// 	Distance:      650,
		// },
		// {
		// 	TrainNumber:   "016А",
		// 	TrainName:     "Арктика",
		// 	OriginStation: "2000000", // Moscow
		// 	DestStation:   "2004000", // Saint Petersburg
		// 	DepartureTime: baseDate.Add(9*time.Hour + 30*time.Minute),
		// 	ArrivalTime:   baseDate.Add(13*time.Hour + 30*time.Minute),
		// 	Duration:      240, // 4 hours (high-speed)
		// 	Carrier:       "РЖД",
		// 	TrainType:     "Скоростной",
		// 	Distance:      650,
		// },
		// {
		// 	TrainNumber:   "104А",
		// 	TrainName:     "Якутия",
		// 	OriginStation: "2000000", // Moscow
		// 	DestStation:   "2100000", // Yakutsk
		// 	DepartureTime: baseDate.Add(15*time.Hour + 20*time.Minute),
		// 	ArrivalTime:   baseDate.Add(168*time.Hour + 45*time.Minute), // ~7 days
		// 	Duration:      9205,                                         // 153+ hours
		// 	Carrier:       "РЖД",
		// 	TrainType:     "Пассажирский",
		// 	Distance:      8300,
		// },
		// {
		// 	TrainNumber:   "056У",
		// 	TrainName:     "Дальневосточный",
		// 	OriginStation: "2000000", // Moscow
		// 	DestStation:   "2080000", // Vladivostok
		// 	DepartureTime: baseDate.Add(14*time.Hour + 10*time.Minute),
		// 	ArrivalTime:   baseDate.Add(166*time.Hour + 35*time.Minute), // ~7 days
		// 	Duration:      9145,                                         // 152+ hours
		// 	Carrier:       "РЖД",
		// 	TrainType:     "Пассажирский",
		// 	Distance:      9259,
		// },
		// {
		// 	TrainNumber:   "010А",
		// 	TrainName:     "Байкал",
		// 	OriginStation: "2000000", // Moscow
		// 	DestStation:   "2060000", // Irkutsk
		// 	DepartureTime: baseDate.Add(16*time.Hour + 45*time.Minute),
		// 	ArrivalTime:   baseDate.Add(88*time.Hour + 20*time.Minute), // ~3 days
		// 	Duration:      4295,                                        // 71+ hours
		// 	Carrier:       "РЖД",
		// 	TrainType:     "Фирменный",
		// 	Distance:      5191,
		// },
	}

	// Filter trains based on origin/destination if provided
	if origin != "" || destination != "" {
		filtered := []Train{}
		for _, train := range trains {
			if (origin == "" || train.OriginStation == origin) &&
				(destination == "" || train.DestStation == destination) {
				filtered = append(filtered, train)
			}
		}
		return filtered, nil
	}

	return trains, nil
}

// GetTickets returns mock ticket pricing for a specific train.
func (c *MockClient) GetTickets(ctx context.Context, trainNumber string) ([]Ticket, error) {
	// Mock tickets with different car classes
	return []Ticket{
		// {
		// 	TrainNumber:    trainNumber,
		// 	CarType:        "Плацкарт",
		// 	Price:          2500.00,
		// 	AvailableSeats: 54,
		// 	ServiceClass:   "3-й класс",
		// },
		// {
		// 	TrainNumber:    trainNumber,
		// 	CarType:        "Купе",
		// 	Price:          4200.00,
		// 	AvailableSeats: 36,
		// 	ServiceClass:   "2-й класс",
		// },
		// {
		// 	TrainNumber:    trainNumber,
		// 	CarType:        "СВ (спальный вагон)",
		// 	Price:          8500.00,
		// 	AvailableSeats: 18,
		// 	ServiceClass:   "1-й класс",
		// },
		// {
		// 	TrainNumber:    trainNumber,
		// 	CarType:        "Сидячий",
		// 	Price:          1200.00,
		// 	AvailableSeats: 62,
		// 	ServiceClass:   "Эконом",
		// },
	}, nil
}
