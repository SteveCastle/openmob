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
`;

function Column({ dark, children, tracing }) {
  return (
    <StyledColumn
      border={0}
      p={3}
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
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Column.defaultProps = {};

export default Column;
