import fs from 'node:fs';

function getFirstLine(buffer) {
    return buffer.toString().split('\n').shift();
}

const RED = '\x1b[0;31m';
const GREEN = '\x1b[0;32m';
const BLUE = '\x1b[1;34m';

const BOLD = '\x1b[1m';

/** End Of Style, removes all attributes (formatting and colors) */
const EOS = '\x1b[0m';

const commitMsgContent = fs.readFileSync('.git/COMMIT_EDITMSG', 'utf-8');
const msg = getFirstLine(commitMsgContent).replace(/\s{2,}/g, ' ');

// TODO: Change and update regex below according to your preferences.
const pattern = /^\[((BF)|(HF)|F|R|(DX)|(CUSTOM))\]: .+$/;

// TODO: Change and update example below according to your preferences.
const exampleMsg = [
    '[F]: added flowcharts',
    '[F]: added customization for flowcharts',
    '[R]: Moved components from ~/cmp to ~/lib/cmp',
    '[HF]: Expearence -> Experience',
    '[BF]: Fixed cache invalidating on main page',
];

const successMsgs = [
    `\n[ok] ${GREEN}*********** Succesfully commit changes. Goodjob! ***********${EOS}`,
    `\n[ok] ${GREEN}*********** All good! Powerrr! ***********${EOS}`,
    `\n[ok] ${GREEN}*********** omg what have I done? jk, All good! ***********${EOS}`,
    `\n[ok] ${GREEN}*********** May the Force be with you ***********${EOS}`,
    `\n[ok] ${GREEN}*********** “It's alive! It's alive!” - Frankenstein, 1931 ***********${EOS}`,
    `\n[ok] ${GREEN}*********** “I'll be back.” - The Terminator, 1984 ***********${EOS}`,
    `\n[ok] ${GREEN}*********** Magic Mirror on the wall, who is the fairest one of all? ***********${EOS}`,
    `\n[ok] ${GREEN}*********** "Bond. James Bond." ***********${EOS}`,
    `\n[ok] ${GREEN}*********** "They may take our lives, but they'll never take our freedom!" ***********${EOS}`,
    `\n[ok] ${GREEN}*********** To infinity and beyond! ***********${EOS}`,
    `\n[ok] ${GREEN}*********** ( ͡• ͜ʖ ͡•) niceeeee ~ ***********${EOS}`,
    `\n[ok] ${GREEN}*********** Another oneeeee! ***********${EOS}`,
    `\n[ok] ${GREEN}*********** We Had To Use Dark Magic To Make This Work ***********${EOS}`,
    `\n[ok] ${GREEN}*********** Never gonna give you up, Never gonna let you down ***********${EOS}`,
    `\n[ok] ${GREEN}*********** Too lazy to write descriptive success message ***********${EOS}`,
    `\n[ok] ${GREEN}*********** hehehehhe nice ***********${EOS}`,
    `\n[ok] ${GREEN}*********** “You is kind. You is smart. You is important.” ***********${EOS}`,
    `\n[ok] ${GREEN}*********** Lorem ipsum dolor sit amet, consectetur adipiscing... ***********${EOS}`,
    `\n[ok] ${GREEN}*********** My mama always said life was like a box of chocolates. You never know what you're gonna get. ***********${EOS}`,
];

const failMsgs = [
    `\n${RED}************* Invalid Git Commit Message *************${EOS}`,
    `\n${RED}************* Houston, we have a problem. *************${EOS}`,
    `\n${RED}************* You can't handle the truth! *************${EOS}`,
    `\n${RED}************* “I am your father.” *************${EOS}`,
    `\n${RED}************* “Hasta la vista, baby.” *************${EOS}`,
    `\n${RED}************* Something bad happened *************${EOS}`,
    `\n${RED}************* Catastrophic failure *************${EOS}`,
    `\n${RED}************* Really????? ( ˘︹˘ ) *************${EOS}`,
    `\n${RED}************* Dont cry, you can commit again with the correct format *************${EOS}`,
    `\n${RED}************* Well, nobody's perfect. Please try again.” *************${EOS}`,
    `\n${RED}************* Bu yao zhi yang :( ” *************${EOS}`,
    `\n${RED}************* I'm sorry. I'm afraid I can't do that. *************${EOS}`,
    `\n${RED}************* We Had To Use Dark Magic To Make This Work, but we still fail :( *************${EOS}`,
];

const result = pattern.test(msg);
const commitResultMsg = result ? 'SUCCESS' : 'FAILED';

if (result) {
    console.log(successMsgs[Math.floor(Math.random() * successMsgs.length)]);
    console.log(`${BOLD}Commit message:${EOS} ${BLUE}${msg}${EOS}\n`);

    process.exit(0);
} else {
    console.log(failMsgs[Math.floor(Math.random() * failMsgs.length)]);

    console.log(`${BOLD}Commit message:${EOS} ${RED}${msg}${EOS}`);
    console.log(`${BOLD}Example of correct message:${EOS}`);
    exampleMsg.forEach((example, i) => {
        console.log(`  ${i + 1}) ${BLUE}${example}${EOS}`);
    });
    console.log();
    process.exit(1);
}
