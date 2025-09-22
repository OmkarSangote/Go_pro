# -*- coding: utf-8 -*-
"""
Number of occurance of each letter in given 
Number of words in given string
"""

s = input("Enter the string\n")
d = {}
for i in s:
    if i in d:
        d[i] = d[i]+1
    else:
        d[i] = 1
print ("Count of all characters in",s,"is\n",(d))
#word count
word={}
l = s.split(' ')
for i in l:
    word[i]=word.get(i,0)+1
print("Dict with number of occurance of word",word)

