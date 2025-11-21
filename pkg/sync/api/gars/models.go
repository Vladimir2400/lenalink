package gars

// Route represents Catalog_Маршруты entity.
type Route struct {
	RefKey           string           `json:"Ref_Key"`
	Description      string           `json:"Description"`
	Number           string           `json:"НомерМаршрута"`
	Distance         float64          `json:"Расстояние"`
	Stops            []RouteStopEntry `json:"Остановки"`
	FirstStopKey     string           `json:"ПерваяОстановка_Key"`
	LastStopKey      string           `json:"ПоследняяОстановка_Key"`
	Comment          string           `json:"Комментарий"`
	TransportTypeKey string           `json:"ВидСообщения"`
}

// RouteStopEntry describes route stop list items from Маршруты.Остановки table part.
type RouteStopEntry struct {
	LineNumber           string  `json:"LineNumber"`
	RefKey               string  `json:"Ref_Key"`
	StopKey              string  `json:"Остановка_Key"`
	PlatformNumber       int     `json:"НомерПлатформы"`
	DistanceFromFirst    float64 `json:"РасстояниеОтПервойОстановки"`
	DistanceFromPrevious float64 `json:"РасстояниеОтПредыдущейОстановки"`
	IncludeInName        bool    `json:"ВключитьВНазвание"`
	PlayAnnouncements    bool    `json:"ВоспроизводитьВОстановкахСледования"`
}

// Stop represents Catalog_Остановки entity.
type Stop struct {
	RefKey          string `json:"Ref_Key"`
	Description     string `json:"Description"`
	FullAddress     string `json:"ПолныйАдрес"`
	Settlement      string `json:"НаселенныйПункт"`
	Region          string `json:"Регион"`
	Country         string `json:"Страна"`
	TimeZone        string `json:"ЧасоваяЗона"`
	Coordinates     string `json:"Координаты"`
	AlternativeName string `json:"ДругиеНазванияОстановки"`
}

// Trip describes Catalog_Рейсы entity (general trip information).
type Trip struct {
	RefKey        string `json:"Ref_Key"`
	Description   string `json:"Description"`
	RouteKey      string `json:"Маршрут_Key"`
	CarrierKey    string `json:"Перевозчик_Key"`
	IsActive      bool   `json:"Активен"`
	Code          string `json:"Code"`
	TransportType string `json:"ВидСообщения"`
}

// TripSchedule describes Catalog_РейсыРасписания entity.
type TripSchedule struct {
	RefKey      string `json:"Ref_Key"`
	TripKey     string `json:"Рейс_Key"`
	ValidFrom   string `json:"ДействуетС"`
	ValidTo     string `json:"ДействуетПо"`
	Periodicity string `json:"Регулярность"`
	State       string `json:"Состояние"`
}

// TripScheduleStop describes schedule stop times for route.
type TripScheduleStop struct {
	LineNumber string  `json:"LineNumber"`
	StopKey    string  `json:"Остановка_Key"`
	Arrival    string  `json:"ВремяПрибытия"`
	Departure  string  `json:"ВремяОтправления"`
	Distance   float64 `json:"Расстояние"`
}

// Fare describes pricing information from InformationRegister_ТарифыРейсов.
type Fare struct {
	TripScheduleKey string  `json:"РейсРасписание_Key"`
	ServiceKey      string  `json:"Услуга_Key"`
	Price           float64 `json:"Цена"`
	Currency        string  `json:"Валюта"`
	FareType        string  `json:"ТипТарифа"`
}

// SeatAvailability describes seat quota entries.
type SeatAvailability struct {
	TripScheduleKey string `json:"РейсРасписание_Key"`
	Date            string `json:"Дата"`
	TotalSeats      int    `json:"КоличествоМест"`
	SoldSeats       int    `json:"КоличествоПроданных"`
	FreeSeats       int    `json:"КоличествоСвободных"`
}

// TripScheduleRegularity describes date list entries for schedule regularity.
type TripScheduleRegularity struct {
	LineNumber string `json:"LineNumber"`
	RefKey     string `json:"Ref_Key"`
	Date       string `json:"Дата"`
}

