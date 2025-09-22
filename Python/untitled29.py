# -*- coding: utf-8 -*-
"""
3. Write a python program to find common characters presents in two words
"""
s1=input("Enter the first string\n").lower()
s2=input("Enter the second string\n").lower()
l=[]
for i in s1:
    for j in s2:
        if i==j:
            l.append(i)
print("Common letters are\n",set(l))