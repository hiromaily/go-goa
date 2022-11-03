const nextJest = require("next/jest");

const createJestConfig = nextJest({
  // includes `next.config.js` and `.env` for test
  dir: "./",
});

// Jest custom settings
const customJestConfig = {
  testEnvironment: "jest-environment-jsdom",
};

module.exports = createJestConfig(customJestConfig);