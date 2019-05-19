import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Input from '../Input';

storiesOf('Layout/Input', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Input', () => (
      <Input onClick={action('clicked')} />
  ))
