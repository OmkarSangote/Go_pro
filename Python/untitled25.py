# -*- coding: utf-8 -*-
"""
Check given number is Armstrong number or not
"""
n=int(input("Enter the number\t"))
s=0
temp=n
while temp>0:
    d=temp%10
    s+=d**3
    temp//=10
if n==s:
    print(n,"is an armstrong number")
else:
    print(n,"is not an armstrong number")
