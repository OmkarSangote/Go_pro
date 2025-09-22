# -*- coding: utf-8 -*-
"""
Generate all prime numbers between n to m excluding those prime that end with digit 3. Use
while with else and continue statement
"""
m=int(input("Enter the m value\t"))
n=int(input("Enter the n value\t"))
for i in range(m,n+1):
  if i>1:
    for j in range(2,i):
        if(i % j==0):
            break
    else:
        if(i%10!=3):
            print(i)