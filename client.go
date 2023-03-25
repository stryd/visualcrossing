package visualcrossing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client  *http.Client
	BaseURL string
	APIKey  string
}

const BASEURL = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline"

func NewClient(apiKey string) *Client {
	return &Client{
		client:  http.DefaultClient,
		APIKey:  apiKey,
		BaseURL: BASEURL,
	}
}

func (c *Client) SetTimeout(to time.Duration) {
	c.client.Timeout = to
}

func (c *Client) GetTimelineForecast(lat, lng string, t *time.Time, args Arguments) (forecast *Forecast, err error) {
	return c.GetTimelineForecastCtx(context.Background(), lat, lng, t, args)
}

func (c *Client) GetTimelineForecastCtx(ctx context.Context, lat, lng string, t *time.Time, args Arguments) (forecast *Forecast, err error) {
	path := fmt.Sprintf("%s,%s", lat, lng)
	if t != nil {
		path += fmt.Sprintf("/%d", t.Unix())
	}
	return c.GetCtx(ctx, path, args)
}

func (c *Client) GetCtx(ctx context.Context, path string, args Arguments) (*Forecast, error) {
	var forecast Forecast

	url := fmt.Sprintf("%s/%s", c.BaseURL, path)

	args["key"] = c.APIKey
	params := args.ToURLValues()
	if len(params) > 0 {
		url = fmt.Sprintf("%s?%s", url, params.Encode())
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("error requesting for /%s: %w", path, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api responded with %v", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&forecast)
	if err != nil {
		return nil, fmt.Errorf("error decoding response for /%s: %w", path, err)
	}

	return &forecast, nil
}
