MENU = {
    "espresso": {
        "ingredients": {
            "water": 50,
            "coffee": 18,
        },
        "cost": 1.5,
    },
    "latte": {
        "ingredients": {
            "water": 200,
            "milk": 150,
            "coffee": 24,
        },
        "cost": 2.5,
    },
    "cappuccino": {
        "ingredients": {
            "water": 250,
            "milk": 100,
            "coffee": 24,
        },
        "cost": 3.0,
    }
}

resources = {
    "water": 300,
    "milk": 200,
    "coffee": 100,
}


def print_report(sale_amount):
  print(f"Water: {resources['water']}ml")
  print(f"Milk: {resources['milk']}ml")
  print(f"Coffee: {resources['coffee']}g")
  print(f"Money: ${sale_amount}")


def check_resources(type):
  coffee_type = MENU[type]
  ingredients = coffee_type['ingredients']
  for key in resources:
    if key in ingredients and resources[key] < ingredients[key]:
      return key

  return ""


def process_coins():
  quaters = 0.25 * int(input("how many quaters?: "))
  dimes = 0.10 * int(input("how many dimes?: "))
  nickles = 0.05 * int(input("how many nickles?: "))
  pennies = 0.01 * int(input("how many pennies?: "))
  return quaters + dimes + nickles + pennies


def reduce_resources(type):
  coffee_type = MENU[type]
  ingredients = coffee_type['ingredients']
  for key in resources:
    if key in ingredients:
      resources[key] -= ingredients[key]


def start_machine():
  continue_serving = True
  total_sale = 0
  while continue_serving:
    coffee_type = input("What would you like? (espresso/latte/cappuccino): ")

    if coffee_type == "off":
      continue_serving = False
    elif coffee_type == "report":
      print_report(total_sale)
    elif coffee_type in MENU:
      insufficient_resource = check_resources(coffee_type)
      if insufficient_resource != "":
        print(f"Sorry there is not enough {insufficient_resource}.")
      else:
        print("Plese insert coins.")
        total = process_coins()

        cost = MENU[coffee_type]["cost"]
        if cost > total:
          print("Sorry that's not enough money. Money refunded.")
        else:
          change = total - cost
          total_sale += cost
          reduce_resources(coffee_type)
          print(f"Here is ${round(change,2)} in change")
          print(f"Here is your {coffee_type} ☕️. Enjoy!")


start_machine()
