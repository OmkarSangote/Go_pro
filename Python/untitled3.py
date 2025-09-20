
"""
Read date of birth of a person and calculate age of a person
"""
m = [31,28,31,30,31,30,31,31,30,31,30,31]
d1=int(input('Enter the birth date'))
m1=int(input("Enter the birth month"))
y1=int(input("Enter the birth year"))
d2=int(input('Enter the present date'))
m2=int(input("Enter the present month"))
y2=int(input("Enter the present year"))
if d1>d2:
    m2=m2-1
    d2=d2+m[m1-1]
if m1>m2:
    y2=y2-1
    m2=m2+12
date=d2-d1
month=m2-m1
year=y2-y1
print("The age is:",year,"years",month,"months",date,"days")


