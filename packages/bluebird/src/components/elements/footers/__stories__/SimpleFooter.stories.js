import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleFooter from '../SimpleFooter';

storiesOf('OpenMob Elements/SimpleFooter', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default SimpleFooter', () => (
    <SimpleFooter onClick={action('clicked')} />
  ));
