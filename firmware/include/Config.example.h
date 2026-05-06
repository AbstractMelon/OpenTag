#pragma once

// Copy this file to Config.h and update with your settings.
// Do NOT commit Config.h to version control.

#define WIFI_SSID "YOUR_WIFI_SSID"
#define WIFI_PASSWORD "YOUR_WIFI_PASSWORD"

// Server base URL (without endpoint path)
// IMPORTANT: ESP32 and server must be on the same network!
// To find server IP:
//   - Linux/Mac: 'ip addr' or 'ifconfig'
//   - Windows: 'ipconfig'
// Example: #define SERVER_URL "http://192.168.1.100:8080"
#define SERVER_URL "http://YOUR_SERVER_IP:8080"

// Unique client identifier for this gun/headband unit
#define CLIENT_NAME "TagBlaster-01"