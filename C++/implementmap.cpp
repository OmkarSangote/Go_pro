#include <iostream>
#include <vector>
#include <unordered_map>
using namespace std;

template<typename K, typename V>
vector<pair<K, V>> mapToVector(const unordered_map<K, V> &map) {
	return std::vector<std::pair<K, V>>(map.begin(), map.end());
}

int main()
{
	std::unordered_map<char,int> map = {
		{'A', 65}, {'B', 66}, {'C', 67}, {'D', 68}, {'E', 69}
	};

	std::vector<std::pair<char,int>> v = mapToVector(map);

	for (std::pair<char,int> p: v) {
		std::cout << '{' << p.first << ", " << p.second << '}' << std::endl;
	}

	return 0;
}
