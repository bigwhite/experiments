
extern crate my_local_crate1;
extern crate my_local_crate2;

fn main() {
    let x = 5;
    let y = my_local_crate1::add_one(x);
    let z = my_local_crate2::multiply_two(y);
    println!("Result: {}", z);
}
