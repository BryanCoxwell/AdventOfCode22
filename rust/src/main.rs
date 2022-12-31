mod day1;
mod day2;

use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();
    if args.len() < 2 {
        println!("Provide the number of the day you want to run.");
        return;
    }
    let day = &args[1];

    match day.as_str() {
        "1" => day1::run(),
        "2" => day2::run(),
        _ => println!("Day {} not implemented yet", day),
    }
}
