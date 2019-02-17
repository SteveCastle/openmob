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
  display
} from 'styled-system';

const StyledRow = styled.div`
  ${borders}
  ${borderRadius}
  ${space}
  ${width}
  ${fontSize}
  ${color}
  ${display}
  box-sizing: border-box;

`;

function Row({ dark, children, tracing }) {
  return (
    <StyledRow
      display={'flex'}
      border={0}
      py={3}
      width={1}
      bg={
        tracing
          ? `fuschia.${tracing}`
          : dark
          ? 'darkBackground'
          : 'lightBackground'
      }
      fontSize={1}
      color={dark ? 'lightText' : 'darkText'}
    >
      {children}
    </StyledRow>
  );
}

Row.propTypes = {
  children: PropTypes.arrayOf(PropTypes.element),
  dark: PropTypes.bool,
  tracing: PropTypes.number
};
Row.defaultProps = {};

export default Row;
