# -*- coding: utf-8 -*-
"""
1. Write a python program find the number of characters present in a string
(with out using len())
"""
x = input("Enter the string\t")
count=0
for i in x:
    count+=1
print("Length of string is",count)
