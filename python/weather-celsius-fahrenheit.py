weather_c = {"Monday": 12, "Tuesday": 14, "Wedensday": 15, "Thursday": 14}

weather_f = {day:((temp*9/5)+32) for (day, temp) in weather_c.items()}

print(weather_f)