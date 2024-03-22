# 1Password Events API Generic Scripts

This repository contains scripts in several languages to get started with the Events API.

All scripts use the `EVENTS_API_TOKEN` environment variable to load your API token. You should use [`op run`](https://developer.1password.com/docs/cli/reference/commands/run) and [secret references](https://developer.1password.com/docs/cli/secrets-reference-syntax/) provided by the [1Password CLI](https://developer.1password.com/docs/cli) to securely load environment variables.

**Example 1** - using an `.env` file, running the PHP script:

```shell
op run --env-file .env -- php eventsapi.php
```

**Example 2** - providing variables inline, running the Go script:

```shell
EVENTS_API_TOKEN="op://Vault/Item/token" op run -- go run eventsapi.go
```

You can generate an API bearer token [on 1Password.com](https://support.1password.com/events-reporting/#appendix-issue-or-revoke-bearer-tokens) or with the [CLI](https://developer.1password.com/docs/cli/reference/management-commands/events-api#events-api-create).

The script will print up to 20 audit events, sign-in attempts, or item usages from the last 24 hours.

Learn more about how to [get started with 1Password Events Reporting](https://support.1password.com/events-reporting/).
