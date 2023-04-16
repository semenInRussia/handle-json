use std::{io::{stdin, Read}, collections::HashMap};

use serde::{Serialize, Deserialize};
use serde_json;


#[derive(Serialize, Deserialize, Debug)]
struct Input {
    data: Vec<Item>
}

#[derive(Serialize, Deserialize, Debug)]
#[serde(tag = "type")]
enum Item {
    #[serde(rename = "user")]
    User(User),
    #[serde(rename = "payment")]
    Payment(Payment),
    #[serde(rename = "address")]
    Address(Address),
}

type ID = u16;

#[derive(Serialize, Deserialize, Debug, Clone)]
struct User {
    id: ID,
    firstname: String,
    lastname: String,
    #[serde(default)]
    payments: Vec<Payment>,
    #[serde(default)]
    addresses: Vec<Address>
}

#[derive(Serialize, Deserialize, Debug, Clone)]
struct Payment {
    id: ID,
    amount: u16,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
struct Address {
    id: ID,
    address: String,
}

fn main() {
    let mut src = String::new();
    stdin().read_to_string(&mut src)
        .expect("Can't read from stdin ???");
    let input: Input = serde_json::de::from_str(&src)
        .expect("Invalid JSON from the stdin");
    let items = input.data;

    let mut users = HashMap::new();

    for it in items.iter() {
        if let Item::User(user) = it {
            users.insert(user.id, user.clone());
        }
    }

    for it in items.into_iter() {
        match it {
            Item::Payment(pm) => {
                users.get_mut(&pm.id)
                    .map(|user| (*user).payments.push(pm));
            },
            Item::Address(adrs) => {
                users.get_mut(&adrs.id)
                    .map(|user| (*user).addresses.push(adrs));
            },
            Item::User(_) => { continue; }
        }
    }

    let users: Vec<User> = users.into_values().collect();

    println!("{}", serde_json::ser::to_string_pretty(&users)
             .expect("Error when serialize users to JSON, bug of serde not mine"));
}
