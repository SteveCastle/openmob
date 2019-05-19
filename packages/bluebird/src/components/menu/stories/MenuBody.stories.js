import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import MenuBody from '../MenuBody';

storiesOf('Layout/MenuBody', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default MenuBody', () => (
      <MenuBody onClick={action('clicked')} />
  ))
