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
  boxShadow,
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
  ${boxShadow}
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
  height: ${({ fillHeight }) => (fillHeight ? '100%' : null)};
`;

function Card({ children, width = 1, fillHeight }) {
  return (
    <StyledCard
      bg="gray.0"
      p={2}
      m={2}
      border={0}
      boxShadow={0}
      borderRadius={2}
      display="inline-block"
      flexWrap="wrap"
      flexDirection="row"
      fontSize={1}
      width={width}
      fillHeight={fillHeight}
    >
      {children}
    </StyledCard>
  );
}

Card.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  width: PropTypes.number,
  fillHeight: PropTypes.bool
};
Card.defaultProps = {};

export default Card;
