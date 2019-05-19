var recursive = require("recursive-readdir");
var path = require("path");
const fs = require('fs');

function generateStory(componentName) {
  return `import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import ${componentName} from '../${componentName}';

storiesOf('Layout/${componentName}', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default ${componentName}', () => (
      <${componentName} onClick={action('clicked')} />
  ))
`
}
function generateTest(componentName) {
  return `

  `
}

recursive(path.join(__dirname, "src/components/layout/sidebar"), ["*stories.js"], function (err, files) {
  files.map(file => {
    fs.mkdir(`${path.dirname(file)}/stories`, { recursive: true }, (err) => {
      if (err) console.log(err);
    });
    fs.writeFile(`${path.dirname(file)}/stories/${path.basename(file, '.js')}.stories.js`, generateStory(path.basename(file, '.js')), { overwrite: true }, function (err) {
      if (err) {
        return console.log(err);
      }

      console.log("The file was saved!");
    });
  })
});