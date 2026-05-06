#include <Arduino.h>
#include "Debugging.h"
#include "Pins.h"

// Setup function that runs on the first time the board is powered
// ESP's that I use are funky, prints here probably won't display
void setup() {
    Serial.begin(115200);
    Serial.println("Hello, world!");
    Serial.println("ESP32 working");
    pinMode(LED_PIN, OUTPUT);

    // LED blink loop
    led_blink();
}

// Main loop function that runs repeatedly after setup() is done
void loop() {

}

