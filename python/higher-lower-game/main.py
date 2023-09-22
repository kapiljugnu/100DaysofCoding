import random
from game_data import data
from art import logo


def pick_random_account():
  return random.choice(data)


# if guess == "a":
#   if celebrity_a["follower_count"] > celebrity_b["follower_count"]:
#     score += 1
#     celebrity_a = celebrity_b
#   else:
#     is_game_over = True


# else:
#   if celebrity_a["follower_count"] < celebrity_b["follower_count"]:
#     score += 1
#     celebrity_a = celebrity_b
#   else:
#     is_game_over = True
def check_answer(guess, a_followers, b_followers):
  if a_followers > b_followers:
    return guess == "A"
  else:
    return guess == "B"


def play_game():
  score = 0
  celebrity_a = pick_random_account()

  is_game_over = False
  while not is_game_over:
    celebrity_b = pick_random_account()
    print("****************************")
    print(
        f"Compare A: {celebrity_a['name']}, a {celebrity_a['description']}, from {celebrity_a['country']}"
    )
    print(
        f"Against B: {celebrity_b['name']}, a {celebrity_b['description']}, from {celebrity_b['country']}"
    )

    guess = input("Who has more followers? Type 'A' or 'B': ")
    a_follower = celebrity_a["follower_count"]
    b_follower = celebrity_b["follower_count"]
    is_correct = check_answer(guess, a_follower, b_follower)
    print("****************************")
    if is_correct:
      score += 1
      celebrity_a = celebrity_b
      print(f"You're right! Current score: {score}")
    else:
      is_game_over = True
      print(f"Sorry, that' wrong. Final score: {score}")


print(logo)
play_game()
