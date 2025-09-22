#include <iostream>
int main()
{
    int i,j,arr[6];
    for (i=0;i<=5;i++)
    {
        std::cout << "Enter the elements of array " << i+1 << "\n";
        std::cin >> arr[i];
    }

    for (j=0;j<=5;j++)
    {
        std::cout << "The elements of array " << j+1 << " = " << arr[j] << "\n";
    }
}