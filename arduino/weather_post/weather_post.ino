#include <Wire.h>
#include <Adafruit_BMP085.h>
#include "DHT.h"

DHT dht(14, DHT22);
Adafruit_BMP085 bmp;
  
void setup() {
  Serial.begin(9600);
  dht.begin();
  bmp.begin();
}
  
void loop() {
  
  float dht22_Humi = dht.readHumidity();
  float dht22_Temp = dht.readTemperature();
  float bmp180_Temp = bmp.readTemperature();
  float bmp180_Pres = bmp.readPressure();

  Serial.print("DHT22 Temp : ");
  Serial.print(dht22_Temp);
  Serial.println(" *C");

  Serial.print("DHT22 Humi : ");
  Serial.print(dht22_Humi);
  Serial.println(" %");
  
  Serial.print("BMP180 Temp : ");
  Serial.print(bmp180_Temp);
  Serial.println(" *C");
  
  Serial.print("BMP180 Pressure : ");
  Serial.print(bmp.readPressure());
  Serial.println(" hPa");
//    
//    Serial.print("Altitude = ");
//    Serial.print(bmp.readAltitude());
//    Serial.println(" meters");
//
//    Serial.print("Pressure at sealevel (calculated) = ");
//    Serial.print(bmp.readSealevelPressure());
//    Serial.println(" Pa");
//
//    Serial.print("Real altitude = ");
//    Serial.print(bmp.readAltitude(101500));
//    Serial.println(" meters");
    
    Serial.println();
    Serial.println("----------------------------");
    Serial.println();
    delay(2000);
}
