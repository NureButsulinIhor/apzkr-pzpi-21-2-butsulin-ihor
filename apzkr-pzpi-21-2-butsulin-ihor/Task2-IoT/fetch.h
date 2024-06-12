#include "constants.h"
#include <HTTPClient.h>
#include <ArduinoJson.h>

void getTime(char* resultString, size_t length){
  tm timeInfo;
  getLocalTime(&timeInfo);
  
  strftime(resultString, length, "%Y-%m-%dT%H:%M:%SZ", &timeInfo);
}

double fetch(double weight)
{
  HTTPClient http;
  
  http.begin(ServerURL + "device/polling");

  // Setup headers
  http.addHeader("Content-Type", "application/json");
  http.addHeader("Authorization", AuthBearer);

  // Getting time
  char formattedTime[40];
  getTime(formattedTime, 40);

  // Create JSON payload
  DynamicJsonDocument jsonDocument(1024);
  jsonDocument["weight"] = weight;
  jsonDocument["time"] = String(formattedTime);

  String jsonPayload;
  serializeJson(jsonDocument, jsonPayload);

  // Make the POST request
  int httpResponseCode = http.POST(jsonPayload);
  double result;

  Serial.printf("HTTP Response code: %d\n", httpResponseCode);
  if (httpResponseCode < 200 || httpResponseCode >= 300) {
    Serial.printf("HTTP Request failed, error: %s\n", http.getString().c_str());

    return 0;
  }  

  // Parse JSON response
  DynamicJsonDocument jsonDocumentResponse(1024);
  deserializeJson(jsonDocumentResponse, http.getString());

  // Extract weight
  result = jsonDocumentResponse["body"]["weight"].as<double>();

  // Close connection
  http.end();

  return result;
}