# -*- coding: utf-8 -*-
"""
2. Write a python program to count the number occurrences all vowels present
in a string
"""
string=input("Enter the string:\t").lower()
vowels=0
for i in string:
      if(i=='a' or i=='e' or i=='i' or i=='o' or i=='u'):
            vowels=vowels+1
print("Number of vowels are:\t",vowels)

