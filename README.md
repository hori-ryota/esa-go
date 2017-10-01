# esa-go

[![CircleCI](https://circleci.com/gh/hori-ryota/esa-go.svg?style=svg)](https://circleci.com/gh/hori-ryota/esa-go)
[![Coverage Status](https://coveralls.io/repos/github/hori-ryota/esa-go/badge.svg?branch=master)](https://coveralls.io/github/hori-ryota/esa-go?branch=master)

esa-go is a client library for esa API v1 written in Go

## Features

- [x] With [context](https://golang.org/pkg/context/)
- [x] Teams API
- [x] Members API
- [x] Posts API
- [x] Comments API
- [x] Star API
- [x] Watch API
- [x] Category API
- [x] Invitation API
- [x] Emoji API
- [x] User API
- [ ] Oauth API

## Requirement

- Go 1.7+ (for context)

## Usage

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/hori-ryota/esa-go/esa"
)

func main() {

	apiToken := os.Getenv("ESA_API_TOKEN")
	teamName := os.Getenv("ESA_TEAM_NAME")

	client := esa.NewClient(token, teamName)

	res, err := client.GetTeam(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(res)
}
```

## Installation

```sh
    go get github.com/hori-ryota/esa-go
```

## Author

[@hori_ryota](https://twitter.com/hori_ryota)
