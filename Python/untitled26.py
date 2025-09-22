# -*- coding: utf-8 -*-
"""
Find sum of the numbers and odd and even count between the limit n and m
"""

n=int(input('Enter n number:'))
m=int(input('Enter m number:'))
evn=0
odd=0
sum=0
for k in range(n,m+1):
    sum=sum+k
    if k%2==0:
        evn=evn+1
    else:
        odd=odd+1
print("The sum of the numbers is:",sum)
print('Even number count is:',evn)
print('Odd number count is:',odd)
