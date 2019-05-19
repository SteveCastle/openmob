import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import App from '../App';

storiesOf('Layout/App', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default App', () => (
      <App onClick={action('clicked')} />
  ))
