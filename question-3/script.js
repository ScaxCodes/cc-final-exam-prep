function checkPasswordStrength(password) {
  if (password === undefined) return "Please enter a password!";

  let points = 0;

  checkLength(password);
  checkForCapital(password);
  checkForLowercase(password);
  checkForNumber(password);
  checkForSpecialCharacters(password);

  if (points === 5) return "very strong";
  else if (points === 4) return "strong";
  else if (points === 3) return "medium";
  else return "weak";

  function checkLength(str) {
    if (str.length > 10) points++;
  }

  function checkForCapital(str) {
    for (let i = 0; i < str.length; i++) {
      let asciiIndex = str.charCodeAt(i);
      if (asciiIndex >= 65 && asciiIndex <= 90) {
        points++;
        return;
      }
    }
  }

  function checkForLowercase(str) {
    for (let i = 0; i < str.length; i++) {
      let asciiIndex = str.charCodeAt(i);
      if (asciiIndex >= 97 && asciiIndex <= 122) {
        points++;
        return;
      }
    }
  }

  function checkForNumber(str) {
    for (let i = 0; i < str.length; i++) {
      let asciiIndex = str.charCodeAt(i);
      if (asciiIndex >= 48 && asciiIndex <= 57) {
        points++;
        return;
      }
    }
  }

  function checkForSpecialCharacters(str) {
    const specialCharacters = [37, 94, 38, 42, 44, 64, 33, 35, 36];

    for (let i = 0; i < str.length; i++) {
      let asciiIndex = str.charCodeAt(i);
      if (specialCharacters.includes(asciiIndex)) {
        points++;
        return;
      }
    }
  }
}

let password = "password";
console.log(checkPasswordStrength(password));

password = "password123";
console.log(checkPasswordStrength(password));

password = "Password123";
console.log(checkPasswordStrength(password));

password = "Pord123!";
console.log(checkPasswordStrength(password));

password = "LongPassword123!";
console.log(checkPasswordStrength(password));

password = "PASSWORD!";
console.log(checkPasswordStrength(password));

password = "1234567890!";
console.log(checkPasswordStrength(password));

password = "";
console.log(checkPasswordStrength(password));
