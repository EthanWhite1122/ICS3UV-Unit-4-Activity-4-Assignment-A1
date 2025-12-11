/**
 * @author Ethan White
 * @version 1.0.0
 * @date 2025-12-11
 * @fileoverview  This program will come up with a random number and make the user guess it
 */
let target = Math.floor(Math.random() * 10) + 1;
let guess = -1;

while (guess !== target && guess !== 0) {
  const input = prompt("Enter a number between 1 and 10 (0 to quit):");

  if (input === null) {
    console.log("Program terminated.");
    break;
  }

  guess = Number(input);

  if (guess === 0) {
    console.log("Program terminated.");
  } else if (guess === target) {
    console.log("You win.");
  } else {
    console.log("Wrong guess. Try again.");
  }
}

console.log("\nDone.")