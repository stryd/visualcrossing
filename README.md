
Go Visual Crossing API for Stryd
================

A #golang package to consume [Visual Crossing](https://www.visualcrossing.com/) Timeline Weather API calls.

## Getting Started with the Visual Crossing API

All usage requires a Visual Crossing API key, which you can obtain from your [Visual Crossing account page](https://www.visualcrossing.com/account).

To use the Go libary client, instantiate a `visualcrossing.Client` with your API key:

```Go
lat := "47.202"
lng := "-123.4167"

client := visualcrossing.NewClient("APIKEY")
client.SetTimeout(10 * time.Second)
now := time.Now()
f, err := client.GetTimeLineForecast(lat, lng, &now, visualcrossing.Defaults)
if err != nil {
  // Handle error
}
```

See the [Forecast](https://github.com/stryd/visualcrossing/blob/main/forecast.go) to get a picture of the shape of the returned data.

You may also want to explore the [Visual Crossing Data Format](https://www.visualcrossing.com/resources/documentation/weather-data/weather-data-documentation)
documentation, which explains when each property is expected to be populated.

## API Arguments

The API accepts a few modification parameters. Set these via a `visualcrossing.Arguments`. If you
want the default behavior, use `visualcrossing.Defaults`. If you're looking only for the `Currently` data
object, then you should use `visualcrossing.CurrentOnly` instead. Examples:

If you'd like to set other arguments, you'll need to construct a `visualcrossing.Arguments` directly. The type is a simple map[string]string:
