print("Welcome to the tip calculator.")
total_bill = input("What was the total bill? $")
tip = input("What percentage tip would you like to give? 10, 12 or 15? ")
people = input("How many people to split the bill? ")
amount = float(total_bill) / float(people) * (1+int(tip)/100)
print(f"Each person should pay: ${round(amount, 2)}")