# -*- coding: utf-8 -*-
"""
Created on Mon May  3 09:10:25 2021

@author: Soumya
"""
"""

python program to create a dictionary which contains n players name along with score. 
Retrieve highest score along with player name.
Print the dictionary in alphabetical order of player name.

"""
player={}
n=int(input("Enter a number of players:"))
for i in range(n):
    name=input("Enter a name of player:")
    score=int(input("Enter a player score:"))
    player.update({name:score})
    #player[name]=score
print("Dicitionary of players with scores:",player)
s=list(player.values())
s.sort(reverse=True)
for k in player:
    if player[k]==s[0]:
        print("{} has scored highest and scored is {}".format(k,s[0]))
        break
player_sorted={}
names=sorted(player.keys())
for k in names:
    player_sorted[k]=player.get(k)
print("Dicitionary in alphabetical order of player name is:",player_sorted)