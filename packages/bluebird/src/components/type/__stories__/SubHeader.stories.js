import React from 'react';
import faker from 'faker';

import { storiesOf } from '@storybook/react';
import SubHeader from '../SubHeader';

storiesOf('Typography/SubHeader', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default SubHeader', () => (
    <SubHeader>{faker.company.catchPhrase()}</SubHeader>
  ))
