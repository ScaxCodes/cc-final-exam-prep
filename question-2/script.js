function getNumberOfMatchingStrings(stringArray, lengthsArray) {
  let counter = 0;

  stringArray.forEach((word) => {
    lengthsArray.forEach((length) => {
      if (word.length === length) counter++;
    });
  });

  return counter;
}

let strings = ["this", "is", "a", "test"];
let numbers = [4, 2];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["apple", "orange", "banana", "grape"];
numbers = [5, 6];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["one", "two", "three", "four", "five"];
numbers = [];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["red", "blue", "green", "yellow"];
numbers = [3, 4, 5];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["I", "am", "not", "a", "robot"];
numbers = [1, 2, 3];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["unique", "words", "have", "specific", "lengths"];
numbers = [5, 6, 7];
console.log(getNumberOfMatchingStrings(strings, numbers));

strings = ["longer", "words", "like", "hippopotamus", "and", "crocodile"];
numbers = [4, 5, 6, 7, 8, 9, 10, 11, 12];
console.log(getNumberOfMatchingStrings(strings, numbers));
