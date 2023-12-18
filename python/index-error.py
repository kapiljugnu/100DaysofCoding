fruits = ["Apple", "Pear", "Orange"]

def make_pie(index):
    try:
        fruit = fruits[index]
    except IndexError:
        print(f"Index {index} not found")
    else:
        print(fruit + " pie")

make_pie(3)