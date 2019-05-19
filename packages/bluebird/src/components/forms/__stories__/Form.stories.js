import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Form from '../Form';

storiesOf('Layout/Form', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Form', () => (
      <Form onClick={action('clicked')} />
  ))
