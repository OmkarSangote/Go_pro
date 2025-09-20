# -*- coding: utf-8 -*-
"""
4. Write a python program to check the given string is palindrome or not
"""
s=input('Enter a string:')
if s==s[::-1]:
    print('string is palindrome')
else:
    print('string is not a palindrome')

