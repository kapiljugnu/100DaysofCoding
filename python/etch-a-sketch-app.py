from turtle import Turtle, Screen


tim = Turtle()
screen = Screen()
screen.listen()

def move_forward():
  tim.forward(10)


def move_backward():
  tim.backward(10)


def move_clockwise():
  tim.right(10)


def move_anti_clockwise():
  tim.left(10)


def clear_screen():
  tim.reset()


screen.onkey(key="w", fun=move_forward)
screen.onkey(key="s", fun=move_backward)
screen.onkey(key="d", fun=move_clockwise)
screen.onkey(key="a", fun=move_anti_clockwise)
screen.onkey(key="c", fun=clear_screen)

screen.exitonclick()
