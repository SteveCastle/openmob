import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleBoycott from '../SimpleBoycott';

storiesOf('OpenMob Elements/SimpleBoycott', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default SimpleBoycott', () => (
    <SimpleBoycott onClick={action('clicked')} />
  ));
