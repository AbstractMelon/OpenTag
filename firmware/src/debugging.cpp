#include <Arduino.h>
#include "Debugging.h"
#include "Pins.h"

// LED blink loop
void led_blink() {
    while (true) {
        delay(1000);
        digitalWrite(LED_PIN, !digitalRead(LED_PIN));
        Serial.println("LED toggled to " + String(digitalRead(LED_PIN) ? "ON" : "OFF"));
    }
}
