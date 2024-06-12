#include <esp32-hal-ledc.h>
#include "fetch.h"
#include <HX711.h>
#include <HTTPClient.h>
#include <LiquidCrystal.h>
#include <WiFi.h>
#include <Wire.h>

HX711 hx;
LiquidCrystal lcd(LcdPinRS, LcdPinE, LcdPinD4, LcdPinD5, LcdPinD6, LcdPinD7);

void setup() {
  Serial.begin(115200);

  Serial.println("Initializing LCD...");
  lcd.begin(LcdConf1, LcdConf2);
  lcd.print("Starting...");

  Serial.println("Connecting to WiFi...");
  WiFi.begin("Wokwi-GUEST", "", 6);
  while (WiFi.status() != WL_CONNECTED) {
    delay(250);
    Serial.println("Connecting to WiFi...");
  }
  Serial.println("WiFi connected");
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  Serial.println("Initializing HX711...");
  hx.begin(DataPinHX711, ScanPinHX711);
  hx.set_scale(420);
  hx.tare();

  Serial.println("Initializing LED...");
  ledcSetup(PwmChannel, PwmFrequency, PwmResolution);
  ledcAttachPin(LedPin, PwmChannel);

  Serial.println("Initializing Time...");
  configTime(UTC_OFFSET, UTC_OFFSET_DST, NTP_SERVER);
  char formattedTime[40];
  getTime(formattedTime, 40);
  Serial.println("Local time: " + String(formattedTime)); 

  Serial.println("End of setup.");
  lcd.clear();
  lcd.print("Started!");
}

void loop() {
  
  double realWeight = hx.get_units(10);
  Serial.printf("Real weight: %.2F\n", realWeight);

  double referenceWeight = fetch(realWeight);
  Serial.printf("Reference weight: %.2F\n", referenceWeight);

  int brightness = abs(referenceWeight - realWeight) * 5.1;
  Serial.printf("Brightness: %d\n", brightness);

  ledcWrite(PwmChannel, brightness); // From 0 to 255
  printToMon(realWeight, referenceWeight);

  delay(2000);
}

void printToMon(double realWeight, double referenceWeight){
  char textOnScreen0[16];
  char textOnScreen1[16];

  sprintf(textOnScreen0, "Act. %.2F", realWeight);
  sprintf(textOnScreen1, "Ref. %.2F", referenceWeight);

  lcd.clear();
  lcd.print(textOnScreen0);
  lcd.setCursor(0, 1); // row, column
  lcd.print(textOnScreen1);
}

