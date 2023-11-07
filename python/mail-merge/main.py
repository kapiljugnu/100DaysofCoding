with open("./input/names/invited_names.txt") as invited_file:
    invites = invited_file.readlines()

with open("./input/letters/starting_letter.txt") as letter_file:
    letter_template = letter_file.read()

for invited in invites:
    name = invited.strip("\n")
    letter = letter_template.replace("[name]", name)
    with open(f"./output/ready_to_send/letter_for_{name}.txt", "w") as letter_file:
        letter_file.write(letter)

