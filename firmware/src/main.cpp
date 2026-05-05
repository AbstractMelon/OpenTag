#include <Arduino.h>

#define LED_PIN 2

// Setup function that runs on the first time the board is powered
// ESP's that I use are funky, prints here probably won't display
void setup() {
    Serial.begin(115200);
    Serial.println("Hello, world!");
    Serial.println("ESP32 working");
    pinMode(LED_PIN, OUTPUT);
}

// Main loop function that runs repeatedly after setup() is done
void loop() {
    delay(1000);
    digitalWrite(LED_PIN, !digitalRead(LED_PIN));
    Serial.println("LED toggled to " + String(digitalRead(LED_PIN) ? "ON" : "OFF"));
}