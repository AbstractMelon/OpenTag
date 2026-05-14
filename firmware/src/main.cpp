#include <Arduino.h>
#include <WiFi.h>
#include "Config.h"
#include "Pins.h"

void setup() {
  // Allow time for the serial monitor to connect????
  // https://forum.arduino.cc/t/esp32-serial-monitor-prints-garbage-after-reset-prints-correctly/1315325/17?u=abstractmelon
  delay(1000); 
  Serial.begin(115200);
  delay(1000); 

  pinMode(LED_PIN, OUTPUT);
  pinMode(BUTTON_PIN, INPUT_PULLUP);

  digitalWrite(LED_PIN, LOW);
  
  WiFi.begin(WIFI_SSID, WIFI_PASSWORD);

  Serial.print("Connecting to WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nConnected! IP: " + WiFi.localIP().toString());
}

void loop() {
  int buttonState = digitalRead(BUTTON_PIN);

  // INPUT_PULLUP means pressed = LOW
  if (buttonState == LOW) {
    digitalWrite(LED_PIN, HIGH);
  } else {
    digitalWrite(LED_PIN, LOW);
  }
}