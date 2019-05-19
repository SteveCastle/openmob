import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Container from '../Container';

storiesOf('Layout/Container', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Container', () => (
      <Container onClick={action('clicked')} />
  ))
