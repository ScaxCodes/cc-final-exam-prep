# cc-final-exam-prep

Some practise questions for my final exam of the course "CourseCareers - Software Development Fundamentals"

## Question 1

Write a function that accepts an array of numbers and returns the 2 largest numbers in the array. The largest values should be returned in an array where the first element in the array is the larger of the two elements. You may assume you will always be given an array that contains at least 2 numbers.

## Question 2

Write a function that has two parameters, the first being an array of strings, the second being an array of unique integers (whole numbers). Your function should return the number of strings in the input array that have a length contained in the second array. If the array of numbers is empty you should return 0. See the sample input for further explanation.

## Question 3

Write a function that accepts a plain text password and returns the relative strength of the password based on the following criteria.

Types of Passwords:

- Weak Password: 0 - 2 points.
- Medium Password: 3 points.
- Strong Password: 4 points.
- Very Strong Password: 5 Points.

Criteria For Points:

- Length: 1 point if more than 10 characters
- Capital letters: 1 point if at least 1 capital letter is present
- Lowercase letters: 1 point if at least 1 lowercase letter is present
- Numbers: 1 point if at least 1 number is present
- Special characters: 1 point if any of the following characters are present: %, ^, &, \*, (, ),
  @, !, #, $ (i.e all the special characters accessible from the number keys 1, 2, 3, 4... etc)

Your function should return “weak”, “medium”, “strong” or “very strong” to classify the given password.

## Question 4

Widgets Inc is a production company with several plants around the world. They have recently started a project to analyze the efficiency of their production process. In each plant, widgets are manufactured in a sequence of steps, where each step can only begin once the previous step has been completed. Each plant has a different sequence of steps, and the time it takes to complete each step also varies between plants.

You are given an array of JavaScript objects, where each object represents a plant and contains the properties "id" and "productionSteps". The "id" property is a unique string identifier for the plant, and "productionSteps" is an array where each element is another object that has a "step" property (a string identifying the step) and a "time" property (a number indicating the time it takes to complete that step).

Your task is to write a JavaScript function that takes this array as input and returns a new JavaScript object. The keys of this object should be the ids of the plants, and the values should be objects where the keys are the steps and the values are the total time it would take to reach that step (including the time for that step itself). The steps should appear in the order they are completed.

## Question 5

ABC corporation has many employees, some of which are managers, others which are employees. Each employee has a manager, and some managers have a manager who oversees them. Each employee knows who manages them but unfortunately the managers have lost their org chart making it difficult for them to remember which employees are below them in the hierarchy.

For this question you’ll be given an array of javascript objects representing employee information. Each employee has a unique “id” property, and every employee that has a manager has a “manager” property that stores the “id” of their manager. Your job is to return a new javascript object that contains the ids of managers as properties, where each property stores an array of the unique ids of all employees beneath that manager in the hierarchy. You may assume any manager referenced will always exist in the array. See the sample inputs for further explanation.

## Question 6

RapidText Inc. specializes in text processing and analytics. They have developed a new tool that can analyze a large text string and perform a series of operations on it. The operations are provided as an array of functions that each take a string as input and return a modified string as output. The functions are applied in the order they are given in the array.

The problem is that some of these operations are computationally intensive and can take a long time to run, so RapidText wants to be able to apply these operations only to the portions of the text that actually need them.

Your task is to write a JavaScript function that takes a text string, an array of operations (functions), and an array of ranges (each range being an array of two numbers indicating start and end positions in the text). The function should apply each operation to only the text within the specified ranges, and return the modified text.

Each operation will be a function of the form function(text) { /_ ... _/ } and will return a string.

Note: Ranges are inclusive and are based on zero-indexing. If a range is [2, 5], it refers to the part of the string from the 3rd to the 6th character, inclusive.
