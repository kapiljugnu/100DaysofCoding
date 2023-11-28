with open('file1.txt') as file1:
  number_list_1 = file1.readlines()
  number_list_1 = [n.strip('\n') for n in number_list_1]
  number_list_1 = [int(n) for n in number_list_1]


with open('file2.txt') as file2:
  number_list_2 = file2.readlines()
  number_list_2 = [n.strip('\n') for n in number_list_2]
  number_list_2 = [int(n) for n in number_list_2]


result = [n for n in number_list_1 if n in number_list_2]

print(number_list_1)
print(number_list_2)
print(result)
