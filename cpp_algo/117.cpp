#include <queue>
#include <iostream>

class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};


Node* connect(Node* root) {
    if (!root) return nullptr;
    std::queue<Node*> q;
    q.push(root);
    while (!q.empty()) {
        int levelSize = q.size();
        Node* prev = nullptr;
        for (int i = 0; i < levelSize; i++) {
            Node* cur = q.front(); q.pop();
            if (prev) prev->next = cur;
            prev = cur;
            if (cur->left)  q.push(cur->left);
            if (cur->right) q.push(cur->right);
        }
    }
    return root;
}
