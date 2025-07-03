// tracker-version-updater.js

const fs = require('fs');

module.exports.readVersion = function (contents) {
  const match = contents.match(/var\s+Version\s+=\s+"([^"]+)"/);
  if (match) {
    return match[1];
  }
  throw new Error('Version string not found');
};

module.exports.writeVersion = function (contents, version) {
  return contents.replace(/var\s+Version\s+=\s+"([^"]+)"/, `var Version = "${version}"`);
};
