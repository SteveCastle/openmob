import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Card from '../Card';

storiesOf('Layout/Card', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Card', () => (
      <Card onClick={action('clicked')} />
  ))
