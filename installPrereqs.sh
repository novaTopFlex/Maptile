#!/bin/bash
go install fyne.io/fyne/v2@latest
go mod tidy
go install fyne.io/x/fyne
go mod tidy
go install github.com/fyne-io/terminal/cmd/fyneterm@latest
go mod tidy
