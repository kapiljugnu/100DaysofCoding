import random

word_list = ["mouse", "house", "fan", "van"]
lives = 6

word = random.choice(word_list)
# display = []
# for _ in range(len(word)):
#   display.append("_")
display = ["_" for w in word]
end_of_game = False

while not end_of_game:
  guess = input("Guess a letter: ").lower()

  if guess in display:
    print(f"You've already guessed {guess}")
  else:
    for i in range(len(word)):
      if guess == word[i]:
        display[i] = guess

    if guess not in word:
      print(f"You guessed {guess}, that's not in the word. You lose a life")
      lives -= 1
      if lives == 0:
        end_of_game = True
        print("You lose.")
  
    if "_" not in display:
      end_of_game = True
      print("You win.")

  print(f' '.join(display))
