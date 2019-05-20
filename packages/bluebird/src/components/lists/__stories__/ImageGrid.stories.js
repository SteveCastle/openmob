import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import ImageGrid from '../ImageGrid';
import GridItem from '../GridItem';

storiesOf('Building Blocks/ImageGrid', module)
  .addParameters({
    info: {
      inline: true,
    },
  })
  .add('Default ImageGrid', () => (
    <ImageGrid onClick={action('clicked')}>
      <GridItem uri={'/images/6.jpg'} />
      <GridItem uri={'/images/6.jpg'} />
      <GridItem uri={'/images/6.jpg'} />
      <GridItem uri={'/images/6.jpg'} />
    </ImageGrid>
  ));
