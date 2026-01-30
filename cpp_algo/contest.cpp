#include <iostream>
#include <vector>
#include <map>
#include <set>
#include <string>
#include <sstream>

using namespace std;

vector<set<string>> read_queries(int n) {
    vector<set<string>> queries(n);
    for (int i = 0; i < n; i++) {
        int m;
        cin >> m;
        cin.ignore();
        string line;
        getline(cin, line);
        stringstream ss(line);
        string word;
        while (ss >> word) {
            queries[i].insert(word);
        }
    }
    return queries;
}

map<string, vector<int>> build_word_map(const vector<set<string>>& queries) {
    map<string, vector<int>> word_map;
    for (int i = 0; i < queries.size(); i++) {
        for (const auto& word : queries[i]) {
            word_map[word].push_back(i);
        }
    }
    return word_map;
}

void find_connected(int idx, 
                   const vector<set<string>>& queries,
                   const map<string, vector<int>>& word_map,
                   vector<bool>& visited,
                   set<string>& current_set) {
    visited[idx] = true;
    for (const auto& word : queries[idx]) {
        current_set.insert(word);
    }
    for (const auto& word : queries[idx]) {
        auto it = word_map.find(word);
        if (it != word_map.end()) {
            for (int neighbor : it->second) {
                if (!visited[neighbor]) {
                    find_connected(neighbor, queries, word_map, visited, current_set);
                }
            }
        }
    }
}

int main() {
    int n;
    cin >> n;
    auto queries = read_queries(n);
    auto word_map = build_word_map(queries);
    vector<bool> visited(n, false);
    int groups_count = 0;
    int max_group_size = 0;
    for (int i = 0; i < n; i++) {
        if (!visited[i]) {
            set<string> current_group;
            find_connected(i, queries, word_map, visited, current_group);
            groups_count++;
            int group_size = current_group.size();
            if (group_size > max_group_size) {
                max_group_size = group_size;
            }
        }
    }
    cout << groups_count << " " << max_group_size << endl;
    return 0;
}