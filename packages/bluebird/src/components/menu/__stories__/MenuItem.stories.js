import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import MenuItem from '../MenuItem';

storiesOf('Navigation/MenuItem', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default MenuItem', () => <MenuItem onClick={action('clicked')} />);
