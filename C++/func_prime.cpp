#include <iostream>
#include <math.h>
#include <iomanip>
int isprime (int);
int main()
{
    int m,n,i;
    std::cout << "Enter the range of numbers: \n";
    std::cin >> m >> n;
    std::cout << "List of prime numbers between " << m << " and " << n << " are: ";
    for (i=m;i<=n;i++)
    {
        if(isprime(i))
        {
        std::cout << std::setw(5) << i;
        }
    }
}

int isprime(int num)
{
    int i;
    for (i=2; i<=sqrt(num);i++)
    {
        if(num%i==0)
        return 0;
    }
    return 1;
}