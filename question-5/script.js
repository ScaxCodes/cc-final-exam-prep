function getHierarchy(arrayOfEmployees) {
  const object = {};

  arrayOfEmployees.forEach((employee) => {
    const managerTracker = [];
    let currentManager = null;

    if ("manager" in employee) {
      currentManager = employee.manager;
      managerTracker.push(currentManager);
      let currentEmployee = arrayOfEmployees[currentManager - 1];
      while ("manager" in currentEmployee) {
        let lastManager = currentManager;
        currentManager = arrayOfEmployees[lastManager - 1].manager;
        managerTracker.push(currentManager);
        currentEmployee = arrayOfEmployees[currentManager - 1];
      }
      managerTracker.forEach((manager) => {
        if (!object[manager]) object[manager] = [];
        object[manager].push(employee.id);
      });
    }
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
loop for each entry:
  if entry has manager, add manager to trackerarray
  repeat until no manager:
    if manager has manager, add manager to trackerarray
  add id to all manager-ids from tracker-array into return-object
*/
