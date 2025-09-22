#include<iostream>
int main()
{
    int a,b,c;
    std::cout << "Enter 3 numbers: \n";
    std::cin >> a >> b >> c;
    if (a>b & a>c)
    {
        if (b>c)
        {
            std::cout << "Largest = " << b;

        }

        else 
        {
            std::cout << "Largest = " << c;
        }
    }

    else if (b>c & b>a)
    {

        if (c>a)
        {
            std::cout << "Second largest = " << c;

        }

        else
        {
            std::cout << "Second largest = " << a;
        } 
    }
    else
        {
            if (a>b)
            {
                std::cout << "Second largest = " << a;
            }

            else
            {
                std::cout << "Second largest = " << b;
            }
        }

 return 0;
}
