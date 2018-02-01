package output

import (
  "io"
  "fmt"
  "strconv"
  "encoding/csv"
  "strings"
  "time"
  "../top"
  "../core"
)

type OutputFormat func(data top.GithubDataPieces, writer io.Writer) error

func RepoNames(data []core.RepoResponse) []string {
  repoNames := make([]string, len(data))
  for i, repo := range data {
    repoNames[i] = repo.Repo
    i++
  }

  return repoNames
}

type UserMoreStats struct {
  NumOriginalRepos int
  NumForkedRepos int
  TotalRepos int
  Languages []string
}

func ParseRepo(data []core.RepoResponse) UserMoreStats {
  userStats := UserMoreStats{NumOriginalRepos: 0, NumForkedRepos: 0, TotalRepos: len(data)}

  for i, repo := range data {
    if repo.Fork {
      userStats.NumForkedRepos += 1
    } else {
      userStats.NumOriginalRepos += 1
    }
    i++
  }

  return userStats
}

func PlainOutput(data top.GithubDataPieces, writer io.Writer) error {
  fmt.Fprintln(writer, "USERS\n--------")
  for i, piece := range data {
    fmt.Fprintf(writer, "#%+v: %+v (%+v):%+v (%+v) %+v repos: %d %+v %+v\n", i + 1, piece.User.Name, piece.User.Login, piece.Contributions, piece.User.Company, strings.Join(piece.Organizations, ","), len(piece.Repos), strings.Join(RepoNames(piece.Repos), ","), strings.Join(piece.Languages, ","))
    fmt.Fprintf(writer, "MORE STATS %+v\n", ParseRepo(piece.Repos))
  }
  fmt.Fprintln(writer, "\nORGANIZATIONS\n--------")
  for i, org := range data.TopOrgs(10) {
    fmt.Fprintf(writer, "#%+v: %+v (%+v)\n", i + 1, org.Name, org.MemberCount)
  }
  return nil
}

func CsvOutput(data top.GithubDataPieces, writer io.Writer) error {
  w := csv.NewWriter(writer)
  if err := w.Write([]string{"rank", "name", "login", "email", "contributions", "repos", "company", "organizations"}); err != nil {
    return err
  }
  for i, piece := range data {
    rank := strconv.Itoa(i + 1)
    name := piece.User.Name
    login := piece.User.Login
    email := piece.User.Email
    contribs := strconv.Itoa(piece.Contributions)
    orgs := strings.Join(piece.Organizations, ",")
    repos := strings.Join(RepoNames(piece.Repos), ",")
    company := piece.User.Company
    if err := w.Write([]string{ rank, name, login, email, contribs, repos, company, orgs }); err != nil {
      return err
    }
  }
  w.Flush()
  return nil
}

func YamlOutput(data top.GithubDataPieces, writer io.Writer) error {
  fmt.Fprintln(writer, "users:")
  for i, piece := range data {
    fmt.Fprintf(
      writer,
      `
  - rank: %+v
    name: '%+v'
    login: '%+v'
    id: %+v
    contributions: %+v
    company: '%+v'
    organizations: '%+v'
`,
      i + 1,
      strings.Replace(piece.User.Name, "'", "''", -1),
      strings.Replace(piece.User.Login, "'", "''", -1),
      piece.User.Id,
      piece.Contributions,
      strings.Replace(piece.User.Company, "'", "''", -1),
      strings.Replace(strings.Join(piece.Organizations, ","), "'", "''", -1))
  }
  fmt.Fprintln(writer, "\norganizations:")

  for i, org := range data.TopOrgs(10) {
    fmt.Fprintf(
      writer,
      `
  - rank: %+v
    name: '%+v'
    membercount: %+v
`,
      i + 1,
      strings.Replace(org.Name, "'", "''", -1),
      org.MemberCount)
  }

  fmt.Fprintf(writer, "generated: %+v\n", time.Now())

  return nil
}
