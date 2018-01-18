package main

type PresetLocations []string

var PRESETS = map[string]PresetLocations{
  "philippines": PresetLocations{"ph", "philippines", "filipinas", "pilipinas", "manila", "makati", "cebu", "davao", "bohol", "bacolod", "iloilo", "baguio", "vigan"}}

func Preset(name string) []string {
  return PRESETS[name]
}
