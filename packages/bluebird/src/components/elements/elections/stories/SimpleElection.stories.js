import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleElection from '../SimpleElection';

storiesOf('OpenMob Elements/SimpleElection', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SimpleElection', () => (
      <SimpleElection onClick={action('clicked')} />
  ))
