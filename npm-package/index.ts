#!/usr/bin/env node

import os from 'os';
import path from 'path';
import { execFileSync } from 'child_process';

function main() {
    const currentPlatform = os.platform().toString();

    const selectedFunction = platformToFunctionMap.get(currentPlatform)

    if (selectedFunction === undefined) {
        console.log(`The platform ${currentPlatform} is not yet supported :(
            Please open an issue if one does not already exist: https://github.com/JoshuaTheMiller/Grapple/issues/new`);
    }

    selectedFunction?.call(selectedFunction, process.argv)
}

const platformToFunctionMap = new Map([
    ["win32", runWindows],
    ["linux", runLinux],
    ["darwin", runDarwin]
]);

function runWindows(args: string[]) {
    runCommandWithArguments("windows", "Grapple.exe", args);
}

function runLinux(args: string[]) {
    runCommandWithArguments("linux", "Grapple", args);
}

function runDarwin(args: string[]) {
    runCommandWithArguments("darwin", "Grapple", args);
}

function runCommandWithArguments(subFolder: string, commandName: string, args: string[]) {
    const userArgs = args.splice(2);
    const actualPath = path.join(__dirname, subFolder, commandName);
    const result = execFileSync(actualPath, userArgs, { encoding: 'utf8', windowsHide: true });
    console.log(result);
}

main();