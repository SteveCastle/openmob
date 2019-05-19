import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Masonry from '../Masonry';

storiesOf('Layout/Masonry', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Masonry', () => (
      <Masonry onClick={action('clicked')} />
  ))
