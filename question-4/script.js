function ParseProductionSteps(arrayOfPlants) {
  if (arrayOfPlants == undefined) return "Please give a valid input!";

  const object = {};

  arrayOfPlants.forEach((plant) => {
    object[plant.id] = {};
    let sum = 0;
    plant.productionSteps.forEach((step) => {
      object[plant.id][step.step] = step.time + sum;
      sum += step.time;
    });
  });

  return object;
}

let plants = [
  {
    id: "Plant3",
    productionSteps: [
      { step: "Step1", time: 1 },
      { step: "Step2", time: 2 },
      { step: "Step3", time: 3 },
    ],
  },
  {
    id: "Plant4",
    productionSteps: [
      { step: "StepA", time: 10 },
      { step: "StepB", time: 10 },
      { step: "StepC", time: 10 },
    ],
  },
];
console.log(ParseProductionSteps(plants));

plants = [
  {
    id: "Plant5",
    productionSteps: [
      { step: "Step1", time: 5 },
      { step: "Step2", time: 6 },
      { step: "Step3", time: 7 },
    ],
  },
  {
    id: "Plant6",
    productionSteps: [
      { step: "StepA", time: 3 },
      { step: "StepB", time: 3 },
      { step: "StepC", time: 3 },
    ],
  },
];
console.log(ParseProductionSteps(plants));

plants = [
  {
    id: "Plant7",
    productionSteps: [
      { step: "Step1", time: 1 },
      { step: "Step2", time: 1 },
      { step: "Step3", time: 1 },
    ],
  },
  {
    id: "Plant8",
    productionSteps: [
      { step: "StepA", time: 1 },
      { step: "StepB", time: 1 },
      { step: "StepC", time: 1 },
    ],
  },
];
console.log(ParseProductionSteps(plants));

plants = [
  {
    id: "Plant1",
    productionSteps: [
      { step: "Phase1", time: 5 },
      { step: "Phase2", time: 10 },
      { step: "Phase3", time: 15 },
    ],
  },
  {
    id: "Plant2",
    productionSteps: [
      { step: "PhaseA", time: 3 },
      { step: "PhaseB", time: 6 },
      { step: "PhaseC", time: 9 },
      { step: "PhaseD", time: 12 },
      { step: "PhaseE", time: 15 },
    ],
  },
];
console.log(ParseProductionSteps(plants));

plants = [
  {
    id: "Factory1",
    productionSteps: [
      { step: "Task1", time: 2 },
      { step: "Task2", time: 4 },
    ],
  },
  {
    id: "Factory2",
    productionSteps: [
      { step: "TaskA", time: 6 },
      { step: "TaskB", time: 12 },
      { step: "TaskC", time: 18 },
      { step: "TaskD", time: 24 },
    ],
  },
  {
    id: "Factory3",
    productionSteps: [
      { step: "TaskI", time: 1 },
      { step: "TaskII", time: 2 },
      { step: "TaskIII", time: 3 },
      { step: "TaskIV", time: 4 },
      { step: "TaskV", time: 5 },
    ],
  },
];
console.log(ParseProductionSteps(plants));

plants = [];
console.log(ParseProductionSteps(plants));
