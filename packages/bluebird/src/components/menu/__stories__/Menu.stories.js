import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Menu from '../Menu';

storiesOf('Layout/Menu', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Menu', () => (
      <Menu onClick={action('clicked')} />
  ))
