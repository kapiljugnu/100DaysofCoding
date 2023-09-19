alphabets = []
for i in range(97, 123):
  alphabets += chr(i)
alphabets += alphabets


def ceaser(start_text, shift_amount, cipher_direction):
  end_text = ""
  direction = " encoded"
  if cipher_direction == "d":
    direction = "decoded"
    shift_amount *= -1
  for char in start_text:
    if char in alphabets:
      position = alphabets.index(char)
      new_position = position + shift_amount
      end_text += alphabets[new_position]
    else:
      end_text += char

  print(f"Here's the {direction} result: {end_text}")


should_continue = True
while should_continue:
  direction = input("Type 'e' to encrypt, type 'd' to decrypt:\n")
  text = input("Type your message:\n").lower()
  shift = int(input("Type the shift number:\n"))
  shift = shift % 26

  ceaser(text, shift, direction)

  restart = input("Type 'y' if you want to go again. Otherwise type 'n'.\n")
  if restart == 'n':
    should_continue = False
    print('Goodbye')
