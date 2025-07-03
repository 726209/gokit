// .config/types.js
const types = [
  {
    value: 'feat',
    name: 'feat:     ✨ 新增功能',
    section: '✨ Features | 新增功能',
    emoji: ':recycle:'
  },
  { value: 'fix', name: 'fix:      🐛 修复缺陷', section: '🐛 Bug Fixes | 修复缺陷' },
  { value: 'docs', name: 'docs:     📝 文档变更', section: '📚 Documentation | 文档变更' },
  {
    value: 'style',
    name: 'style:    💄 代码格式(空白、格式化、去掉分号等，不影响功能)',
    section: '💄 Style | 代码格式',
    hidden: true
  },
  {
    value: 'refactor',
    name: 'refactor: ♻️  重构代码(不包括Bug修复、功能新增)',
    section: '♻️ Refactor | 代码重构'
  },
  { value: 'perf', name: 'perf:     ⚡️ 性能优化', section: '⚡ Performance | 性能优化' },
  { value: 'test', name: 'test:     ✅ 添加测试', section: '✅ Tests | 测试', hidden: true },
  {
    value: 'build',
    name: 'build:    🔨 构建相关（如Webpack、Rollup配置）',
    section: '🔨‍Build System | 构建相关',
    hidden: true
  },
  {
    value: 'chore',
    name: 'chore:    📦 其他修改配置、工具等杂项',
    section: '📦 Chores | 其他更新',
    hidden: true
  },
  {
    value: 'ci',
    name: 'ci:       🚀 持续集成（修改 CI 配置、脚本等）',
    section: '🚀Continuous Integration | CI 配置',
    hidden: true
  },
  {
    value: 'revert',
    name: 'revert:   ⏪ 回退提交',
    section: '⏪ Reverts | 回退提交',
    hidden: true
  },
  {
    value: 'WIP',
    name: 'WIP:         正在进行的工作',
    section: 'WIP: 正在进行的工作',
    hidden: true
  }
];

module.exports = types;
