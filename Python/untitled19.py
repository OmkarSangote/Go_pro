# -*- coding: utf-8 -*-
"""
Generate multiplication table between m to n. Read m and n from keyboard
"""
m = int(input("Enter the 'm' value\n"))
n = int(input("Enter the 'n' value\n"))
while(m<=n):
    for i in range(1,11):
        print(m*i,"\t",end="")
    m=m+1
    print()