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

function Column({ dark, children, tracing, size }) {
  return (
    <StyledColumn
      border={0}
      p={3}
      mx={2}
      width={size / 12}
      bg={
        tracing
          ? `blue.${tracing}`
          : dark
          ? 'darkBackground'
          : 'lightBackground'
      }
      fontSize={1}
      color={dark ? 'lightText' : 'darkText'}
    >
      {children}
    </StyledColumn>
  );
}

Column.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  size: PropTypes.number,
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Column.defaultProps = { size: 12 };

export default Column;
