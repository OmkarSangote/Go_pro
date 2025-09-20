# -*- coding: utf-8 -*-
"""
Print the pattern
1           * * * *
1 2         * * *
1 2 3       * *
1 2 3 4     *
"""
for i in range(4,-1,-1):
    for j in range(i):
        print("*",end="\t")
    print()
    
for i in range(1,5):
    for j in range(1,i+1):
        print(j,end="\t")
    print()
    
