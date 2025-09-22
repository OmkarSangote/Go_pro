# -*- coding: utf-8 -*-
"""
Retrieve all words having the lenght of atleat 4 chars
"""
import re
s = re.findall(r'\w{4,5}','one two three four five six seven eight 8 9 10')
print(s)
