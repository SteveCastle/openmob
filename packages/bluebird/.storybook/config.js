import { configure } from '@storybook/react';
import { addDecorator } from '@storybook/react';
import { withThemesProvider } from 'storybook-addon-styled-component-theme';
import { withInfo } from '@storybook/addon-info';
import themes from '../src/themes';

function loadStories() {
  require('../src/stories');
}
addDecorator(withInfo);

// Setup themes.
const themeList = [...themes];
addDecorator(withThemesProvider(themeList));

configure(loadStories, module);
