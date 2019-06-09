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
  flexWrap
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
  justify-content: stretch;
  list-style: none;
`;

function CardGrid({ children }) {
  return (
    <StyledGrid display="flex" flexWrap="wrap" width={1}>
      {children}
    </StyledGrid>
  );
}

CardGrid.propTypes = {
  children: PropTypes.arrayOf(
    PropTypes.shape({
      uri: PropTypes.string,
      width: PropTypes.number,
      height: PropTypes.number,
      title: PropTypes.string,
      id: PropTypes.string,
      slug: PropTypes.string
    })
  )
};
CardGrid.defaultProps = {};

export default CardGrid;
