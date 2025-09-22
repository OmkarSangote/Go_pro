# -*- coding: utf-8 -*-
"""
8. Write a python program to read a date (dd-mm-yyyy) and print the month
name according the month number.
"""
d,m,y=input("Enter the date in (dd-mm-format):").split("-")
d,m,y=int(d),int(m),int(y)
print("Month is")
if m==1:
    print("JANUARY")
elif m==2:
    print("FEBRUARY")
elif m==3:
    print("MARCH")
elif m==4:
    print("APRIL")
elif m==5:
    print("MAY")
elif m==6:
    print("JUNE")
elif m==7:
    print("JULY")
elif m==8:
    print("AUGUST")
elif m==9:
    print("SEPTEMBER")
elif m==10:
    print("OCTOBER")
elif m==11:
    print("NOVEMBER")
elif m==12:
    print("DECEMBER")
else:
    print("Invalid")