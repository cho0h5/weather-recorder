// http
#include <Arduino.h>
#include <ESP8266WiFi.h>
#include <ESP8266WiFiMulti.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

ESP8266WiFiMulti WiFiMulti;

// sensors
#include <Wire.h>
#include <Adafruit_BMP085.h>
#include "DHT.h"

DHT dht(14, DHT22);
Adafruit_BMP085 bmp;

void setup() {
  // serial
  Serial.begin(9600);

  // wifi
  WiFi.mode(WIFI_STA);
  WiFiMulti.addAP("seryu-elec", "cho1418.");

  // dht, bmp
  dht.begin();
  bmp.begin();
}

void loop() {
  if ((WiFiMulti.run() == WL_CONNECTED)) {
    WiFiClient client;
    HTTPClient http;

    if (http.begin(client, "http://192.168.127.13:8080/headers")) {
      int httpCode = http.GET();
      Serial.printf("httpCode: %d\n", httpCode);
      http.end();
    } else {
      Serial.printf("failed\n");
    }
    
  }

  delay(2000);
}
