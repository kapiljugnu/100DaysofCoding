import tkinter

FONT= ("Arial", 24, "bold")

window = tkinter.Tk()
window.title("Mile to Km Converter")
window.minsize(width=500, height=300)
window.config(padx=20, pady=20)

# miles input
miles_input = tkinter.Entry()
miles_input.grid(row=0, column=1)

# miles label
miles_label = tkinter.Label(text="Miles", font=FONT)
miles_label.grid(row=0, column=2)

# is_equal_to label
is_equal_to_label = tkinter.Label(text="is equal to", font=FONT)
is_equal_to_label.grid(row=1, column=0)

# converted value label
converted_value_label = tkinter.Label(text=0, font=FONT)
converted_value_label.grid(row=1, column=1)

# km label
km_label = tkinter.Label(text="Km", font=FONT)
km_label.grid(row=1, column=2)

# calculate button
def calculate():
    input = miles_input.get()
    result = float(input)*1.609
    converted_value_label["text"] = result


calculate_button = tkinter.Button(text="Calculate", font=FONT, command=calculate)
calculate_button.grid(row=2, column=1)


window.mainloop()