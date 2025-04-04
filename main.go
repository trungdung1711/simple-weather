/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/trungdung1711/simple-weather/cmd"
	"github.com/trungdung1711/simple-weather/internal/fetch"
)

func main() {
	cmd.Execute()

	fetch.Fetch()
}
