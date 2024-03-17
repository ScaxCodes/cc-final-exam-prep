function getHierarchy(arrayOfEmployees) {
  const object = {};
  const trackManagers = [];

  arrayOfEmployees.forEach((employee) => {
    if ("manager" in employee) {
      trackManagers.push(employee.manager);
      if (!object[employee.manager]) object[employee.manager] = [];
      if (trackManagers.includes(employee.id)) {
        object[employee.id].forEach((employeeBelow) => {
          object[employee.manager] = [];
          object[employee.manager].push(employeeBelow);
        });
      }
      object[employee.manager].push(employee.id);
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
