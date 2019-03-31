import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  maxWidth,
  height,
  fontSize,
  color,
  borders,
  borderRadius,
  display,
  flexWrap,
  flexDirection,
  alignItems,
  justifyContent
} from 'styled-system';

const StyledCard = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${maxWidth}
  ${height}
  ${fontSize}
  ${display}
  ${flexWrap}
  ${flexDirection}
  ${color}
  ${display}
  ${alignItems}
  ${justifyContent}
`;

function Card({ children }) {
  return (
    <StyledCard
      display="flex"
      bg="gray.0"
      p={2}
      border={0}
      borderRadius={2}
      flexWrap="wrap"
      flexDirection="column"
      fontSize={1}
    >
      {children}
    </StyledCard>
  );
}

Card.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element)
};
Card.defaultProps = {};

export default Card;
