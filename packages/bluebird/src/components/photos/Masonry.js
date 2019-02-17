import chunk from 'lodash/chunk';
import sum from 'lodash/sum';

import React from 'react';
import PropTypes from 'prop-types';
import { Box } from 'rebass';

const Masonry = ({ images = [], itemsPerRow: itemsPerRowByBreakpoints }) => {
  const aspectRatios = images.map(image => image.aspectRatio);

  // For each breakpoint, calculate the aspect ratio sum of each row's images
  const rowAspectRatioSumsByBreakpoints = itemsPerRowByBreakpoints.map(
    itemsPerRow =>
      // Split images into groups of the given size
      chunk(aspectRatios, itemsPerRow).map(rowAspectRatios =>
        // Sum aspect ratios of images in the given row
        sum(rowAspectRatios)
      )
  );

  return (
    <div>
      {images.map((image, i) => (
        <Box
          key={image.src}
          as={'img'}
          src={image.src}
          fluid={image}
          title={image.caption}
          width={rowAspectRatioSumsByBreakpoints.map(
            // Return a value for each breakpoint
            (rowAspectRatioSums, j) => {
              // Find out which row the image is in and get its aspect ratio sum
              const rowIndex = Math.floor(i / itemsPerRowByBreakpoints[j]);
              const rowAspectRatioSum = rowAspectRatioSums[rowIndex];

              return `${(image.aspectRatio / rowAspectRatioSum) * 100}%`;
            }
          )}
          css={{ display: 'inline-block' }}
        />
      ))}
    </div>
  );
};

export default Masonry;

Masonry.propTypes = {
  images: PropTypes.arrayOf(
    PropTypes.shape({
      caption: PropTypes.string,
      src: PropTypes.string.isRequired,
      aspectRatio: PropTypes.number.isRequired
    })
  ).isRequired,
  itemsPerRow: PropTypes.number.isRequired
};
Masonry.defaultProps = { images: [], itemsPerRow: 3 };
