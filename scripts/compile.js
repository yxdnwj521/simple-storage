const fs = require('fs');
const path = require('path');
const solc = require('solc');

// 使用相对路径
const contractPath = path.resolve(__dirname, '../SimpleStorage.sol');
const source = fs.readFileSync(contractPath, 'utf8');

const input = {
    language: 'Solidity',
    sources: {
        'SimpleStorage.sol': {
            content: source,
        },
    },
    settings: {
        outputSelection: {
            '*': {
                '*': ['abi', 'evm.bytecode.object'],
            },
        },
    },
};

const output = JSON.parse(solc.compile(JSON.stringify(input)));

// 打印完整的编译输出
console.log('Complete compilation output:');
console.log(JSON.stringify(output, null, 2));

if (output.errors) {
    console.error('Compilation errors:');
    let hasFatalError = false;
    output.errors.forEach(error => {
        console.error(`- ${error.formattedMessage}`);
        if (error.severity === 'error') {
            hasFatalError = true;
        }
    });
    if (hasFatalError) {
        process.exit(1);
    }
}

// 检查 contracts 是否存在
if (!output.contracts || !output.contracts['SimpleStorage.sol']) {
    console.error('No contracts found in the compilation output.');
    process.exit(1);
}

const contracts = output.contracts['SimpleStorage.sol'];

if (!fs.existsSync(path.resolve(__dirname, '../build'))) {
    fs.mkdirSync(path.resolve(__dirname, '../build'));
}

for (let contract in contracts) {
    fs.writeFileSync(
        path.resolve(__dirname, `../build/${contract}.json`),
        JSON.stringify(contracts[contract], null, 2)
    );
}