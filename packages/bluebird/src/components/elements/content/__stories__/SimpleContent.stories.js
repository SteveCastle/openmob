import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import SimpleContent from '../SimpleContent';

storiesOf('OpenMob Elements/SimpleContent', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default SimpleContent', () => (
    <SimpleContent onClick={action('clicked')} />
  ));
