import React from 'react';

import { storiesOf } from '@storybook/react';
import {
  DataTable,
  TableCell,
  TableRow,
  TableHeader,
  TableHeaderCell,
} from '../';

storiesOf('Building Blocks/DataTable', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default DataTable', () => (
    <DataTable>
      <TableHeader>
        <TableHeaderCell>Greetings</TableHeaderCell>
        <TableRow>
          <TableCell>Hello</TableCell>
        </TableRow>
      </TableHeader>
    </DataTable>
  ));
