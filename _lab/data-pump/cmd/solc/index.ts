import { readFileSync, writeFile } from 'fs'
import solc from 'solc'

const CONTRACT_FILE = 'Registry'
const content = readFileSync(`../../common/${CONTRACT_FILE}.sol`).toString()

const input = {
  language: 'Solidity',
  sources: {
    [CONTRACT_FILE]: {
      content: content,
    },
  },
  settings: {
    outputSelection: {
      '*': {
        '*': ['*'],
      },
    },
  },
}

const compiled = solc.compile(JSON.stringify(input))
const output = JSON.parse(compiled)

const abi: object = output.contracts[CONTRACT_FILE][CONTRACT_FILE].abi

const bytecode: string =
  output.contracts[CONTRACT_FILE][CONTRACT_FILE].evm.bytecode.object

writeFile(
  `../../common/${CONTRACT_FILE}.json`,
  JSON.stringify(abi, null, 2),
  (err) => {
    if (err) {
      return console.error(err)
    }
    return console.log(`Saved ${CONTRACT_FILE}.json!!`)
  }
)

writeFile(`../../common/${CONTRACT_FILE}.bin`, bytecode, (err) => {
  if (err) {
    return console.error(err)
  }
  return console.log(`Saved ${CONTRACT_FILE}.bin!!`)
})
