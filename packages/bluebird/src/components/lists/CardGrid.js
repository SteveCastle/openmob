import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
} from 'styled-system';

const StyledGrid = styled.ul`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  ${flexWrap}
  list-style: none;
`;

function ImageGrid({ children }) {
  return (
    <StyledGrid display="flex" flexWrap="wrap" width={1}>
      {children}
    </StyledGrid>
  );
}

ImageGrid.propTypes = {
  children: PropTypes.arrayOf(
    PropTypes.shape({
      uri: PropTypes.string,
      width: PropTypes.number,
      height: PropTypes.number,
      title: PropTypes.string,
      id: PropTypes.string,
      slug: PropTypes.string,
    })
  ),
};
ImageGrid.defaultProps = {};

export default ImageGrid;
