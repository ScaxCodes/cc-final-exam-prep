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