// TripScheduleSeatQuota keeps per-seat quotas for schedule.
type TripScheduleSeatQuota struct {
	LineNumber       string `json:"LineNumber"`
	RefKey           string `json:"Ref_Key"`
	CancellationTime string `json:"ВремяОтмены"`
	SeatNumber       int    `json:"НомерМеста"`
	DestinationKey   string `json:"ПунктНазначения_Key"`
	OriginKey        string `json:"ПунктОтправления_Key"`
	SeatType         string `json:"ТипМеста"`
}

// ActualTrip describes InformationRegister_АктуальныеРейсы entries.
type ActualTrip struct {
	BusKey       string `json:"Автобус_Key"`
	EffectiveOn  string `json:"ДатаДействия"`
	RouteKey     string `json:"Маршрут_Key"`
	ScheduleKey  string `json:"РейсРасписания_Key"`
	ReferenceKey string `json:"Ссылка_Key"`
}

// TripSaleStatus describes InformationRegister_СостоянияПродажиРейсов entries.
type TripSaleStatus struct {
	Period  string `json:"Period"`
	StopKey string `json:"Остановка_Key"`
	TripKey string `json:"Рейс_Key"`
	Status  string `json:"Состояние"`
}

// ActiveFare represents InformationRegister_ДействующиеТарифы records.
type ActiveFare struct {
	RecorderKey string             `json:"Recorder_Key"`
	RecordSet   []ActiveFareRecord `json:"RecordSet"`
}

// ActiveFareRecord stores fare details inside ActiveFare record set.
type ActiveFareRecord struct {
	Active                bool    `json:"Active"`
	LineNumber            string  `json:"LineNumber"`
	Period                string  `json:"Period"`
	CurrencyKey           string  `json:"Валюта_Key"`
	FareTypeKey           string  `json:"ВидТарифа_Key"`
	PriceTypeKey          string  `json:"ВидЦены_Key"`
	ValidTo               string  `json:"ДатаОкончанияДействия"`
	ForAdditionalTrips    bool    `json:"ДляДополнительныхРейсов"`
	ForCharterTrips       string  `json:"ДляЗаказныхРейсов"`
	ServiceClassKey       string  `json:"КлассОбслуживания_Key"`
	RouteKey              string  `json:"Маршрут_Key"`
	CarrierKey            string  `json:"Перевозчик_Key"`
	DestinationKey        string  `json:"ПунктНазначения_Key"`
	OriginKey             string  `json:"ПунктОтправления_Key"`
	ScheduleKey           string  `json:"РейсРасписания_Key"`
	Fare                  float64 `json:"Тариф"`
	RemoveForCharterTrips bool    `json:"УдалитьДляЗаказныхРейсов"`
}

// ServicePrice describes InformationRegister_ЦеныНаУслуги entries.
type ServicePrice struct {
	Period     string  `json:"Period"`
	OriginKey  string  `json:"ПунктОтправления_Key"`
	ServiceKey string  `json:"Услуга_Key"`
	Price      float64 `json:"Цена"`
}

// Fee describes Catalog_Сборы entity minimal subset.
type Fee struct {
	RefKey              string  `json:"Ref_Key"`
	Description         string  `json:"Description"`
	Code                string  `json:"Code"`
	Comment             string  `json:"Комментарий"`
	UpperLimit          float64 `json:"ВерхнееОграничениеСбора"`
	LowerLimit          float64 `json:"НижнееОграничениеСбора"`
	Amount              float64 `json:"Размер"`
	CalculationMethod   string  `json:"СпособРасчета"`
	RoundUp             bool    `json:"ОкруглятьВБольшуюСторону"`
	RoundingMode        string  `json:"ПорядокОкругления"`
	CashRegisterSection int     `json:"НомерСекцииККМ"`
	CarrierFee          bool    `json:"СборПеревозчика"`
	DoNotPrintReceipt   bool    `json:"НеПечататьВЧеке"`
	Archived            bool    `json:"Архивный"`
}
