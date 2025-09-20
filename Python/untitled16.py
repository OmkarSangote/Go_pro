# -*- coding: utf-8 -*-
"""
Find the roots of a quadratic equation ax 2 +bx+c=0
"""
from math import*
a=int(input('Enter the "a" value\t'))
b=int(input('Enter the "b" value\t'))
c=int(input('Enter the "c" value\t'))
d=(b*b)-(4*a*c)
if(d > 0):
    r1 = ((-b + sqrt(d)) / (2 * a))
    r2 = ((-b - sqrt(d)) / (2 * a))
    print("Two Distinct Real Roots Exists: root1 = %.2f and root2 = %.2f" %(r1, r2))
elif(d == 0):
    r1 = r2 = (-b) / (2 * a)
    print("Two Equal and Real Roots Exists: root1 = %.2f and root2 = %.2f" %(r1, r2))
elif(d < 0):
    r1 = r2 = (-b )/ (2 * a)
    i = (sqrt(-d)) / (2 * a)
    print("Two Distinct Complex Roots Exists: root1 = %.2f+%.2f and root2 = %.2f-%.2f" %(r1, i, r2, i))

