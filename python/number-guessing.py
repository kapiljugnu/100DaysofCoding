# to modify global scope variable use global keyword as in line number 15

# there is no block sope in python
# local scope variable is at function level so variable declared inside the if block inside the function 
# is accessible anywhere inside the function 

import random

print("Welcome to the Number Guessing Game!")
print("I'm thinking of a number between 1 and 100")

difficulty = input("Choose a difficulty. Type 'easy' or 'hard': ")

lives = 5
if difficulty == "easy":
  lives = 10


def reduce_live():
  global lives
  lives -= 1
  if lives == 0:
    print("You've run out of guesses, you lose")
  else:
    print("Guess again.")


computer_number = random.randint(1, 101)
while lives > 0:
  print(f"You have {lives} attempts remaining to guess the number.")
  guess = int(input("Make a guess: "))
  if guess == computer_number:
    print(f"You got it! The answer was {computer_number}")
    lives = 0
  elif guess > computer_number:
    print("Too high.")
    reduce_live()
  else:
    print("Too low.")
    reduce_live()
