# Codeforces-Hacker

[![Build Status][1]][2] [![codecov.io][3]][4] [![codebeat badge](https://codebeat.co/badges/bd93263b-c637-4a10-8ac1-0e178d4a632b)](https://codebeat.co/projects/github-com-limstash-codeforces-hacker-master)

[1]: https://dev.azure.com/limstash/Codeforces-Hacker/_apis/build/status/limstash.Codeforces-Hacker?branchName=master "Build Status badge"
[2]: https://dev.azure.com/limstash/Codeforces-Hacker/_build?definitionId=1 "Azure Build Status"
[3]: https://codecov.io/gh/limstash/Codeforces-Hacker/branch/master/graph/badge.svg?token=6pMHmpIYtG "Coverage badge"
[4]: https://codecov.io/gh/limstash/Codeforces-Hacker "Codecov Status"

Help you automatically hack in Educational Round and Div 3

## Build

You could download pre-compile executable files directly at [here](https://github.com/limstash/Codeforces-Hacker/releases)

You could also download source files and use ``make`` to compile

## Run

You need to set up ``config.json``

Here is an example

```json
{
    "target" : {
        "contest" : 1176, // The contest ID you want to hack
        "index" : "B" // The problem index you want to hack
    },

    "testcase" : {
        "inputFile" : "./data.in", // Input testcase filepath
        "outputFile" : "./data.ans" // Output testcase filepath
    },

    "account" : {
        "username" : "user", // Your Codeforces account username (Optional)
        "password" : "pass" // Your Codeforces account Password (Optional)
    },

    "settings": {
        "autoLogin" : true, // Login your account automatically (need to fill in the account info)
        "autoHack" : true, // Hacked automatically (need check on autoLogin)
    }
}
```