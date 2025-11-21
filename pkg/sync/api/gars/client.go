package gars

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DefaultHTTPTimeout defines default timeout for the HTTP client if none specified in Config.
const DefaultHTTPTimeout = 30 * time.Second

// Config keeps configuration for Client.
type Config struct {
	BaseURL  string
	Username string
	Password string
	// HTTPClient allows to provide custom http.Client. When nil default client with timeout is used.
	HTTPClient *http.Client
	// Timeout overrides default timeout when HTTPClient is nil.
	Timeout time.Duration
}

// Client encapsulates access to GARS OData API.
type Client struct {
	baseURL *url.URL
	user    string
	pass    string
	http    *http.Client
}

// NewClient constructs Client for communicating with GARS API.
func NewClient(cfg Config) (*Client, error) {
	if cfg.BaseURL == "" {
		return nil, errors.New("base URL must be provided")
	}

	parsed, err := url.Parse(cfg.BaseURL)
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
		baseURL: parsed,
		user:    cfg.Username,
		pass:    cfg.Password,
		http:    httpClient,
	}, nil
}

// List retrieves collection for entity and decodes the response into provided target pointer.
// Target must be pointer to slice or struct compatible with JSON structure returned by API.
func (c *Client) List(ctx context.Context, entity string, target interface{}, opts ...Option) (*ListMetadata, error) {
	if target == nil {
		return nil, errors.New("target must not be nil")
	}

	req, err := c.newRequest(ctx, http.MethodGet, entity, opts)
	if err != nil {
		return nil, err
	}

	body, err := c.do(req)
	if err != nil {
		return nil, err
	}

	var envelope listEnvelope
	if err := decodeResponse(body, &envelope); err != nil {
		return nil, err
	}

	if err := json.Unmarshal(envelope.Value, target); err != nil {
		return nil, fmt.Errorf("decode list value: %w", err)
	}

	return &ListMetadata{Count: envelope.Count}, nil
}

// ListRaw returns collection decoded as slice of generic map for use-cases when schema is unknown.
func (c *Client) ListRaw(ctx context.Context, entity string, opts ...Option) ([]map[string]any, *ListMetadata, error) {
	var payload []map[string]any
	meta, err := c.List(ctx, entity, &payload, opts...)
	if err != nil {
		return nil, nil, err
	}
	return payload, meta, nil
}

// Get retrieves single entity instance using key expression.
func (c *Client) Get(ctx context.Context, entity, key string, target interface{}, opts ...Option) error {
	req, err := c.newRequest(ctx, http.MethodGet, entity+"("+key+")", opts)
	if err != nil {
		return err
	}

	body, err := c.do(req)
	if err != nil {
		return err
	}

	return decodeResponse(body, target)
}

func (c *Client) newRequest(ctx context.Context, method, entity string, opts []Option) (*http.Request, error) {
	if entity == "" {
		return nil, errors.New("entity must not be empty")
	}
	params := defaultQueryParams()
	for _, opt := range opts {
		opt.apply(params)
	}

	endpoint, err := url.Parse(entity)
	if err != nil {
		return nil, fmt.Errorf("invalid entity path: %w", err)
	}

	base := *c.baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.Path = strings.TrimSuffix(base.Path, "/") + "/" + strings.TrimPrefix(endpoint.Path, "/")

	query := base.Query()
	params.populateQuery(query)
	base.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, method, base.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.user != "" || c.pass != "" {
		req.SetBasicAuth(c.user, c.pass)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

func (c *Client) do(req *http.Request) ([]byte, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &Error{StatusCode: resp.StatusCode, Body: string(body)}
	}

	return body, nil
}

// Error describes non-successful API responses.
type Error struct {
	StatusCode int
	Body       string
}

func (e *Error) Error() string {
	return fmt.Sprintf("gars api error: status=%d body=%s", e.StatusCode, e.Body)
}

// ListMetadata contains optional metadata returned by list endpoint.
type ListMetadata struct {
	Count *int
}

// listEnvelope is internal helper to unwrap OData collection response.
type listEnvelope struct {
	Count *int            `json:"@odata.count,omitempty"`
	Value json.RawMessage `json:"value"`
}
