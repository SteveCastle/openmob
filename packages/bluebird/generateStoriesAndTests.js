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
  return `import React from 'react';
import { shallow } from 'enzyme';
import ${componentName} from '../${componentName}';

it('renders without crashing', () => {
  shallow(<${componentName} />);
});
`
}

recursive(path.join(__dirname, "src/components"), ["*stories.js"], function (err, files) {
  files.map(file => {
    fs.mkdir(`${path.dirname(file)}/stories`, { recursive: true }, (err) => {
      if (err) console.log(err);
    });
    fs.writeFile(`${path.dirname(file)}/stories/${path.basename(file, '.js')}.stories.js`, generateStory(path.basename(file, '.js')), { overwrite: false }, function (err) {
      if (err) {
        return console.log(err);
      }
      console.log("Story written!");
    });
    fs.mkdir(`${path.dirname(file)}/__tests__`, { recursive: true }, (err) => {
      if (err) console.log(err);
    });
    fs.writeFile(`${path.dirname(file)}/__tests__/${path.basename(file, '.js')}.spec.js`, generateTest(path.basename(file, '.js')), { overwrite: false }, function (err) {
      if (err) {
        return console.log(err);
      }
      console.log("Test written!");
    });
  })
});