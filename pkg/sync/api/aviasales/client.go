package aviasales

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultBaseURL is the default Aviasales Data API base URL.
	DefaultBaseURL = "https://api.travelpayouts.com/v2"
	// DefaultHTTPTimeout defines default timeout for HTTP client.
	DefaultHTTPTimeout = 30 * time.Second
)

// Config keeps configuration for Aviasales Client.
type Config struct {
	// BaseURL for Aviasales Data API. Defaults to DefaultBaseURL.
	BaseURL string
	// APIToken is the required API token from Travelpayouts.
	APIToken string
	// HTTPClient allows to provide custom http.Client. When nil, default client with timeout is used.
	HTTPClient *http.Client
	// Timeout overrides default timeout when HTTPClient is nil.
	Timeout time.Duration
}

// Client encapsulates access to Aviasales Data API.
type Client struct {
	baseURL  *url.URL
	apiToken string
	http     *http.Client
}

// NewClient constructs Client for communicating with Aviasales API.
func NewClient(cfg Config) (*Client, error) {
	if cfg.APIToken == "" {
		return nil, errors.New("API token must be provided")
	}

	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}

	parsed, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	httpClient := cfg.HTTPClient
	if httpClient == nil {
		timeout := cfg.Timeout
		if timeout == 0 {
			timeout = DefaultHTTPTimeout
		}
		httpClient = &http.Client{Timeout: timeout}
	}

	return &Client{
		baseURL:  parsed,
		apiToken: cfg.APIToken,
		http:     httpClient,
	}, nil
}

// GetAirports retrieves list of all airports.
func (c *Client) GetAirports(ctx context.Context) ([]Airport, error) {
	endpoint := "/data/en/airports.json"

	var response AirportsResponse
	if err := c.get(ctx, endpoint, nil, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("aviasales API returned success=false")
	}

	return response.Data, nil
}

// GetCities retrieves list of all cities.
func (c *Client) GetCities(ctx context.Context) ([]City, error) {
	endpoint := "/data/en/cities.json"

	var response CitiesResponse
	if err := c.get(ctx, endpoint, nil, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("aviasales API returned success=false")
	}

	return response.Data, nil
}

// GetPrices retrieves flight prices for a specific route.
// Parameters:
//   - origin: origin IATA code (e.g., "MOW")
//   - destination: destination IATA code (e.g., "LED")
//   - departureDate: departure date in "YYYY-MM" or "YYYY-MM-DD" format
func (c *Client) GetPrices(ctx context.Context, origin, destination, departureDate string) ([]Flight, error) {
	endpoint := "/prices/latest"

	params := url.Values{}
	params.Set("origin", origin)
	params.Set("destination", destination)
	if departureDate != "" {
		params.Set("depart_date", departureDate)
	}
	params.Set("currency", "rub")
	params.Set("limit", "1000")

	var response PriceResponse
	if err := c.get(ctx, endpoint, params, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("aviasales API returned success=false")
	}

	return response.Data, nil
}

// GetFlightSchedules retrieves flight schedules for a specific route and date range.
func (c *Client) GetFlightSchedules(ctx context.Context, origin, destination string, startDate, endDate time.Time) ([]Flight, error) {
	endpoint := "/prices/month-matrix"

	params := url.Values{}
	params.Set("origin", origin)
	params.Set("destination", destination)
	params.Set("show_to_affiliates", "true")
	params.Set("currency", "rub")

	var response PriceResponse
	if err := c.get(ctx, endpoint, params, &response); err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("aviasales API returned success=false")
	}

	return response.Data, nil
}

func (c *Client) get(ctx context.Context, endpoint string, params url.Values, target interface{}) error {
	reqURL, err := url.Parse(c.baseURL.String() + endpoint)
	if err != nil {
		return fmt.Errorf("invalid endpoint: %w", err)
	}

	if params == nil {
		params = url.Values{}
	}
	params.Set("token", c.apiToken)
	reqURL.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &Error{StatusCode: resp.StatusCode, Body: string(body)}
	}

	if target != nil {
		if err := json.Unmarshal(body, target); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

// Error describes non-successful API responses.
type Error struct {
	StatusCode int
	Body       string
}

func (e *Error) Error() string {
	return fmt.Sprintf("aviasales api error: status=%d body=%s", e.StatusCode, e.Body)
}
