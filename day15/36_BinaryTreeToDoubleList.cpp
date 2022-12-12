// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;

    Node() {}

    Node(int _val) {
        val = _val;
        left = NULL;
        right = NULL;
    }

    Node(int _val, Node* _left, Node* _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};

class Solution {
public:
    Node* treeToDoublyList(Node* root) {
        if (root == nullptr) {
            return nullptr;
        }
        build(root);

        prev->right = head;
        head->left = prev;
        return head;
    }

private:
    Node *prev, *head;

    void build(Node *curr) {
        if (curr == nullptr) {
            return;
        }
        build(curr.left);

        if (prev == nullptr) {
            head = curr;
        } else {
            prev.right = curr;
        }
        curr.left = prev;

        prev = curr;
        build(curr.right);
    }
};
