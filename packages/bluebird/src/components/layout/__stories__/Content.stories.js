import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Content from '../Content';

storiesOf('Layout System/Content', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default Content', () => <Content onClick={action('clicked')} />);
