#!/usr/bin/env node

const fs = require('fs');

function getNamedArgumentValue(argumentList, name) {
    const potentialValue = argumentList.find(elem => elem.startsWith(name));

    if (potentialValue === null) {
        return "";
    }

    const lengthOfName = name.length + 1;
    return potentialValue.substring(lengthOfName);
};

function readUtf8File(filePath) {
    // Encoding has to be specified, otherwise weirdness will occur
    return fs.readFileSync(filePath, { encoding: "utf-8" })
}

// tenets
// prefer short emoji over long
// prefer sensible emoji

// feat --> :bulb: (new idea?)
// fix --> :bug: (bug fix :D)
// docs --> :scroll: or :ledger: or :book: or :memo: or :pencil: (short, still represents text)
// style --> :art: (being stylist is an artform)
// refactor --> :soon: (soon, this codebase will be even better... Feeling sarcastic today I think)
// test --> :mag: (inspecting via tests)
// chore --> :zzz: (chores can be boring, but are still necessary!)
// break --> :warning: or :x: or :boom: (something to grab attention, as this is potentially disruptive)

// Yes, I did go through the entire emoji list for this...

// The first element of process.argv is the execution path,
// and the second is the path for the js file. If you want
// to access the params related to git hooks, they will start at 
// index 2!
const arguments = process.argv;
const pathToCommitMessage = arguments[2];
const pathToConfigFile = getNamedArgumentValue(arguments, "FilePath");

const config = JSON.parse(readUtf8File(pathToConfigFile ?? "./CommitMessageTags.json"));

const codeToSkipValidation = config.tagToAllowSkip.tag;
const recognizedShortMessages = config.standardTags.sort((a, b) => (a.category > b.category) ? 1 : -1);

const message = readUtf8File(pathToCommitMessage);
const flattenedRecognized = (recognizedShortMessages.map(rsm => rsm.tags)).flat();
flattenedRecognized.push(codeToSkipValidation);

for (const shortCode of flattenedRecognized) {
    const startsWithShortCode = message.startsWith(shortCode);

    if (startsWithShortCode && shortCode === codeToSkipValidation) {
        fs.writeFileSync(pathToCommitMessage, message.substring(2).trimStart());
    }

    if (startsWithShortCode) {
        console.log("\x1b[32mNice commit message!\x1b[0m (passes Standard Tags validation)")
        process.exit();
    }
}

console.log(`\x1b[37m \x1b[41m
-----------------------------
!! COMMIT MESSAGE REJECTED !!
-----------------------------\x1b[0m

Current Git Hooks are configured such that commit messages must start with
a tag from one of the following categories! To skip validation, start the
commit message with the characters \x1b[33m${codeToSkipValidation}\x1b[0m. Those characters will be removed
from the final message.
`);

for (const shortMessage of recognizedShortMessages) {
    console.log(`\x1b[33m${shortMessage.category}\x1b[0m: ${shortMessage.description}`)
    const tagsToPrint = shortMessage.tags.join("\x1b[0m, \x1b[32m");
    console.log(" | Tags --> ", "\x1b[32m", tagsToPrint, "\x1b[0m");
}
console.log("");
console.log("-----------------------------");

process.exit(1);