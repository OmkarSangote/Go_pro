# -*- coding: utf-8 -*-
"""
7.1 Write a python program to find out how many interface are “up” (from the
below input).
Interface ethernet 0 is up
Interface ethernet 1 is down
Interface serial 0 is down
Interface serial 1 is up
"""
x= '''Interface ethernet0 is up
Interface ethernet1 is down
Interface serial0 is down
Interface serial1 is up'''
print("number of up interfaces is",x.count('up'))
print("interface in up state are:")
for i in x.splitlines():
    if'up' in i:
        print(i.split())
