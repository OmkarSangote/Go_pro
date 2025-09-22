# -*- coding: utf-8 -*-
"""
Created on Tue May  4 22:46:48 2021

@author: Soumya
"""

n=int(input('Enter the rows:'))
m=int(input('Enter the columns:'))
l=[]
for i in range(n):
    k=[]
    for j in range(m):
        p=int(input('Enter a number:'))
        k.append(p)
    l.append(k)
sum=0
print('Entered matrix is:')
for i in range(n):
    for j in range(m):
        print(l[i][j],end=' ')
        sum=sum+l[i][j]
    print()
print('sum of elements of a matrix is:',sum)
