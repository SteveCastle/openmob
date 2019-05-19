import React from 'react';
import faker from 'faker';

import { storiesOf } from '@storybook/react';
import Paragraph from '../Paragraph';

storiesOf('Typography/Paragraph', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default Paragraph', () => (
    <Paragraph>{faker.lorem.paragraph()}</Paragraph>
  ))
