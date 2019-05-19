import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Navbar from '../Navbar';

storiesOf('Layout/Navbar', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Navbar', () => (
      <Navbar onClick={action('clicked')} />
  ))
