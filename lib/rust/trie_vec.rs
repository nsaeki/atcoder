const ALPHABET_SIZE: usize = 26;
const BASE: u8 = b'a';

struct TrieNode {
    children: [Option<Box<TrieNode>>; ALPHABET_SIZE],
    term: bool,
}

impl TrieNode {
    fn new() -> Self {
        Self {
            children: std::array::from_fn(|_| None),
            term: false,
        }
    }
}

struct Trie {
    root: TrieNode,
}

impl Trie {
    fn new() -> Self {
        Trie {
            root: TrieNode::new(),
        }
    }

    fn register(&mut self, word: &str) {
        let mut node = &mut self.root;
        for c in word.bytes() {
            let index = (c - BASE) as usize;
            node = node.children[index].get_or_insert_with(|| Box::new(TrieNode::new()));
        }
        node.term = true;
    }

    fn search(&self, word: &str) -> bool {
        let mut node = &self.root;
        for c in word.bytes() {
            let index = (c - BASE) as usize;
            if let Some(next) = node.children[index].as_ref() {
                node = next;
            } else {
                return false;
            }
        }
        node.term
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn can_search_word() {
        let mut trie = Trie::new();
        trie.register("test");
        assert!(trie.search("test"));
    }

    #[test]
    fn can_search_same_prefix_words() {
        let mut trie = Trie::new();
        trie.register("test");
        trie.register("tests");
        assert!(trie.search("test"));
    }

    #[test]
    fn search_fails_non_registered_word() {
        let mut trie = Trie::new();
        trie.register("test");
        assert!(!trie.search("tes"));
    }
}
