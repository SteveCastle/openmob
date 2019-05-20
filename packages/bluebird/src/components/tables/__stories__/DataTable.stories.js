import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import DataTable from '../DataTable';

storiesOf('Building Blocks/DataTable', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default DataTable', () => <DataTable onClick={action('clicked')} />);
