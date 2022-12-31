use std::fs::File;
use std::io::prelude::*;
use std::cmp;

const INPUT: &str = "../../inputs/day1_input.txt";

pub fn run(){
    println!("Day 1, Part 1: {}", part1());
    println!("Day 1, Part 2: {}", part2());
}

fn part1() -> i32 {
    let file: File = File::open(INPUT).expect("File not found.");
    let mut lines = std::io::BufReader::new(file).lines();

    let mut elfs_max: i32 = 0;
    let mut overall_max: i32 = 0;

    while let Some(line) = lines.next() {
        let line = line.unwrap();
        if line.len() == 0 {
            overall_max = cmp::max(elfs_max, overall_max);
            elfs_max = 0;
        } else {
            let val: i32 = line.parse().unwrap();
            elfs_max += val;
        }
    }
    return overall_max;
}

fn part2() -> i32 {
    println!("Part 2 not implemented!");
    0
}