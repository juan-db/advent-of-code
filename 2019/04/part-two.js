const start = 146888;
const end = 612564;

let valid = 0;
for (let i = start; i <= end; ++i) {
	if (isValid(i)) {
		valid += 1;
	}
}
console.log(valid);

function isValid(num) {
	num = `${num}`;

	for (let i = 1, last = num[0]; i < num.length; ++i) {
		// Verify the numbers increase
		if (num[i] < last) {
			return false;
		}
		last = num[i];
	}

	for (let i = 0; i < num.length;) {
		let count = 0;
		let last = num[i];
		while (i < num.length && num[i] == last) {
			++count;
			++i;
		}
		if (count == 2) {
			return true;
		}
	}
	return false;
}
