import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimplePetition from '../SimplePetition';

storiesOf('Layout/SimplePetition', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimplePetition', () => (
      <SimplePetition onClick={action('clicked')} />
  ))
