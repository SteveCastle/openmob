import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Spinner from '../Spinner';

storiesOf('Layout/Spinner', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Spinner', () => (
      <Spinner onClick={action('clicked')} />
  ))
