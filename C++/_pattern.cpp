#include <iostream>
int main()
{
    int row,col;
    for (row=1;row<=4;row++)
    {
        for(col=1;col<=row;col++)
        {
            std::cout << "*";
        }
        std::cout << "\n";
    }    
}