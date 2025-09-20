# -*- coding: utf-8 -*-
"""
Retrieve all words having the lenght of with 3,4 and 5 chars
"""
import re
s= re.findall(r'\w{3,5}','on one two three four five six seven eight 8 9 10')
print(s)

