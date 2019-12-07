// As with all my previous solutions, I'm doing it the simple way instead of the
// optimal way.

// Shamelessly taken from Rosetta Code
struct DigitIter(i32, i32);
 
impl Iterator for DigitIter {
    type Item = i32;
    fn next(&mut self) -> Option<Self::Item> {
        if self.0 == 0 {
            None
        } else {
            let ret = self.0 % self.1;
            self.0 /= self.1;
            Some(ret)
        }
    }
}
 
fn is_valid(num: i32) -> bool {
    let mut iter = DigitIter(num, 10);
    let mut last = iter.next().unwrap();
    let mut found_double = false;
    for digit in iter {
        if digit > last {
            return false;
        } else if digit == last {
            foundDouble = true;
        } else {
            last = digit;
        }
    }
    found_double
}

fn main() {
    let start = 146888;
    let end = 612564;
    let mut count = 0;

     for i in start..=end {
         if is_valid(i) {
             count += 1;
         }
     }

     println!("{}", count)
}

