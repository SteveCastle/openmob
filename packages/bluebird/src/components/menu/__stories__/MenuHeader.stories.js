import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import MenuHeader from '../MenuHeader';

storiesOf('Layout/MenuHeader', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default MenuHeader', () => (
      <MenuHeader onClick={action('clicked')} />
  ))
