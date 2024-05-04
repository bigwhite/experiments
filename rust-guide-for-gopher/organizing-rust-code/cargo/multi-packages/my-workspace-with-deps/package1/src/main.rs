extern crate package2;
extern crate rand;

use rand::Rng;

fn main() {
    let mut rng = rand::thread_rng();
    let num: u32 = rng.gen();
    println!("Random number: {}", num);
    let result = package2::add(2, 2);
    println!("result: {}", result);
}
