from tkinter import *
from quiz_brain import QuizBrain

THEME_COLOR = "#375362"

class QuizInterface:

    def __init__(self, quiz_brain: QuizBrain) -> None:
        self.quiz = quiz_brain
        window = Tk()
        window.title("Quizzler")
        window.config(padx=20, pady=20, bg=THEME_COLOR, width=400, height=500)
        self.window = window

        score_label = Label(text="Score: 0", bg=THEME_COLOR, fg="white")
        score_label.grid(row=0, column=1)
        self.score_label = score_label

        question_canvas = Canvas(width=300, height=250, bg="white")
        question_canvas.grid(row=1, column=0, columnspan=2, pady=50)
        question_text = question_canvas.create_text(150, 125, text="Question text", font=("Arial", 20, "italic"), fill=THEME_COLOR, width=280)
        self.question_canvas = question_canvas
        self.question_text = question_text

        true_image = PhotoImage(file="images/true.png")
        true_button = Button(image=true_image, highlightthickness=0, command=self.true_pressed)
        true_button.grid(row=2, column=0)
        self.true_button = true_button

        false_image = PhotoImage(file="images/false.png")
        false_button = Button(image=false_image, highlightthickness=0, command=self.false_pressed)
        false_button.grid(row=2, column=1)
        self.false_button = false_button

        self.get_next_question()


        window.mainloop()
        

    def get_next_question(self):
        self.question_canvas.config(bg="white")
        if self.quiz.still_has_questions():
            self.score_label.config(text=f"Score: {self.quiz.score}")
            q_text = self.quiz.next_question()
            self.question_canvas.itemconfig(self.question_text, text=q_text)
        else:
            self.question_canvas.itemconfig(self.question_text, text="You've reached the end of the quiz.")
            self.true_button.config(state="disabled")
            self.false_button.config(state="disabled")

    def true_pressed(self):
        is_right = self.quiz.check_answer("True")
        self.give_feedback(is_right)

    def false_pressed(self):
        is_right = self.quiz.check_answer("False")
        self.give_feedback(is_right)

    def give_feedback(self, is_right):
        if is_right:
            self.question_canvas.config(bg="green")
        else:
            self.question_canvas.config(bg="red")
        self.window.after(1000, self.get_next_question)