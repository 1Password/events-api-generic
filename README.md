# 1Password Events API Generic Scripts

## Go

To get started, replace `token` with your Events API token and `url` with the Events API URL corresponding to your 1Password account region.

Then run `go run main.go`.

The script will start a Sign In Attempt Worker and Item Usage Worker.

Any events from the last 24 hours and any new events will be printed.

Optionally, tools such as [jq](https://stedolan.github.io/jq/) can be used to format the ouput data.

## Python

To get started, replace `token` with your Events API token and `url` with the Events API URL corresponding to your 1Password account region.

Then run `python3 main.go`

The script will print at most 20 sign in attempts and item usage events from the last 24 hours.
