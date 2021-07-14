# 1Password Events API Generic Scripts

To get started, replace `token` with your Events API token and `url` with the Events API URL corresponding to your 1Password account region.

Then run `python3 eventsapi.py`

The script will print at most 20 sign in attempts and item usage events from the last 24 hours.

Optionally, tools such as [jq](https://stedolan.github.io/jq/) can be used to format the ouput data.

For more information, check out our support page [here](https://support.1password.com/events-reporting/).
