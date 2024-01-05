import requests
import os
from twilio.rest import Client

API_KEY = os.environ.get("OWM_API_KEY")
LAT = os.environ.get("APP_LAT")
LON = os.environ.get("APP_LON")

response = requests.get("https://api.openweathermap.org/data/2.5/forecast", params={ "lat": LAT, "lon": LON, "appid": API_KEY, "cnt": 4 })
response.raise_for_status()

weather_data = response.json()
will_rain = False
for hour_data in weather_data["list"]:
    condition_code = hour_data["weather"][0]["id"]
    if int(condition_code) < 700:
        will_rain = True
        break



if will_rain:
    account_sid = os.environ.get("TWILIO_ACCOUNT_SID")
    auth_token = os.environ.get("TWILIO_AUTH_TOKEN")
    client = Client(account_sid, auth_token)
    message = client.messages \
                .create(
                     body="Join Earth's mightiest heroes. Like Kevin Bacon.",
                     from_='+15017122661',
                     to='+15558675310'
                 )