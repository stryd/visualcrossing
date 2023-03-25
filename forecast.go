package visualcrossing

type DataPoints struct {
	Datetime       string   `json:"datetime"`
	DatetimeEpoch  int      `json:"datetimeEpoch"`
	Temp           float64  `json:"temp"`
	Feelslike      float64  `json:"feelslike"`
	Humidity       float64  `json:"humidity"`
	Dew            float64  `json:"dew"`
	Precip         float64  `json:"precip"`
	Precipprob     float64  `json:"precipprob"`
	Snow           float64  `json:"snow"`
	Snowdepth      float64  `json:"snowdepth"`
	Preciptype     []string `json:"preciptype"`
	Windgust       float64  `json:"windgust"`
	Windspeed      float64  `json:"windspeed"`
	Winddir        float64  `json:"winddir"`
	Pressure       float64  `json:"pressure"`
	Visibility     float64  `json:"visibility"`
	Cloudcover     float64  `json:"cloudcover"`
	Solarradiation float64  `json:"solarradiation"`
	Solarenergy    float64  `json:"solarenergy"`
	Uvindex        float64  `json:"uvindex"`
	Conditions     string   `json:"conditions"`
	Icon           string   `json:"icon"`
	Stations       []string `json:"stations"`
	Source         string   `json:"source"`
}

type Forecast struct {
	QueryCost       int     `json:"queryCost"`
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	ResolvedAddress string  `json:"resolvedAddress"`
	Address         string  `json:"address"`
	Timezone        string  `json:"timezone"`
	Tzoffset        float64 `json:"tzoffset"`
	Description     string  `json:"description"`
	Days            []struct {
		DataPoints
		Tempmax      float64 `json:"tempmax"`
		Tempmin      float64 `json:"tempmin"`
		Feelslikemax float64 `json:"feelslikemax"`
		Feelslikemin float64 `json:"feelslikemin"`
		Precipcover  float64 `json:"precipcover"`
		Severerisk   float64 `json:"severerisk"`
		Sunrise      string  `json:"sunrise"`
		SunriseEpoch int     `json:"sunriseEpoch"`
		Sunset       string  `json:"sunset"`
		SunsetEpoch  int     `json:"sunsetEpoch"`
		Moonphase    float64 `json:"moonphase"`
		Description  string  `json:"description"`
		Hours        []struct {
			DataPoints
			Severerisk float64 `json:"severerisk"`
		} `json:"hours"`
	} `json:"days"`
	Alerts []struct {
		Event       string `json:"event"`
		Description string `json:"description"`
	} `json:"alerts"`
	CurrentConditions struct {
		DataPoints
		Sunrise      string  `json:"sunrise"`
		SunriseEpoch int     `json:"sunriseEpoch"`
		Sunset       string  `json:"sunset"`
		SunsetEpoch  int     `json:"sunsetEpoch"`
		Moonphase    float64 `json:"moonphase"`
	} `json:"currentConditions"`
}
