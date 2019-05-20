import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Widget from '../Widget';

storiesOf('Forms/Widget', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default Widget', () => <Widget onClick={action('clicked')} />);
