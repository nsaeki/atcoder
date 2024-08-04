struct TrieNode {
    children: Vec<Option<TrieNode>>,
    term: bool,
}

impl TrieNode {
    fn new() -> Self {
        let mut children = Vec::new();
        for _ in 0..26 {
            children.push(None);
        }

        TrieNode {
            children,
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
            let index = (c - b'a') as usize;
            if node.children[index].is_none() {
                node.children[index] = Some(TrieNode::new());
            }
            node = node.children[index].as_mut().unwrap();
        }
        node.term = true;
    }

    fn search(&self, word: &str) -> bool {
        let mut node = &self.root;
        for c in word.bytes() {
            let index = (c - b'a') as usize;
            if let Some(next) = node.children[index].as_ref() {
                node = next;
            } else {
                return false;
            }
        }
        node.term == true
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
