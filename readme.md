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
| 36   | Stock trading news alert project | 
| 37   | Habit Tracker Project | 
| 38   | Workout Tracker Project |
| 39   | Flight Deal Finder Part 1 |
| 40   | Flight Deal Finder Part 2 |
| 41   | Web Foundation - HTML, Movie Ranking Project |
| 42   | Web Foundation - Intermediate HTML, Birthday Invite Project |
| 43   | Web Foundation - CSS, Color Vocab Project |
| 44   | Web Foundation - Intermediate CSS, Motivation Meme Project |
| 45   | Web scraping with Beautiful Soup |
| 46   | Webscraping billboard |
| 47   | Webscraping product price |
| 48   | Seleium Webscraping, Python upcoming events, Cookie Clicker Project |
| 49   | Automated job application |
| 50   | Tinder swipe bot |
| 65   | Web Design |
| 72   | Advance Data Exploration with Pandas |
| 73   | Data Visualisation with Matplotlib |
| 74   | Lego Analysis |
| 75   | Google Trend Data |



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

#### Web scraping
Data that is publicily available and not copyrighted is fair game for web crawlers.
You can't commercialise copyrighted content
You can't scrape data behind authentication



### Web Design Principles
#### Colour Theory
Art and science of creating the right palette for your website

While picking color beaware the mood of your color palette and what message I am trying to convey to the user and pick appropriate color.
Red - Love, energy and intensity.
Yellow - Joy, intellect and attention.
Green - Freshness, safety and growth
Blue - Stability, trust and serenity
Purple - Royality, wealth and femininity

How to choose color?
- taking 2 color right next to each other on color wheel. (Nav, body) (called - Analogue color), The approach is not good at standing out
- you can choose color from opposite end of color wheel. (called - complementary color), Don't use this approach for text and text background
- Draw a traingle to create color palette
- Draw a perfect square

https://colorhunt.co/

https://color.adobe.com/create/color-wheel

#### Typography
Font matters, Different font have different moods, in design stick to 2 fonts
large family
- serif (feet) (moods - Traditonal, stable, respectable)
- sans-serif (perfect right angles) (moods - sensible, simple, straightforward)

#### User Interface Design (UI)
#### Define Heirarchy
- using colors
- size
#### Layout
- use blocks instead of flat structure
- 40 to 60 character per line
#### Aligment
- reduce the number of aligment points
#### White space
empty space around the text , element
#### Audience
Think about audience

#### User Experience Design (UX)
##### Simplicity
If choice between complexity and simplicity choose simplicity
##### Consistency
User confuse when navigating
#### Reading patterns
- F pattern (important part on left)
- Z pattern (eyes from left to right and then zigzag down)
#### All platform design (mobile responsive)
#### Don't use your power for evil
To perform an action which user not want


https://www.dailyui.co/


