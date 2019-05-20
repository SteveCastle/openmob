import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import Layout from '../Layout';

storiesOf('Layout System/Layout', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default Layout', () => <Layout onClick={action('clicked')} />);
