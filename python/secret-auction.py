auction = {}
new_bid = True
while new_bid:
  name = input("What is your name?: ")
  bid_amount = int(input("What's your bid?: $"))
  auction[name] = bid_amount
  restart = input("Are there any other bidders? Type 'y' or 'n' ")

  if restart == 'n':
    new_bid = False

max_amount = 0
bidder_name = ""
for name in auction:
  if auction[name] > max_amount:
    max_amount = auction[name]
    bidder_name = name

print(f"The winner is {bidder_name} with a bid of ${max_amount}")
