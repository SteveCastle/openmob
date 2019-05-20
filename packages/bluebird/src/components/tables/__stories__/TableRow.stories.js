import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import TableRow from '../TableRow';

storiesOf('Building Blocks/TableRow', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default TableRow', () => <TableRow onClick={action('clicked')} />);
