<?php
# For more information, check out our support page
# https://support.1password.com/events-reporting

$api_token = "Bearer ".getenv('EVENTS_API_TOKEN');
$url = "https://events.1password.com";

$start_time = (new \DateTime())->modify('-24 hours');

$headers = array(
  "Content-Type: application/json",
  "Authorization: $api_token"
);
$payload = array(
  "limit" => 20,
  "start_time" => $start_time->format('Y-m-d\TH:i:s\Z')
);

# Alternatively, use the cursor returned from previous responses to get any new events
# $payload = array("cursor" => $cursor);

$context = stream_context_create(
  array(
    'http' => array(
      'method' => 'POST',
      'content' => json_encode($payload),
      'header' => $headers,
    )
  )
);

$signin_attempts = file_get_contents($url."/api/v1/signinattempts", false, $context);
print($signin_attempts);

$item_usages = file_get_contents($url."/api/v1/itemusages", false, $context);
print($item_usages);

# For more information on the response, check out our support page
# https://support.1password.com/cs/events-api-reference/
