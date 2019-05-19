import React from 'react';

import { storiesOf } from '@storybook/react';
import Masonry from '../Masonry'

storiesOf('Photo Galleries/Masonry', module)
  .addParameters({
    info: {
      inline: true
    }
  })
  .add('Masonry', () => (
    <Masonry
      itemsPerRow={[2, 3]}
      images={[
        { src: '/images/1.jpg', aspectRatio: 3968 / 2976 },
        { src: '/images/2.jpg', aspectRatio: 5344 / 3563 },
        { src: '/images/3.jpg', aspectRatio: 5653 / 3769 },
        { src: '/images/4.jpg', aspectRatio: 3648 / 5472 },
        { src: '/images/5.jpg', aspectRatio: 4570 / 3264 },
        { src: '/images/6.jpg', aspectRatio: 5472 / 3648 },
        { src: '/images/7.jpg', aspectRatio: 122 / 182 },
        { src: '/images/8.jpg', aspectRatio: 122 / 182 }
      ]}
    />
  ));