package main

import (
  "time"
  "os"

  framework "github.com/triipme/dfinity-oracle-framework"
  "github.com/triipme/dfinity-oracle-framework/models"
)

func main() {
  weatherApiKey := os.Getenv("WEATHERAPI_API_KEY")
  weatherBitApiKey := os.Getenv("WEATHERBIT_API_KEY")
  tokyoEndpoints := []models.Endpoint{
    {
      Endpoint: "http://api.weatherapi.com/v1/current.json?key="+weatherApiKey+"&q=Tokyo,JP",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.current.temp_c",
      },
    },
    {
      Endpoint: "https://api.weatherbit.io/v2.0/current?key="+weatherBitApiKey+"&city=Tokyo&country=JP",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.data[0].temp",
      },
    },
  }
  delhiEndpoints := []models.Endpoint{
    {
      Endpoint: "http://api.weatherapi.com/v1/current.json?key="+weatherApiKey+"&q=Delhi,IN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.current.temp_c",
      },
    },
    {
      Endpoint: "https://api.weatherbit.io/v2.0/current?key="+weatherBitApiKey+"&city=Delhi&country=IN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.data[0].temp",
      },
    },
  }
  hcmcEndpoints := []models.Endpoint{
    {
      Endpoint: "http://api.weatherapi.com/v1/current.json?key="+weatherApiKey+"&q=HoChiMinh,VN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.current.temp_c",
      },
    },
    {
      Endpoint: "https://api.weatherbit.io/v2.0/current?key="+weatherBitApiKey+"&city=HoChiMinh&country=VN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.data[0].temp",
      },
    },
  }
  haNoiEndpoints := []models.Endpoint{
    {
      Endpoint: "http://api.weatherapi.com/v1/current.json?key="+weatherApiKey+"&q=HaNoi,VN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.current.temp_c",
      },
    },
    {
      Endpoint: "https://api.weatherbit.io/v2.0/current?key="+weatherBitApiKey+"&city=HaNoi&country=VN",
      JSONPaths: map[string]string{
        "temperature_celsius": "$.data[0].temp",
      },
    },
  }
  config := models.Config{
    CanisterName:   "dfinity_oracle_weather",
    UpdateInterval: 1 * time.Minute,
    ReplaceCanisterCode: false,
  }
  engine := models.Engine{
    Metadata: []models.MappingMetadata{
      {Key: "Tokyo", Endpoints: tokyoEndpoints},
      {Key: "Delhi", Endpoints: delhiEndpoints},
      {Key: "HoChiMinh", Endpoints: hcmcEndpoints},
      {Key: "HaNoi", Endpoints: haNoiEndpoints},
    },
  }
  oracle := framework.NewOracle(&config, &engine)
  oracle.Bootstrap()
  oracle.Run()
}