import { addParameters, configure } from '@storybook/react';
import { addDecorator } from '@storybook/react';
import { withThemesProvider } from 'storybook-addon-styled-component-theme';
import { withInfo } from '@storybook/addon-info';
import { withBackgrounds } from '@storybook/addon-backgrounds';
import themes from '../src/themes';
import bluebird from './bluebird';

addParameters({
  options: {
    theme: bluebird,
  },
});

const req = require.context('../src/components', true, /\.stories\.js$/);

function loadStories() {
  req.keys().forEach(filename => req(filename));
}

// Setup info plugin
addDecorator(withInfo);

// Setup backgrounds plugin.
addDecorator(
  withBackgrounds([{ name: 'Light', value: '#fff', default: true }])
);

// Setup themes.
const themeList = [...themes];
addDecorator(withThemesProvider(themeList));

configure(loadStories, module);
