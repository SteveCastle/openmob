import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Paragraph from '../Paragraph';

storiesOf('Layout/Paragraph', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Paragraph', () => (
      <Paragraph onClick={action('clicked')} />
  ))
