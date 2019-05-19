import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Header from '../Header';

storiesOf('Layout/Header', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Header', () => (
      <Header onClick={action('clicked')} />
  ))
