import datetime
import requests
import os

# For more information, check out our support page
# https://support.1password.com/events-reporting

api_token = f"Bearer {os.environ['EVENTS_API_TOKEN']}"
url = "https://events.1password.com"

start_time = datetime.datetime.now() - datetime.timedelta(hours=24)

headers = {
  "Content-Type": "application/json",
  "Authorization": api_token
}
payload = {
  "limit": 20,
  "start_time": start_time.astimezone().replace(microsecond=0).isoformat()
}

# Alternatively, use the cursor returned from previous responses to get any new events
# payload = { "cursor": cursor }

r = requests.post(f"{url}/api/v1/signinattempts", headers=headers, json=payload)
if (r.status_code == requests.codes.ok):
  print(r.json())
else:
  print("Error getting sign in attempts: status code", r.status_code)

r = requests.post(f"{url}/api/v1/itemusages", headers=headers, json=payload)
if (r.status_code == requests.codes.ok):
  print(r.json())
else:
  print("Error getting item usages: status code", r.status_code)

# For more information on the response, check out our support page
# https://support.1password.com/cs/events-api-reference/
