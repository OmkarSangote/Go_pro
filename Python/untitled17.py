# -*- coding: utf-8 -*-
"""
Check given number is palindrome or not
"""
x = input("Enter the number\n")
y = reversed(x)
if list(x) == list(y):
    print("Given number is palindrome")
else:
    print("Given number is not palindrome")
