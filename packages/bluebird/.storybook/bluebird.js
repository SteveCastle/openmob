import { create } from '@storybook/theming';

export default create({
    base: 'light',
  
    colorPrimary: 'hotpink',
    colorSecondary: 'deepskyblue',

    // Typography
    fontBase: '"Open Sans", sans-serif',
    fontCode: 'monospace',
  
  
    // Form colors
    inputBg: 'white',
    inputBorder: 'silver',
    inputTextColor: 'black',
    inputBorderRadius: 4,
  
    brandTitle: 'OpenMob: BlueBird',
    brandUrl: 'https://openmob.us',
  });