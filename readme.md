| Day  | Description | 
| --- | -----------   | 
| 1    | Variables |
| 2    | Data Types and string manipulation | 
| 3    | Control Flow and Logical Operators |
| 4    | Randomisation and List |
| 5    | for loop and python range |
| 6    | Functions and While loop using [reeborg](https://reeborg.ca/reeborg.html?lang=en&mode=python&menu=worlds%2Fmenus%2Freeborg_intro_en.json&name=Hurdle%201&url=worlds%2Ftutorial_en%2Fhurdle1.json) |
| 7    | Hangman Game |
| 8    | Function Parameters (aka Function with input) |
| 9    | Dictionaries |
| 10   | Functions with output |
| 11   | Blackjack Capstone Project |
| 12   | Scope |
| 13   | Debugging |
| 14   | Higher Lower Game |
| 15   | Env setup and Coffee machine |
| 16   | Object Oriented Programming |
| 17   | Benefits of OOP, Quiz game, Creating class |
| 18   | Turtle & GUI , Tuples |
| 19   | Instances, State, Higher order functions, functions as input |
| 20   | Snake Game Part 1 |
| 21   | Class Inheritance, slicing and Snake Game Part 2 |
| 22   | Pong Game |
| 23   | Turtle Capstone Project |
| 24   | Files, directories, paths, mail merge project |
| 25   | Working with csv data, panda library |
| 26   | List, Dictionary Comprehension, and NATO Alphabet Project  |
| 27   | GUI with Tkinter and function arguments (default, args, and kwargs) |
| 28   | Pomodoro with tinker |
| 29   | Password Manager GUI |
| 30   | Errors, Exceptions and JSON data |
| 31   | Flash Card capstone project |
| 32   | SMTP, datetime module, pythonanywhere.com |
| 33   | API, ISS overhead notifier |
| 34   | API Practice, GUI Quiz |
| 35   | API Keys, Authentication, Environment Variables |



#### Problem Solving
- Break a problem into list of smaller problem
- Build Todos list
- Turn the problem into comments
- Write code -> Run code -> Fix code 
- Build Flowcharts

#### Debugging
- Describe the problem
- Reproducing the bug
- Play computer
- Fix errors and watching for red underlines (editor suggestions)
- Print is your friend
- Using a debugger eg - [pythontutor](https://pythontutor.com/render.html#mode=edit)
- Take a break (brain downtime, ask friend)
- Ask stackoverflow

#### Procedual Programming (earliest paradigm)
one function is changing some variable and the same function is doing something to another variable at some point over code start to look very very spagitteand at this point it is hard to track and what is going in over code.
Where we set one procedure to do one particular thing and one procedure leads to another procedure and all in all computer is jumping from top to bottom and jumping into the function as needed.
It is difficult to create a large project with procedual programming.

#### Object Oriented Programming
We can split the larger task into smaller task and that task become reusable. Each have individual role.

Try to real world object which ***has*** attributes (variables)(not free floating inside the file) and ***does*** methods (function)(not free floating inside the file)

#### Higher order functions
💡 Function that can work with other function.

### Application Programming Interface (API)
API is a set of commands, functions, protocols, and objects that programmers can use to create software or interact with an external system.
API is a bearer between your program and external system. You are trying to use the rules that API have describe to make a request to the external system for some piece of data.

#### API Endpoint
I we want to get data from external service we need to know what location that is stored.
eg - money out of bank you need to know where the bank is and what it address.

In api lingo addrees is called API Endpoint.


#### API Request
Similar to try withdraw some money out of bank.
Teller (Validator) - Can I help you? To Prevent you from accessing the resource.

#### Response Codes
1xx - Hold on , something is happening , this is not final
2xx - Here you go, everything is successful , you should be getting the data you spected
3xx - You don't have permission to do this thing, Go Away.
4xx - You Screwed up, or what you are looking for doesn't exists.
5xx - I screwed up.

#### API Parameters
Give input when make request , so that you get different pieces of data back depend on input.

#### API Key
With API Key the API provider can track how much you are using the API and to authorise your access or deny your access.
Like Your personal account number and password


#### Environment Variables
Used for
- Convenience : change without touching large code base 
- Security : uploading codebase somewhere, not a good idea to store api key, auth keys to store with it.



