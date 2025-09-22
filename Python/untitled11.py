# -*- coding: utf-8 -*-
"""

"""
d={}
n=int(input("Enter the limit\t"))
for i in range (n):
    name=input("Enter the name\n")
    score=input("Enter the score\n")
    score=int(score)
    d[name]=score
print("The score of the players are\n",d)
max_score=0
player_name=0
for i in d:
    if max_score<d[i]:
        max_score=d[i]
        player_name=i
print("The highest scorer is",player_name,"scoring",max_score)
print("The players in alphabetical order is:")
for i in sorted(d):
    print(i,":",d[i])