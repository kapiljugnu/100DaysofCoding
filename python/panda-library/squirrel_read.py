import pandas
data = pandas.read_csv("Squirrel_Census_20231108.csv")
unique_color_count = data["Primary Fur Color"].value_counts()
unique_color_count.to_csv("squirrel_count.csv")