list_of_strings = input('enter the numbers comma seperated ').split(',')

numbers = [int(n) for n in list_of_strings]

even_num = [n for n in numbers if n % 2 == 0]

print(even_num)