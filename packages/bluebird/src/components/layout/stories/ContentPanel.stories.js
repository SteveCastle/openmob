import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import ContentPanel from '../ContentPanel';

storiesOf('Layout System/ContentPanel', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default ContentPanel', () => (
      <ContentPanel onClick={action('clicked')} />
  ))
