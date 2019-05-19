import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import index from '../index';

storiesOf('Layout/index', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default index', () => (
      <index onClick={action('clicked')} />
  ))
