/** @type {import('ts-jest').JestConfigWithTsJest} **/
module.exports = {
  testEnvironment: "node",
  transform: {
    "^.+.tsx?$": ["ts-jest", {}],
  },
  clearMocks: true,
  collectCoverage: true,
  collectCoverageFrom: [
    './src/**/*.ts?(x)',
    '!./src/containers/dashboadr.tsx',
    '!./src/index.tsx',
    '!./src/containers/login/index.ts'
  ],
  coverageDirectory: "<rootDir>/coverage",
  coverageReporters: [
    ['html', {
      subdir: 'html'
    }]
  ],
  preset: 'ts-jest',
};