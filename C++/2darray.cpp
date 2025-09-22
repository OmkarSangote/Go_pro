#include <iostream>
#include <iomanip>
int main()
{
    int i,j,m,n,a[10][10];
    std::cout << "Enter the number of rows and colums : \n";
    std::cin >> m >> n;
    std::cout << "Enter the matrix elements: \n";
    for(i=0;i<m;i++)
    {
        for(j=0;j<n;j++)
        {
            std::cin >> a[i][j];
        }
    }
    std::cout << "The matrix is\n ";
    for (i=0;i<m;i++)
    {
        for(j=0;j<n;j++)
        {
            std::cout << std::setw(4) << a[i][j];
            
        }
        std::cout << "\n";
    }
    
}