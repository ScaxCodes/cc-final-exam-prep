// TODO: Debug script v1 step by step, what need to be done after first loop?
/*

1 Loop each entry
if manager in entry
  add manager id to managertracker
if manager id not in newobject
  add manager id to newobject

problem: wenn manager am ende weitere einträge bekommt
lösung: erst am ende wenn alles ausgelesen ist berechnen
*/

function getHierarchy(arrayOfEmployees) {
  const object = {};
  let counterlvl1 = 0;

  arrayOfEmployees.forEach((employee) => {
    const managerTracker = [];
    let currentManager = null;
    let counterlvl2 = 0;

    if ("manager" in employee) {
      currentManager = employee.manager;
      managerTracker.push(currentManager);
      let test = arrayOfEmployees[currentManager - 1];
      // console.log(arrayOfEmployees);
      // console.log(arrayOfEmployees[1]);
      while ("manager" in test) {
        let lastManager = currentManager;
        // console.log(arrayOfEmployees[lastManager - 1]);
        currentManager = arrayOfEmployees[lastManager - 1].manager;
        managerTracker.push(currentManager);
        test = arrayOfEmployees[currentManager - 1];
        counterlvl2++;
        // console.log(`lvl2: ${counterlvl2}`);
      }
      managerTracker.forEach((manager) => {
        if (!object[manager]) object[manager] = [];
        object[manager].push(employee.id);
      });
    }

    counterlvl1++;
    // console.log(`lvl1: ${counterlvl1}`);
  });

  return object;
}

let employees = [
  { id: "1", manager: "2" },
  { id: "2", manager: "3" },
  { id: "3" },
  { id: "4", manager: "2" },
  { id: "5", manager: "4" },
];
console.log(getHierarchy(employees));

employees = [{ id: "1", manager: "2" }, { id: "2" }];
console.log(getHierarchy(employees));

employees = [{ id: "1" }, { id: "2" }, { id: "3" }, { id: "4" }];
console.log(getHierarchy(employees));

employees = [
  { id: "1", manager: "2" },
  { id: "2", manager: "3" },
  { id: "3", manager: "4" },
  { id: "4", manager: "5" },
  { id: "5" },
];
console.log(getHierarchy(employees));

employees = [{ id: "1", manager: "3" }, { id: "2", manager: "3" }, { id: "3" }];
console.log(getHierarchy(employees));
/* 
SOLUTION PSEUDO CODE
if entry has manager, add to trackerarray
repeat until no manager:
  if manager has manager, add to trackerarray
add id to all newObjectID from trackerarray
*/
