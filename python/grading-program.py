student_scores = {
  "Harry": 81,
  "Ron": 78,
  "Hermione": 99, 
  "Draco": 74,
  "Neville": 62,
}

student_grades = {}

for student in student_scores:
    grade = ""
    value = student_scores[student]
    if value > 90:
        grade = "Outstanding"
    elif value > 80:
        grade = "Exceeds Expectations"
    elif value > 70:
        grade = "Acceptable"
    else:
        grade = "Fail"
    student_grades[student] = grade
    

print(student_grades)