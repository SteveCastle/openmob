import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleShop from '../SimpleShop';

storiesOf('Layout/SimpleShop', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimpleShop', () => (
      <SimpleShop onClick={action('clicked')} />
  ))
