// .config/.cz-config.js
const types = require('./types');

console.info('>>> local config in repo: ', __dirname);

module.exports = {
  types: types.map((t) => ({ value: t.value, name: t.name, emoji: t.emoji })),
  useEmoji: true,
  // emojiAlign: 'center',

  // scope 类型，针对当前项目
  // scopes: [{ name: '默认' }, { name: '组件' }],
  scopes: [
    ['default', '默认'],
    ['components', '组件'],
    ['module', '模块'],
    ['dependency', '依赖'],
    ['other', '其他']
    // 如果选择 custom ,后面会让你再输入一个自定义的 scope , 也可以不设置此项， 把后面的 allowCustomScopes 设置为 true
    // ['custom', '以上都不是？我要自定义']
  ].map(([value, description]) => {
    return {
      value,
      name: `${value.padEnd(30)} (${description})`
    };
  }),

  // usePreparedCommit: false, // to re-use commit from ./.git/COMMIT_EDITMSG
  // allowTicketNumber: false,
  // isTicketNumberRequired: false,
  // ticketNumberPrefix: 'TICKET-',
  // ticketNumberSuffix:'',
  // ticketNumberRegExp: '\\d{1,5}',

  // 可以设置 scope 的类型跟 type 的类型匹配项，例如: 'fix'
  // it needs to match the value for field type. Eg.: 'fix'
  /*
    scopeOverrides: {
      fix: [

        {name: 'merge'},
        {name: 'style'},
        {name: 'e2eTest'},
        {name: 'unitTest'}
      ]
    },
    */
  // override the messages, defaults are as follows
  messages: {
    type: '请选择本次提交类型:',
    scope: '\n请选择本次提交的影响范围 (可选):',
    // used if allowCustomScopes is true
    customScope: '请输入本次提交的自定义影响范围:',
    subject: '请简要描述本次提交说明(必填):\n',
    body: '请输入本次提交的详细描述(可选，使用"|"换行):\n',
    breaking: '非兼容性说明 (可选，使用"|"换行):\n',
    footer: '列举出所有变更的 ISSUES CLOSED (可选，例如：#31, #34):',
    confirmCommit: '确认要使用以上信息提交?'
  },
  // 是否允许自定义填写 scope ，设置为 true ，会自动添加两个 scope 类型 [{ name: 'empty', value: false },{ name: 'custom', value: 'custom' }]
  allowCustomScopes: true,
  // allowBreakingChanges: ['feat', 'fix'],
  // skip any questions you want
  // skipQuestions: ['body'],

  // limit subject length
  subjectLimit: 100
  // breaklineChar: '|', // It is supported for fields body and footer.
  // footerPrefix : 'ISSUES CLOSED:'
  // askForBreakingChangeFirst : true, // default is false
};
