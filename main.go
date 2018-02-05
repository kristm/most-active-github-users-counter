package main

import (
  "flag"
  "log"
  "bufio"
  "os"
  "fmt"
  "io/ioutil"
  "./top"
  "./output"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
    return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func contains(arr []string, str string) bool {
   for _, a := range arr {
      if a == str {
         return true
      }
   }
   return false
}

func targetOutput(location string) (string, string) {
  if contains(os.Args, "--csv") {
    return "csv", fmt.Sprintf("%s-%s.csv", os.Args[0], location)
  } else {
    return "plain", ""
  }
}

var locations arrayFlags

func main() {
  secret, err := ioutil.ReadFile("secret")
  if err != nil {
    log.Fatal(err)
  }

  targetLocation := func () string { if len(os.Args) >= 2 { return os.Args[1] } else { return "default" } }()
  targetOutput, targetFile := targetOutput(targetLocation)
  log.Printf("output %s to %s", targetOutput, targetFile)

  token := flag.String("token", string(secret)[:len(secret)-1], "Github auth token")
  amount := flag.Int("amount", 20, "Amount of users to show")
  considerNum := flag.Int("consider", 100, "Amount of users to consider")
  outputOpt := flag.String("output", targetOutput, "Output format: plain, csv")
  fileName := flag.String("file", targetFile, "Output file (optional, defaults to stdout)")
  preset := flag.String("preset", targetLocation, "Preset (optional)")

  flag.Var(&locations, "location", "Location to query")
  flag.Parse()

  if *preset != "" {
    locations = Preset(*preset)
  }

  data, err := top.GithubTop(top.TopOptions { Token: *token, Locations: locations, Amount: *amount, ConsiderNum: *considerNum })

  if err != nil {
    log.Fatal(err)
  }

  var format output.OutputFormat

  if *outputOpt == "plain" {
    format = output.PlainOutput
  } else if *outputOpt == "yaml" {
    format = output.YamlOutput
  } else if *outputOpt == "csv" {
    format = output.CsvOutput
  }

  var writer *bufio.Writer
  if *fileName != "" {
    f, err := os.Create(*fileName)
    if err != nil {
      log.Fatal(err)
    }
    writer = bufio.NewWriter(f)
    defer f.Close()
  } else {
     writer = bufio.NewWriter(os.Stdout)
  }

  err = format(data, writer)
  if err != nil {
    log.Fatal(err)
  }
  writer.Flush()
}
