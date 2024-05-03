extern crate rand;

use rand::Rng;

fn main() {
    let mut rng = rand::thread_rng();
    let num: u32 = rng.gen();
    println!("Random number: {}", num);
}
