#include <iostream>
#include <set>

using namespace std;

int main()
{
    int n;
    cout << "Enetr n" << endl;
    cin >> n;
    cout << " n is "<< n << endl ;

    set<int> s;
    
    int a;
    cout << "Enetr set" << endl;
    for (int i = 0; i < n; i++)
    {
        cin >> a;
        s.insert(a);
    }
    
    set<int>::iterator i;
     cout << "o/p set" << endl;
    for (i = s.begin(); i != s.end(); ++i)
    cout << *i << " " ;

cout << s.size() << "is size of s" << endl;

cout << "lower bound of 8" << *s.upper_bound(8)<< endl;

    return 0;
}