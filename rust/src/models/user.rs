use std::collections::HashMap;
use std::sync::{Arc, Mutex};

#[derive(Clone, Debug)]
pub struct User {
    pub id: u32,
    pub name: String,
    pub email: String,
}

impl User {
    pub fn new(id: u32, name: String, email: String) -> Self {
        User { id, name, email }
    }

    pub fn to_json(&self) -> String {
        format!(r#"{{"id":{},"name":"{}","email":"{}"}}"#, self.id, self.name, self.email)
    }
}

pub type UserStore = Arc<Mutex<HashMap<u32, User>>>;
