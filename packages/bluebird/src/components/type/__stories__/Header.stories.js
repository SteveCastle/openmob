import React from 'react';
import faker from 'faker';
import { storiesOf } from '@storybook/react';
import Header from '../Header';

storiesOf('Typography/Header', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default Header', () => <Header>{faker.name.findName()}</Header>)
  .add('Dark Header', () => <Header dark>{faker.name.findName()}</Header>);
