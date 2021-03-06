package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "default": PresetLocations{"id", "indonesia", "jp", "japan", "mx", "mexico", "ph", "philippines", "uk", "london"},
  "id": PresetLocations{"id", "indonesia", "jakarta", "bandung", "banjarmasin", "makassar", "java", "surabaya", "bekasi", "sumedang"},
  "jp":PresetLocations{"jp", "japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "mx": PresetLocations{"mx", "mexico", "guadalajara", "chihuahua", "juarez", "cancun", "mexicali", "veracruz", "oaxaca"},
  "ph": PresetLocations{"ph", "philippines", "filipinas", "pilipinas", "manila", "makati", "cebu", "davao", "bohol", "bacolod", "iloilo", "baguio", "vigan"},
  "uk": PresetLocations{"uk","london","birmingham","leeds","glasgow","sheffield","bradford","manchester","edinburgh","liverpool","bristol","cardiff","belfast","leicester","wakefield","coventry","nottingham","newcastle"},
}

func Preset(name string) []string {
  return PRESETS[name]
}
