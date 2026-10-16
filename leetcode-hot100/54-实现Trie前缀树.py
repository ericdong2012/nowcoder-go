"""208. 实现 Trie（前缀树）
题目：实现 insert / search / startsWith。
思路：每个节点一个 children 字典 + is_end 标记；
      search 要求 is_end，startsWith 不要求。
复杂度：三个操作均 O(单词长度)。
"""


class Trie:
    def __init__(self):
        self.children = {}
        self.is_end = False

    def insert(self, word: str) -> None:
        node = self
        for ch in word:
            node = node.children.setdefault(ch, Trie())
        node.is_end = True

    def _walk(self, prefix: str):
        node = self
        for ch in prefix:
            node = node.children.get(ch)
            if node is None:
                return None
        return node

    def search(self, word: str) -> bool:
        node = self._walk(word)
        return node is not None and node.is_end

    def startsWith(self, prefix: str) -> bool:
        return self._walk(prefix) is not None


if __name__ == "__main__":
    t = Trie()
    t.insert("apple")
    assert t.search("apple") is True
    assert t.search("app") is False
    assert t.startsWith("app") is True
    t.insert("app")
    assert t.search("app") is True
    print("54 OK")
