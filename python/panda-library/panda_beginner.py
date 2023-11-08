# # reading the file
# with open("weather_data.csv") as weather_data:
#     data = weather_data.readlines()
#
# print(data)

# # using inbuilt csv library
# import csv
#
# with open("weather_data.csv") as weather_data:
#     data = csv.reader(weather_data)
#     temperatures = []
#     for row in data:
#         if not row[1] == 'temp':
#             temperatures.append(int(row[1]))
#
#     print(temperatures)


# # python data analytics library "panda" which is super powerful to perform data analysis on tabular data
import pandas
data = pandas.read_csv("weather_data.csv")
# print(data)
# temp_list = data["temp"].to_list()
# avg = sum(temp_list)/len(temp_list)
# print(avg)
# or
# print(data["temp"].mean())
# print(data["temp"].max())

# Get data in column
# print(data["condition"])
# or
# print(data.condition)

# Get data in row
# data[data.day]
# or
# data[data["day"]]
# or
# print(data[data.day == "Monday"])

# max_temp = data["temp"].max()
# print(data[data.temp == max_temp])


# monday = data[data.day == "Monday"]
# monday_temp = monday.temp[0]
# monday_temp_f = monday_temp * 9/5 + 32
# print(monday_temp_f)

# create a data frame
# data_dict = {
#     "students": ["Amy", "James", "Angels"],
#     "scores": [76, 56, 65]
# }
# data = pandas.DataFrame(data_dict)
# data.to_csv("new_data.csv")
