import turtle
import pandas

us_state_data = pandas.read_csv("50_states.csv")
all_states = us_state_data.state.to_list()
guessed_state = []

screen = turtle.Screen()
screen.title("U.S States Game")
image = "blank_states_img.gif"
screen.addshape(image)
turtle.shape(image)

while len(guessed_state) < 50:
    answer_state = screen.textinput(title=f"{len(guessed_state)}/50 State correct", prompt="What's another state's name?")
    if answer_state is None:
        break
    answer_state = answer_state.title()
    if answer_state == "Exit":
        break
    if answer_state in all_states:
        guessed_state.append(answer_state)
        state_data = us_state_data[us_state_data.state == answer_state]
        state_x = int(state_data.x.iloc[0])
        state_y = int(state_data.y.iloc[0])
        t = turtle.Turtle()
        t.penup()
        t.hideturtle()
        t.goto(state_x, state_y)
        t.write(state_data.state.item())

states_to_learn = []
for state in all_states:
    if state not in guessed_state:
       states_to_learn.append(state)

df = pandas.DataFrame(states_to_learn)
df.to_csv("states_to_learn.csv")
