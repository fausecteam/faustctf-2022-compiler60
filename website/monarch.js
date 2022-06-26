({
    defaultToken: 'invalid',
    brackets: [ ["'BEGIN'","'END'",'keyword.block'],
        ['[',']','delimiter.square'],
        ['(',')','delimiter.parenthesis']],
    escapes: /\\(?:[nrt\\"]|x[0-9A-Fa-f]{1,4}|[0-7]{1,3})/,
    
    tokenizer: {
    root: [
        // keywords
        [/\'/, {token: 'keyword', next: 'keyword', goBack: 1}],
        // identifiers
        [/[a-zA-Z][a-zA-Z0-9 \t\r\n]*/, 'identifier'],
        // whitespace
        [/[ \t\r\n]+/, 'white'],
        // brackets and operators
        [/[()\[\]]/, '@brackets'],
        [/[=<>≤≥≠¬∧∨⊃×÷+-]|(\:[ \t\r\n]*\=)/, 'operator'],
        // numbers
        [/[0-9][0-9 \t\r\n]*/, 'number'],
        // delimiter
        [/[;,]/, 'delimiter'],
        // strings
        [/"/,  { token: 'string.quote', bracket: '@open', next: '@string' }]
    ],
    keyword: [
    // 'BEGIN'
    [/\'[ \t\r\n]*B[ \t\r\n]*E[ \t\r\n]*G[ \t\r\n]*I[ \t\r\n]*N[ \t\r\n]*\'/, 
        {token: 'keyword.block', bracket: '@open', next: '@pop'}],
    // 'END'
    [/\'[ \t\r\n]*E[ \t\r\n]*N[ \t\r\n]*D[ \t\r\n]*\'/,
        {token: 'keyword.block', bracket: '@open', next: '@pop'}],
    // 'COMMENT'
    [/\'[ \t\r\n]*C[ \t\r\n]*O[ \t\r\n]*M[ \t\r\n]*M[ \t\r\n]*E[ \t\r\n]*N[ \t\r\n]*T[ \t\r\n]*\'/,
        {token: 'comment', switchTo: 'comment'}],
    ["'", {token: 'keyword', switchTo: 'in_keyword'}]
    ],
    in_keyword: [
        [/[A-Z \t\r\n]+/, 'keyword'],
        [/\'/, 'keyword', '@pop']
    ],
    comment: [
        [/[^;]+/, 'comment' ],
        [";", 'comment', '@pop'],
    ],
    string: [
        [/[^\\"]+/,  'string'],
        [/@escapes/, 'type'],
        [/\\./,      'invalid'],
        [/"/,        { token: 'string.quote', bracket: '@close', next: '@pop' } ]
    ]
    },
})