function processText(text, operations, ranges) {
  let chunks = [];
  // Add the beginning part of the text that is not included in the ranges.
  if (ranges[0][0] > 0) {
    chunks.push(text.substring(0, ranges[0][0]));
  }
  // Process each range.
  for (let i = 0; i < ranges.length; i++) {
    let range = ranges[i];
    let chunk = text.substring(range[0], range[1] + 1);
    // Apply each operation to the chunk.
    for (let operation of operations) {
      chunk = operation(chunk);
    }
    chunks.push(chunk);
    // Add the part of the text between this range and the next range.
    let nextRangeStart = i < ranges.length - 1 ? ranges[i + 1][0] : text.length;
    let middleChunk = text.substring(range[1] + 1, nextRangeStart);
    chunks.push(middleChunk);
  }
  console.log(chunks.join(""));
  return chunks.join("");
}

let text = "Hello, World!";
let operations = [
  function (text) {
    return text.toUpperCase();
  },
  function (text) {
    return text.replace(/o/g, "0");
  },
];
let ranges = [
  [0, 4],
  [7, 12],
];
processText(text, operations, ranges);

text = "Programming";
operations = [
  function (text) {
    return text.split("").reverse().join("");
  },
  function (text) {
    return text.charAt(0).toUpperCase() + text.slice(1);
  },
];
ranges = [
  [0, 4],
  [6, 10],
];
processText(text, operations, ranges);

text = "Hello, World!";
operations = [
  function (text) {
    return text.split("").reverse().join("");
  },
  function (text) {
    return text.charAt(0).toUpperCase() + text.slice(1);
  },
];
ranges = [
  [6, 6],
  [8, 11],
];
processText(text, operations, ranges);

text = "ABCD";
operations = [
  function (text) {
    return text.split("").reverse().join("");
  },
];
ranges = [[1, 2]];
processText(text, operations, ranges);

text = "Hello, OpenAI!";
operations = [
  function (text) {
    return text.split("").join("_");
  }, // replace spaces with underscores
  function (text) {
    return text.charAt(0).toUpperCase() + text.slice(1).toLowerCase();
  }, // make first character uppercase, rest lowercase
];
ranges = [
  [0, 4],
  [7, 12],
];
processText(text, operations, ranges);

text = "JavaScript is awesome!";
operations = [
  function (text) {
    return text.split(" ").join("");
  }, // remove all spaces
  function (text) {
    return (
      text.slice(0, text.length / 2) + text.slice(text.length / 2).toUpperCase()
    );
  }, // make the second half of the text uppercase
];
ranges = [
  [0, 9],
  [14, 21],
];
processText(text, operations, ranges);
