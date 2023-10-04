from turtle import Turtle, Screen, colormode

timmy_the_turtle = Turtle()

# square
for _ in range(4):
  timmy_the_turtle.forward(100)
  timmy_the_turtle.left(90)

timmy_the_turtle.reset()

# dotted line
for _ in range(15):
  timmy_the_turtle.forward(10)
  timmy_the_turtle.penup()
  timmy_the_turtle.forward(10)
  timmy_the_turtle.pendown()

timmy_the_turtle.reset()

# shapes
import random
colours = ["CornflowerBlue", "DarkOrchid", "IndianRed", "DeepSkyBlue", "LightSeaGreen", "wheat", "SlateGray", "SeaGreen"]
for i in range(3, 10):
  angle = 360 / i
  timmy_the_turtle.color(random.choice(colours))
  for j in range (i):
    timmy_the_turtle.forward(100)
    timmy_the_turtle.left(angle)

timmy_the_turtle.reset()

#random walk
direction = [0, 90, 180, 270]
timmy_the_turtle.pensize(15)
timmy_the_turtle.speed("fastest")
for i in range(200):
  timmy_the_turtle.color(random.choice(colours))
  timmy_the_turtle.forward(30)
  timmy_the_turtle.setheading(random.choice(direction))

timmy_the_turtle.reset()

# random color
colormode(255)
def random_color():
  r = random.randint(0, 255)
  g = random.randint(0, 255)
  b = random.randint(0, 255)
  return (r, g, b)

for i in range(200):
  timmy_the_turtle.color(random_color())
  timmy_the_turtle.forward(30)
  timmy_the_turtle.setheading(random.choice(direction))

timmy_the_turtle.reset()

# Spirograph
timmy_the_turtle.pensize(1)
def draw_spirograph(size_of_graph):
  for i in range(int(360/size_of_graph)):
    timmy_the_turtle.color(random_color())
    timmy_the_turtle.circle(100)
    timmy_the_turtle.setheading(timmy_the_turtle.heading() + size_of_graph )

draw_spirograph(10)

screen = Screen()
screen.exitonclick()