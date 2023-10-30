import time
from turtle import Screen
from player import Player
from scoreboard import Scoreboard
from car_manager import CarManager

screen = Screen()
screen.setup(width=600, height=600)
screen.tracer(0)

player = Player()
scoreboard = Scoreboard()
car_manager = CarManager()

screen.listen()
screen.onkey(player.move_forward, "Up")

game_is_on = True
while game_is_on:
    time.sleep(0.1)
    screen.update()

    car_manager.create_car()
    car_manager.move_cars()

    for car in car_manager.all_cars:
        if car.distance(player) < 15:
            game_is_on = False
            scoreboard.game_over()

    # advance to new level
    if player.reached_finish_line():
        scoreboard.update_scoreboard()
        player.move_to_start_position()
        car_manager.increase_speed()


screen.exitonclick()

