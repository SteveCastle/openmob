import React from 'react';

import { storiesOf } from '@storybook/react';
import { action } from '@storybook/addon-actions';
import ImageGrid from '../ImageGrid';
import ImageGridItem from '../ImageGridItem';

storiesOf('Building Blocks/ImageGrid', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Default ImageGrid', () => (
    <ImageGrid onClick={action('clicked')}>
      <ImageGridItem uri={'/images/6.jpg'} />
      <ImageGridItem uri={'/images/6.jpg'} />
      <ImageGridItem uri={'/images/6.jpg'} />
      <ImageGridItem uri={'/images/6.jpg'} />
    </ImageGrid>
  ));
