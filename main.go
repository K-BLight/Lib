package blight_lib

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

// #region Internals

// #region Types/Constants/Variables

type (
	Client struct {
		// The base URL for the Busylight API (default: http://localhost:8989)
		APIBaseURL string
	}

	ColorName string
)

const (
	Green  ColorName = "Green"
	Red    ColorName = "Red"
	Blue   ColorName = "Blue"
	Purple ColorName = "Purple"
)

var (
	// The base URL for the Busylight HTTP Server API.
	DefaultAPIBaseURL = "http://localhost:8989"

	rClient = resty.New()

	// The RGB values for the color green.
	GreenMap = map[string]string{
		"red":   "0",
		"green": "255",
		"blue":  "0",
	}

	// The RGB values for the color red.
	RedMap = map[string]string{
		"red":   "255",
		"green": "0",
		"blue":  "0",
	}

	// The RGB values for the color blue.
	BlueMap = map[string]string{
		"red":   "0",
		"green": "0",
		"blue":  "255",
	}

	// The RGB values for the color purple.
	PurpleMap = map[string]string{
		"red":   "255",
		"green": "0",
		"blue":  "255",
	}

	// The RGB values for the color black/off.
	DefaultMap = map[string]string{
		"red":   "0",
		"green": "0",
		"blue":  "0",
	}
)

// #endregion Types/Constants/Variables

func Init(apiBaseURL string) *Client {
	if apiBaseURL == "" {
		apiBaseURL = DefaultAPIBaseURL
	}

	return &Client{
		APIBaseURL: apiBaseURL,
	}
}

func convertColor(cn ColorName) map[string]string {
	switch cn {
	case Green:
		return GreenMap
	case Red:
		return RedMap
	case Blue:
		return BlueMap
	case Purple:
		return PurpleMap
	default:
		return DefaultMap
	}
}

// #endregion Internals

// #region Standard Endpoint Functions

// Turns the light on with the specified color.
func (c *Client) TurnOn(colorName ColorName) (resp *resty.Response, err error) {
	queryParams := convertColor(colorName)

	fmt.Println("[#TurnOn][REQUEST] c.APIBaseURL: ", c.APIBaseURL)
	fmt.Println("[#TurnOn][REQUEST] Color: ", colorName)
	fmt.Println("[#TurnOn][REQUEST] QueryParams: ", queryParams)

	requestUrl := fmt.Sprintf("%s/?action=light&red=%s&green=%s&blue=%s", c.APIBaseURL, queryParams["red"], queryParams["green"], queryParams["blue"])
	fmt.Println("[#TurnOn][REQUEST] URL: ", requestUrl)

	resp, err = rClient.R().Get(requestUrl)
	if err != nil {
		fmt.Println("[#TurnOn][ERROR]: ", err)
	}

	fmt.Println("[#TurnOn][RESPONSE] Status: ", resp.Status())
	fmt.Println("[#TurnOn][RESPONSE] StatusCode: ", resp.StatusCode())
	fmt.Println("[#TurnOn][RESPONSE] Body: ", resp)

	return resp, err
}

// Turns the light off.
func (c *Client) TurnOff() (resp *resty.Response, err error) {
	resp, err = rClient.R().Get(fmt.Sprintf("%s/?action=off", c.APIBaseURL))
	if err != nil {
		fmt.Println("[#TurnOff][ERROR]: ", err)
	}

	fmt.Println("[#TurnOff][RESPONSE] Status: ", resp.Status())
	fmt.Println("[#TurnOff][RESPONSE] StatusCode: ", resp.StatusCode())
	fmt.Println("[#TurnOff][RESPONSE] Body: ", resp)

	return resp, err
}

// #endregion Standard Endpoint Functions
