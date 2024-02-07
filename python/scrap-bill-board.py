import requests
from bs4 import BeautifulSoup


date = input("Which year do you want to travel to? Type the date in this format YYYY-MM-DD: ")
response = requests.get(f"https://www.billboard.com/charts/hot-100/{date}")
# print(response.text)

soup = BeautifulSoup(response.text, 'html.parser')
titles_el = soup.find_all(name="h3" , class_="c-title")
titles = [title.getText().strip() for title in titles_el]
print(titles)

