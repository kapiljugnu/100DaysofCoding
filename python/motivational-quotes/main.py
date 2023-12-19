import random
import smtplib
import datetime as dt

my_email="sender@gmail.com"
password="sender-password"

today_day = dt.datetime.now().weekday() # weekday start with 0 (Monday)
if today_day == 1:
  with open("quotes.txt") as file:
    all_quotes = file.readlines()
    quote = random.choice(all_quotes)
  
  with smtplib.SMTP("smtp.gmail.com") as connection:
    connection.starttls()
    connection.login(user=my_email, password=password)
    connection.sendmail(
      from_addr=my_email,
      to_addrs="receipent@yahoo.com",
      msg=f"Subject:Today's quote\n\n{quote}"
    )
