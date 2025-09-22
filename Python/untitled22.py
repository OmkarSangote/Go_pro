# -*- coding: utf-8 -*-
"""
Genrate the first n terms of the Fibonacci series
"""
n = int(input("Enter the nth value\t"))
a = 0
b = 1
sum = 0
print("Fibonacci Series : ", end = " ")
while(sum <= n):
     print(sum, end = " ")
     a = b
     b = sum
     sum = a + b