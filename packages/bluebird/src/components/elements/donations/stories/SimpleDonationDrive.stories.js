import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleDonationDrive from '../SimpleDonationDrive';

storiesOf('OpenMob Elements/SimpleDonationDrive', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimpleDonationDrive', () => (
      <SimpleDonationDrive onClick={action('clicked')} />
  ))
