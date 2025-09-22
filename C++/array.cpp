// #include <iostream>
//  using namespace std;

//  int main()
//  {
//      int n;
//      cout << "Enter n" << endl;
//      cin >> n;
    
//     int a[n];
//     cout << "Enter array" << endl;
//     for (int i = 0; i < n; i++)
//     {
//         cin >> a[i];
//     }
    
//     cout << "o/p array" << endl;
//     for (int i = 0; i < n; i++)
//     {
//         cout << a[i] << " ";
//     }

//     return 0;
//  }

#include <iostream>
#include <vector>

using namespace std;

int main()
{
    
    int n;
    cout << "Enetr n" << endl;
    cin >> n;
    cout << " n is "<< n << endl ;
    int a;

    vector<int> g;
    cout << "Enetr vector" << endl;
    for (int i = 0; i < n; i++)
    {
        cin >> a;
        g.push_back(a);
    }
    
    vector<int>::iterator i;
     cout << "o/p vector" << endl;
    for (i = g.begin(); i != g.end(); ++i)
    cout << *i << " " ;
    
    return 0;
}