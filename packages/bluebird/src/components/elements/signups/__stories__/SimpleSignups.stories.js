import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleSignups from '../SimpleSignups';

storiesOf('OpenMob Elements/SimpleSignups', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimpleSignups', () => (
      <SimpleSignups onClick={action('clicked')} />
  ))
