# -*- coding: utf-8 -*-
"""
Print alphabet pattern "T" and "Z"
"""
#For T
for i in range(7):
    print("*",end="")
print()
for i in range(5):
    print("   *")    
    
    
    
#For Z
for i in range(7):
    print("*",end="")
print()
for i in range(4,-1,-1):
    print(i*" ","*")
for i in range(7):
    print("*",end="")