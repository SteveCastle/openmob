import { configure } from '@storybook/react';
import { addDecorator } from '@storybook/react';
import { withThemesProvider } from 'storybook-addon-styled-component-theme';
import themes from '../src/themes';

function loadStories() {
  require('../src/stories');
}
// Setup themes.
const themeList = [...themes];
addDecorator(withThemesProvider(themeList));

configure(loadStories, module);
