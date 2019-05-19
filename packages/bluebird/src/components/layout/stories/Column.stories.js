import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Column from '../Column';

storiesOf('Layout/Column', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Column', () => (
      <Column onClick={action('clicked')} />
  ))
