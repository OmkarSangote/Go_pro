# -*- coding: utf-8 -*-
"""
Read string from keyboard if it is alphabetic then check is it in uppercase if not convert it to
uppercase otherwise is it numeric if numeric print its binary representation
"""
x = input('Enter the string\n')
if x.isalpha() == True:
   print("Given string is alphabet")
   if x.isupper() == True:
       print("Aphabets in given string are in upper case")
   else:
       print("The given string is converted into Upper case",x.upper())
else:
    print("The given string is a numerical value")
    x=int(x)
    print("Binary represent of",x,"is",bin(x))
    
