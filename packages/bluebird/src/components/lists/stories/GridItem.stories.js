import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import GridItem from '../GridItem';

storiesOf('Building Blocks/GridItem', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default GridItem', () => (
      <GridItem onClick={action('clicked')} />
  ))
