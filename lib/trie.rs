use std::collections::HashMap;
use std::rc::Rc;
use std::cell::RefCell;

#[derive(Debug)]
struct Trie {
    root: Rc<RefCell<TrieNode>>,
}

impl Trie {
    fn new() -> Self {
        Self{
            root: Rc::new(RefCell::new(TrieNode::new())),
        }
    }

    fn register(&self, s: &str) {
        self.root.borrow_mut().register(s)
    }

    fn search(&self, s: &str) -> bool {
        self.root.borrow().search(s)
    }
}

#[derive(Debug)]
struct TrieNode {
    chars: HashMap<char, Rc<RefCell<TrieNode>>>,
    term: bool,
}

impl TrieNode {
    fn new() -> Self {
        Self{
            chars: HashMap::new(),
            term: false,
        }
    }

    fn register(&mut self, s: &str) {
        if s.is_empty() {
            self.term = true;
            return;
        }

        let c = s.chars().next().unwrap();
        if !self.chars.contains_key(&c) {
            let new_node = Rc::new(RefCell::new(TrieNode::new()));
            self.chars.insert(c, new_node);
        }
        self.chars.get(&c).unwrap().borrow_mut().register(&s[1..])
    }

    fn search(&self, s: &str) -> bool {
        if s.is_empty() && self.term {
            return true;            
        }

        let c = s.chars().next().unwrap();
        if let Some(next) = self.chars.get(&c) {
            next.borrow_mut().search(&s[1..])
        } else {
            false
        }
    }
}
