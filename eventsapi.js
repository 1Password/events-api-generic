// For more information, check out our support page
// https://support.1password.com/events-reporting

const apiToken = process.env.EVENTS_API_TOKEN;
const url = "https://events.1password.com";

const startTime = new Date();
startTime.setHours(startTime.getHours() - 24);

const headers = {
  "Content-Type": "application/json",
  Authorization: `Bearer ${apiToken}`,
};
const payload = {
  limit: 20,
  start_time: startTime.toISOString().replace(/\.\d{3}Z$/, "Z"),
};

// Alternatively, use the cursor returned from previous responses to get any new events
// payload = { "cursor": cursor }

const options = {
  method: "POST",
  body: JSON.stringify(payload),
  headers,
};

fetch(`${url}/api/v1/signinattempts`, options).then(async (response) => {
  if (response.ok) {
    console.log(await response.json());
  } else {
    console.log("Error getting sign in attempts: status code", response.status);
  }
});

fetch(`${url}/api/v1/itemusages`, options).then(async (response) => {
  if (response.ok) {
    console.log(await response.json());
  } else {
    console.log("Error getting item usages: status code", response.status);
  }
});

fetch(`${url}/api/v1/auditevents`, options).then(async (response) => {
  if (response.ok) {
    console.log(await response.json());
  } else {
    console.log("Error getting audit events: status code", response.status);
  }
});

// For more information on the response, check out our support page
// https://support.1password.com/cs/events-api-reference/
