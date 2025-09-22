# -*- coding: utf-8 -*-
"""
6. Write a python program to read an IP address from stdin and check
whether it is valid or not
'''
Interface ethernet 0 is up
Interface ethernet 1 is down
Interface serial 0 is down
Interface serial 1 is up'''
"""
x=input("Enter IP address:\n")
if x.count(",") !=3:
    print("Invalid IP")
else:
    l=list(x.split("."))
    for i in l:
        if int(i)<0 or int(i)>255:
            print("Invalid IP")
            break
    else:
        print("Valid IP")