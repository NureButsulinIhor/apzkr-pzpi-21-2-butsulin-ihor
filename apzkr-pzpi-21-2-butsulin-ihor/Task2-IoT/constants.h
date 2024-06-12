#define NTP_SERVER     "pool.ntp.org"
#define UTC_OFFSET     3600
#define UTC_OFFSET_DST 0

const int DataPinHX711 = 32;
const int ScanPinHX711 = 33;

const int PwmChannel = 0;
const int PwmFrequency = 5000; // 5 kHz frequency
const int PwmResolution = 8; // 8-bit resolution (values from 0 to 255)

const int LcdPinRS = 19;
const int LcdPinE = 18;
const int LcdPinD4 = 5;
const int LcdPinD5 = 17;
const int LcdPinD6 = 16;
const int LcdPinD7 = 4;
const int LcdConf1 = 16;
const int LcdConf2 = 2;

const int LedPin = 23;

const String ServerURL = "http://91.211.15.6:63784/";
const String AuthBearer = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImEzYTMyNjFkLTY0MGUtNGMxYy04MjE3LTE5MjQ1MDZjZjc2NiJ9.gcDKrAxmJAxXu9xISCyg2i0_2QoZOe-dJhVuB1bEjpg";