package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "default": PresetLocations{"id", "indonesia", "jp", "japan", "mx", "mexico", "ph", "philippines", "uk", "london"},
  "indonesia": PresetLocations{"id", "indonesia", "jakarta", "bandung", "makassar", "java", "surabaya", "bekasi", "sumedang"},
  "japan":PresetLocations{"jp", "japan", "tokyo", "yokohama", "osaka", "nagoya", "sapporo", "kobe", "kyoto", "fukuoka", "kawasaki", "saitama", "hiroshima", "sendai"},
  "mexico": PresetLocations{"mx", "mexico", "méxico", "guadalajara", "chihuahua", "juarez", "cancun", "mexicali", "veracruz", "oaxaca"},
  "philippines": PresetLocations{"ph", "philippines", "filipinas", "pilipinas", "manila", "makati", "cebu", "davao", "bohol", "bacolod", "iloilo", "baguio", "vigan"},
  "uk": PresetLocations{"uk","london","birmingham","leeds","glasgow","sheffield","bradford","manchester","edinburgh","liverpool","bristol","cardiff","belfast","leicester","wakefield","coventry","nottingham","newcastle"},
}

func Preset(name string) []string {
  return PRESETS[name]
}
