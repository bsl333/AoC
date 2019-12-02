const fs = require("fs");
const { promisify } = require("util");
const path = require("path");
const readFilePromise = promisify(fs.readFile);

const calcFuel = mass => Math.floor(mass / 3) - 2;

readFilePromise(path.join(__dirname, "..", "input.txt"), {
  encoding: "utf-8"
})
  .then(data => {
    console.log(`*****************\nSolution1:`);
    const fuelNeeded = data
      .split("\n")
      .reduce((acc, mass) => acc + calcFuel(+mass), 0);
    console.log(fuelNeeded);
    console.log(`*****************\n`);
    return data;
  })
  .then(data => {
    console.log(`*****************\nSolution 2:`);
    const fuelNeeded = data.split("\n").reduce((acc, mass) => {
      let fuel = calcFuel(mass);
      let totalFuel = 0;
      while (fuel > 0) {
        totalFuel += fuel;
        fuel = calcFuel(fuel);
      }
      return acc + totalFuel;
    }, 0);
    console.log(fuelNeeded);
    console.log("*****************");
  });
