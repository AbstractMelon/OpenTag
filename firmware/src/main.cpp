#include <Arduino.h>
#include <WiFi.h>
#include <HTTPClient.h>
#include "Config.h"

void setup() {
  Serial.begin(115200);
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

  Serial.print("Connecting to WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nConnected! IP: " + WiFi.localIP().toString());

  // Connect to server
  HTTPClient http;
  http.begin(SERVER_URL);
  http.addHeader("Content-Type", "application/json");

  String requestBody = "{\"client_name\":\"" + String(CLIENT_NAME) + "\"}";
  int httpResponseCode = http.POST(requestBody);

  if (httpResponseCode > 0) {
    Serial.printf("Response code: %d\n", httpResponseCode);
    Serial.println(http.getString());
  } else {
    Serial.printf("Error code: %d\n", httpResponseCode);
  }
  http.end();
}

void loop() {
}
