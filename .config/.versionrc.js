// .config/.versionrc.js
const types = require('./types');
const { name } = require('../package.json');

const path = require('path');
// const glob = require('glob')
//
// const find_podspec_file = ()=>{
//     const files = glob.sync(__dirname + '/*.podspec')
//     if (files && files.length > 0) {
//         return path.basename(files[0]);
//     }
//     throw new Error(`目录${__dirname}不存在 *.podspec`)
// }
//
// const podspec = find_podspec_file()
// const pod_name = path.parse(podspec).name
// const sdk_header = `${pod_name}/Classes/${pod_name}.h`
console.log(`配置文件：${name}`);

module.exports = {
  // skip: {
  //     tag: true,
  // },
  // preset: 'conventionalcommits',
  // tagPrefix: '',
  header: '# 📝 Changelog\n\n> 本文件自动生成，请勿手动修改。\n',
  //types为Conventional Commits标准中定义，目前支持
  //https://github.com/conventional-changelog/commitlint/tree/master/%40commitlint/config-conventional
  types: types.map((t) => ({ type: t.value, section: t.section, hidden: t.hidden })),

  //hash链接
  // commitUrlFormat: "http://gitlab.cmss.com/BI/{{repository}}/commit/{{hash}}",
  //issue链接
  // issueUrlFormat: "http://jira.cmss.com/browse/{{id}}",
  //server-version自动commit的模板
  // releaseCommitMessageFormat: 'build: 发布 {{currentTag}} 🎉 \n\nCode Source From: Self Code \nDescription: \nJira: # \n市场项目编号（名称）：\n\n变更说明详见 CHANGELOG.md\n'
  //需要server-version更新版本号的文件
  bumpFiles: [
    {
      filename: 'package.json',
      // The `json` updater assumes the version is available under a `version` key in the provided JSON document.
      type: 'json'
    },

    {
      filename: 'internal/version/version.go',
      updater: '.config/version.js'
    }
  ]
};
