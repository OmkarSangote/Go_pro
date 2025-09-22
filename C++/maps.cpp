#include <iostream>
#include <map>

using namespace std;

int main()
{
    int n;
    cout << "Enetr n" << endl;
    cin >> n;
    cout << " n is "<< n << endl ;

    map<int, int> m;
    
    int a; int b;

    cout << "Enetr set" << endl;
    for (int i = 0; i < n; i++)
    {
        cin >> a >> b;
        m.insert(pair<int,int>(a,b));
    }
    
    map<int, int>::iterator i;
     cout << "o/p set" << endl;
    for (i = m.begin(); i != m.end(); ++i)
    cout << i->first << "," << i->second << "\n";

cout << m.size() << "\nis size of m" << endl;

cout << "lower bound of 8" << m.upper_bound(8)->first<< endl;

    return 0;
}