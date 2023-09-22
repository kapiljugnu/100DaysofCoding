import random

MAX_SCORE = 21
BLACK_JACK = 999999
ACE = 11


def calculate_score(cards):
  total = sum(cards)

  if len(cards) == 2 and total == MAX_SCORE:
    return BLACK_JACK

  if total > MAX_SCORE and ACE in cards:
    cards.remove(ACE)
    cards.append(1)
    return sum(cards)

  return total


def deal_card():
  cards = [ACE, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10]
  return random.choice(cards)


def compare(player_score, computer_score):
  if player_score > MAX_SCORE and computer_score > MAX_SCORE:
    return "You went over. You lose."

  if player_score == computer_score:
    return "Draw 🙃"
  elif computer_score == BLACK_JACK:
    return "Lose, opponent has Blackjack 😱"
  elif player_score == BLACK_JACK:
    return "Win with a Blackjack 😎"
  elif player_score > MAX_SCORE:
    return "You went over you lose 😭"
  elif computer_score > MAX_SCORE:
    return "Opponent went over. You win 😀"
  elif player_score > computer_score:
    return "You win 😀"
  else:
    return "You lose 😭"


def game():

  print("Welcome to Blackjack game.")
  player_cards = []
  computer_cards = []

  for _ in range(2):
    player_cards.append(deal_card())
    computer_cards.append(deal_card())

  is_game_over = False
  while not is_game_over:
    player_score = calculate_score(player_cards)
    computer_score = calculate_score(computer_cards)

    print(f"Your cards : {player_cards}, current score: {player_score}")
    print(f"Computer's first card: {computer_cards[0]}")

    if player_score == BLACK_JACK or computer_score == BLACK_JACK or player_score > MAX_SCORE:
      is_game_over = True
    else:
      another_card = input("Type 'y' to get another card, type 'n' to pass: ")
      if another_card == "y":
        player_cards.append(deal_card())
      else:
        is_game_over = True

  while computer_score != BLACK_JACK and computer_score < 17:
    computer_cards.append(deal_card())
    computer_score = calculate_score(computer_cards)

  print(f"Your final hand: {player_cards}, final score: {player_score}")
  print(
      f"Computer's final hand: {computer_cards}, final score: {computer_score}"
  )
  print(compare(player_score, computer_score))


game()
