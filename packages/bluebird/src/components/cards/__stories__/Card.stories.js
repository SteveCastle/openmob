import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Card from '../Card';

storiesOf('Building Blocks/Card', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Basic Cards', () => (
    <Card onClick={action('clicked')} label="Primary Card" />
  ));
