require 'date'
require 'json'
require 'net/http'
require 'uri'

# For more information, check out our support page
# https://support.1password.com/events-reporting

api_token = ENV['EVENTS_API_TOKEN']
url = 'https://events.1password.com'

start_time = DateTime.now - 24

headers = {
  "Content-Type": "application/json",
  "Authorization": "Bearer #{api_token}"
}
payload = {
  "limit": 20,
  "start_time": start_time.strftime("%FT%TZ")
}

# Alternatively, use the cursor returned from previous responses to get any new events
# payload = { "cursor": cursor }

uri = URI.parse(url+"/api/v1/signinattempts")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
req = Net::HTTP::Post.new(uri.request_uri, headers)
req.body = payload.to_json
res = http.request(req)
if (res.code == '200')
  puts(JSON.parse(res.body))
else
  puts("Error getting sign in attempts: status code #{res.code}")
end

uri = URI.parse(url+"/api/v1/itemusages")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
req = Net::HTTP::Post.new(uri.request_uri, headers)
req.body = payload.to_json
res = http.request(req)
if (res.code == '200')
  puts(JSON.parse(res.body))
else
  puts("Error getting item usages: status code #{res.code}")
end

uri = URI.parse(url+"/api/v1/auditevents")
http = Net::HTTP.new(uri.host, uri.port)
http.use_ssl = true
req = Net::HTTP::Post.new(uri.request_uri, headers)
req.body = payload.to_json
res = http.request(req)
if (res.code == '200')
  puts(JSON.parse(res.body))
else
  puts("Error getting audit events: status code #{res.code}")
end

# For more information on the response, check out our support page
# https://support.1password.com/cs/events-api-reference/
