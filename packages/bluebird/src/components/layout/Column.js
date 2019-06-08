import React from 'react';
import PropTypes from 'prop-types';
import styled from 'styled-components';
import {
  space,
  width,
  fontSize,
  color,
  borders,
  borderRadius
} from 'styled-system';

const StyledColumn = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  box-sizing: border-box;
`;

function Column({ children, tracing, size, disableSpacing }) {
  return (
    <StyledColumn
      border={0}
      p={disableSpacing ? 0 : 3}
      mx={disableSpacing ? 0 : 2}
      my={disableSpacing ? 0 : 2}
      width={[1, size / 12]}
      bg={tracing ? `blue.${tracing}` : null}
      fontSize={1}
    >
      {children}
    </StyledColumn>
  );
}

Column.propTypes = {
  children: PropTypes.oneOfType([
    PropTypes.arrayOf(PropTypes.node),
    PropTypes.node
  ]),
  size: PropTypes.number,
  tracing: PropTypes.number,
  disableSpacing: PropTypes.bool
};
Column.defaultProps = { size: 12 };

export default Column;
