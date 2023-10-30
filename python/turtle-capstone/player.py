from turtle import Turtle

START_POSITION = (0, -270)
MOVE_DISTANCE = 10
FINISH_LINE_Y = 280


class Player(Turtle):
    def __init__(self):
        super().__init__()
        self.shape("turtle")
        self.penup()
        self.setheading(90)
        self.goto(START_POSITION)

    def move_forward(self):
        new_y = self.ycor() + MOVE_DISTANCE
        self.sety(new_y)

    def move_to_start_position(self):
        self.goto(START_POSITION)

    def reached_finish_line(self):
        return self.ycor() > FINISH_LINE_Y
