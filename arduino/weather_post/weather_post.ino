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
      // data
      float dht22_Humi = dht.readHumidity();
      float dht22_Temp = dht.readTemperature();
      float bmp180_Temp = bmp.readTemperature();
      float bmp180_Pres = bmp.readPressure();
  
      String data = generateJSON(dht22_Humi, dht22_Temp,
        bmp180_Temp, bmp180_Pres);

      // http post
      int httpCode = http.POST(data);
      Serial.printf("httpCode: %d\n", httpCode);
      http.end();
    } else {
      Serial.printf("failed\n");
    }
    
  }

  delay(2000);
}

String generateJSON(float dht22_Humi, float dht22_Temp,
  float bmp180_Temp, float bmp180_Pres) {
  String data = "{";
  data += "\"dht22_Humi\":" + (String)dht22_Humi + ",";
  data += "\"dht22_Temp\":" + (String)dht22_Temp + ",";
  data += "\"bmp180_Temp\":" + (String)bmp180_Temp + ",";
  data += "\"bmp180_Pres\":" + (String)bmp180_Pres + "";
  data += "}";

  return data;
}
