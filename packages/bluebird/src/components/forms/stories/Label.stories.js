import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Label from '../Label';

storiesOf('Layout/Label', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Label', () => (
      <Label onClick={action('clicked')} />
  ))
