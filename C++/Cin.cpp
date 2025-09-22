#include <iostream>
using namespace std;

int main()
{
    int n;
    

    cout << "Enetr n" << endl;
    cin >> n;
    cout << " n is "<< n << endl ;

    int a[n];
    cout << " Enetr array elements" << endl;
    for(int i = 0; i < n; i++)
    {
       cin >> a[i];
    }

    cout << "Array elements are" << endl;
    for (int i = 0; i < n ; i++)
    cout << a[i] << " ";
    
    return 0;

}
