#include <iostream>
int main()
{
    int i,n,fibo[100];
    std::cout << "Enter the number of elemenrts for the series: ";
    std::cin >> n;
    fibo[0] = 0;
    fibo[1] = 1;
    for (i=2;i<=n;i++)
    {
        fibo[i] = fibo [i-2] + fibo [i-1];
    }

    for (i=0;i<=n;i++)
    {
        std::cout << fibo[i] << "\t";
    }
}