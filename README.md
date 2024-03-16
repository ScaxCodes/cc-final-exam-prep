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
