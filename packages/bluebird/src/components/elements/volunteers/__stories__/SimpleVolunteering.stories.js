import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleVolunteering from '../SimpleVolunteering';

storiesOf('OpenMob Elements/SimpleVolunteering', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default SimpleVolunteering', () => (
    <SimpleVolunteering onClick={action('clicked')} />
  ));
