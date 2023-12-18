from tkinter import *
from tkinter import messagebox
import string
import random
import json

DEFAULT_EMAIL = "angela@gmail.com"

# ---------------------------- PASSWORD GENERATOR ------------------------------- #
def generate_password():
    pwd = ''.join(random.choice(string.ascii_letters+string.digits+'!@#$%^&*()') for _ in range(12))
    password_input.delete(0, END)
    password_input.insert(0, pwd)

# ---------------------------- SAVE PASSWORD ------------------------------- #
def save_password():
    website = website_input.get()
    username = username_input.get()
    password = password_input.get()

    if len(website) == 0 or len(username) == 0 or len(password) == 0:
        messagebox.showinfo(title="Oops", message="Please don't leave any fields empty!")
        return

    is_ok = messagebox.askokcancel(title=website, message=f"These are the details entered:\n Email: {username}\n Password: {password}\n Is it ok to save?")
    
    if is_ok:
        new_data = {
            website: {
                "email": username,
                "password": password
            }
        }

        try:
            with open('data.json', "r") as f:
                data = json.load(f)
        except FileNotFoundError:
            data = new_data
        else:
            data.update(new_data)

        with open('data.json', "w") as file:
            json.dump(data, file, indent=4)
            website_input.delete(0, END)
            password_input.delete(0, END)
            username_input.delete(0, END)
            username_input.insert(0, DEFAULT_EMAIL)


# ---------------------------- Find Password ------------------------------- #
def find_password():
    website = website_input.get()
    try:
        with open("data.json") as file:
            data = json.load(file)
    except FileNotFoundError:
         messagebox.showinfo(title="Error", message=f"No Data File Found")
    else:
        if website in data:
            email = data[website]["email"]
            password = data[website]["password"]
            messagebox.showinfo(title=website, message=f"Email: {email}\n Password: {password}")
        else:
            messagebox.showinfo(title="Error", message=f"No details for {website} exists")

# ---------------------------- UI SETUP ------------------------------- #

window = Tk()
window.title("Password Manager")
window.config(padx=50, pady=50)

canvas = Canvas(width=200, height=200)
logo_image = PhotoImage(file="logo.png")
canvas.create_image(100, 100, image=logo_image)
canvas.grid(row=0, column=1)

website_label = Label(text="Website:")
website_label.grid(row=1, column=0)

website_input = Entry()
website_input.grid(row=1, column=1)
website_input.focus()

username_label = Label(text="Email/Username:")
username_label.grid(row=2, column=0)

username_input = Entry(width=37)
username_input.grid(row=2, column=1, columnspan=2)
username_input.insert(0, DEFAULT_EMAIL)

password_label = Label(text="Password:")
password_label.grid(row=3, column=0)

password_input = Entry()
password_input.grid(row=3, column=1)

generate_password_button = Button(text="Generate Password", command=generate_password)
generate_password_button.grid(row=3, column=2)

add_button = Button(text="Add", width=35, command=save_password)
add_button.grid(row=4, column=1, columnspan=2)

search_button = Button(text="Search", command=find_password)
search_button.grid(row=1, column=2)



window.mainloop()